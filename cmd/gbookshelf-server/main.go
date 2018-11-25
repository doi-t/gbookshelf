package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpc "google.golang.org/grpc"
)

func main() {
	// Create a gRPC Server with gRPC interceptor.
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	// Register gbookshelf-server gRPC service implementations.
	var bookshelf bookShelfServer
	gbookshelf.RegisterBookShelfServer(srv, bookshelf)
	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.Register(srv)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	// Start your http server for prometheus.
	go func() {
		if err := http.ListenAndServe(":9000", nil); err != nil { // TODO: make port number environment variable

			log.Fatalf("Unable to start a http server for Prometheus: %v", err)
		}
	}()

	l, err := net.Listen("tcp", ":8888") // TODO: make port number environment variable
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type bookShelfServer struct{}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb" //TODO: where should I define it and keep it configurable for testing?
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

		var book gbookshelf.Book
		if err := proto.Unmarshal(b[:l], &book); err != nil {
			return nil, fmt.Errorf("cloud not read book: %v", err)
		}
		b = b[l:]
		books.Books = append(books.Books, &book)
	}
}

func (bookShelfServer) Add(ctx context.Context, book *gbookshelf.Book) (*gbookshelf.Book, error) {
	b, err := proto.Marshal(book)
	if err != nil {
		return nil, fmt.Errorf("could not encode book: %v", err)
	}

	// TODO: find the best place to manage protobuf data other than a local file
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

	return book, nil
}

func (bss bookShelfServer) Remove(ctx context.Context, rb *gbookshelf.Book) (*gbookshelf.Book, error) {
	l, err := bss.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return nil, err
	}
	removed := false
	var newList gbookshelf.Books
	for _, book := range l.Books {
		if book.Title == rb.Title {
			log.Printf("Remove %v from bookshelf\n", book)
			removed = true
			continue
		}
		newList.Books = append(newList.Books, book)
	}

	if removed != true {
		return nil, fmt.Errorf("could not find a book that you specified. Check title again: %v", rb)
	}

	// TODO: find a better way to remove a book from db
	err = os.Remove(dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not remove %s: %v", dbPath, err)
	}

	for _, book := range newList.Books {
		bss.Add(ctx, book)
	}

	return rb, nil
}

func (bss bookShelfServer) Update(ctx context.Context, b *gbookshelf.Book) (*gbookshelf.Book, error) {
	l, err := bss.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return nil, err
	}
	updated := false
	var newList gbookshelf.Books
	for _, book := range l.Books {
		if book.Title == b.Title {
			var p int32
			if b.Page == 0 {
				p = book.Page
			} else {
				p = b.Page
			}

			var c int32
			if b.Current == 0 {
				c = book.Current
			} else {
				c = b.Current
			}

			if c > p {
				return nil, fmt.Errorf("The current page position (%d) can be not larger than the number of page (%d) of %s", c, p, book.Title)
			}

			book = &gbookshelf.Book{
				Title:   book.Title,
				Page:    p,
				Done:    b.Done,
				Current: c,
			}
			log.Printf("Update %v-> %v\n", b, book)
			updated = true
		}
		newList.Books = append(newList.Books, book)
	}
	if updated != true {
		return nil, fmt.Errorf("could not find a book title: %v", b.Title)
	}

	// TODO: find a better way to update a book in db
	err = os.Remove(dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not remove %s: %v", dbPath, err)
	}

	for _, book := range newList.Books {
		bss.Add(ctx, book)
	}

	return b, nil
}
