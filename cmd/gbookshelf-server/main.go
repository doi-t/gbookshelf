package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var bookshelf bookShelfServer
	gbookshelf.RegisterBookShelfServer(srv, bookshelf)
	l, err := net.Listen("tcp", ":8888") // TODO: make it environment variable
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type bookShelfServer struct{}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func (bookShelfServer) List(ctx context.Context, void *gbookshelf.Void) (*gbookshelf.Books, error) {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return nil, fmt.Errorf("cloud not read %s: %v", dbPath, err)
	}

	var books gbookshelf.Books
	for {
		if len(b) == 0 {
			return &books, nil
		} else if len(b) < sizeOfLength {
			return nil, fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return nil, fmt.Errorf("cloud not decode message length: %v", err)
		}

		b = b[sizeOfLength:]

		var book gbookshelf.BookState
		if err := proto.Unmarshal(b[:l], &book); err != nil {
			return nil, fmt.Errorf("cloud not read book: %v", err)
		}
		b = b[l:]
		books.Books = append(books.Books, &book)
		fmt.Printf("Loaded book: %v\n", book)
	}
}

func (bookShelfServer) Add(ctx context.Context, book *gbookshelf.Book) (*gbookshelf.BookState, error) {
	bs := &gbookshelf.BookState{
		Title: book.Title,
		Page:  book.Page,
		Done:  false,
	}
	b, err := proto.Marshal(bs)
	if err != nil {
		return nil, fmt.Errorf("could not encode book: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("cloud not open %s: %v", dbPath, err)
	}

	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return nil, fmt.Errorf("could not encode length of message: %v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return nil, fmt.Errorf("could not write book to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("cloud not close file %s: %v", dbPath, err)
	}

	return bs, nil
}

func (bss bookShelfServer) Remove(ctx context.Context, rb *gbookshelf.Book) (*gbookshelf.Void, error) {
	l, err := bss.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return &gbookshelf.Void{}, err
	}
	var newList gbookshelf.Books
	for _, book := range l.Books {
		if book.Title == rb.Title {
			log.Printf("Remove %v from bookshelf\n", book)
			continue
		}
		newList.Books = append(newList.Books, book)
	}

	err = os.Remove(dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not remove %s: %v", dbPath, err)
	}

	for _, book := range newList.Books {
		b, err := proto.Marshal(book)
		if err != nil {
			return nil, fmt.Errorf("could not encode book: %v", err)
		}

		var f *os.File
		f, err = os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("cloud not open %s: %v", dbPath, err)
		}

		if err := binary.Write(f, endianness, length(len(b))); err != nil {
			return nil, fmt.Errorf("could not encode length of message: %v", err)
		}

		_, err = f.Write(b)
		if err != nil {
			return nil, fmt.Errorf("could not write book to file: %v", err)
		}

		if err := f.Close(); err != nil {
			return nil, fmt.Errorf("cloud not close file %s: %v", dbPath, err)
		}
	}

	return &gbookshelf.Void{}, nil
}
