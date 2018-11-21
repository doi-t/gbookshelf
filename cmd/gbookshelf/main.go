package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/doi-t/gbookshelf/gbookshelf"
	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func add(title string) error {
	book := &gbookshelf.Book{
		Title: title,
		Page:  100,
		Done:  false,
	}
	b, err := proto.Marshal(book)
	if err != nil {
		return fmt.Errorf("could not encode book: %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("cloud not open %s: %v", dbPath, err)
	}

	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return fmt.Errorf("could not encode length of message: %v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("could not write book to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("cloud not close file %s: %v", dbPath, err)
	}

	return nil
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("cloud not read %s: %v", dbPath, err)
	}

	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("remaining odd %d bytes, what to do?", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return fmt.Errorf("cloud not decode message length: %v", err)
		}

		b = b[sizeOfLength:]

		var book gbookshelf.Book
		if err := proto.Unmarshal(b[:l], &book); err != nil {
			return fmt.Errorf("cloud not read book: %v", err)
		}
		b = b[l:]

		if book.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s (P%d)\n", book.Title, book.Page)
	}
}
