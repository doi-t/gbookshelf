package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	grpc "google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	client := gbookshelf.NewBookShelfClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(context.Background(), client)
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, client gbookshelf.BookShelfClient, title string) error {
	_, err := client.Add(ctx, &gbookshelf.Title{Title: title})
	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully")
	return nil
}

func list(ctx context.Context, client gbookshelf.BookShelfClient) error {
	l, err := client.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return fmt.Errorf("cloud not fetch books: %v", err)
	}
	for _, b := range l.Books {
		if b.Done {
			fmt.Printf("👍")
		} else {
			fmt.Printf("😱")
		}
		fmt.Printf(" %s (P%d)\n", b.Title, b.Page)
	}
	return nil
}
