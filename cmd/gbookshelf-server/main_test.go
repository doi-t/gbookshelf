package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
)

var books = []gbookshelf.Book{
	{Title: "Designing Data-Intensive Applications", Page: 624, Done: false, Current: 20},
	{Title: "default value book"},
}

func initBookshelf(dbPath string, t *testing.T) {
	t.Helper()
	err := os.Remove(dbPath)
	if err != nil {
		t.Errorf("%s did not exist but it's fine.", dbPath)
	}
	bss := bookShelfServer{}
	for _, b := range books {
		_, err := bss.Add(context.Background(), &b)
		if err != nil {
			t.Fatalf("failed to initialize bookshelf: %v", err)
		}
	}
}

func TestAdd(t *testing.T) {
	tt := []struct {
		name    string
		title   string
		page    int32
		current int32
	}{
		{"add a book", "Designing Data-Intensive Applications", 624, 20},
		{"add a book with default values", "default value book", 0, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bss := bookShelfServer{}
			b := &gbookshelf.Book{
				Title:   tc.title,
				Page:    tc.page,
				Current: tc.current,
			}
			a, err := bss.Add(context.Background(), b)
			if err != nil {
				t.Fatalf("Add should be succeeded; failed: %v", err) // TODO: test it
			}
			if b != a {
				t.Fatalf("Added book should be %v; got %v", b, a)
			}
		})
	}
}

func TestList(t *testing.T) {
	tt := []struct {
		name    string
		title   string
		page    int32
		current int32
		num     int
	}{
		{"list added book", "Designing Data-Intensive Applications", 624, 20, 2},
	}
	initBookshelf(dbPath, t) // Using a global variable in main.go
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bss := bookShelfServer{}
			l, err := bss.List(context.Background(), &gbookshelf.Void{})
			if err != nil {
				t.Fatalf("List should be succeeded; failed: %v", err)
			}
			if len(l.Books) != tc.num {
				t.Fatalf("the number of books should be %d; got %d", tc.num, len(l.Books))
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tt := []struct {
		name    string
		title   string
		page    int32
		done    bool
		current int32
		err     string
	}{
		{name: "update added book", title: "Designing Data-Intensive Applications", current: 400},
		{name: "update zero page book", title: "default value book", page: 111},
		{name: "finish to read", title: "Designing Data-Intensive Applications", done: true},
		{name: "invalid current page position", title: "Designing Data-Intensive Applications", current: 9999, err: "The current page position (9999) can be not larger than the number of page (624) of Designing Data-Intensive Applications"},
		{name: "unknown titlel", title: "Unknown book", err: "could not find a book title: Unknown book"},
	}
	initBookshelf(dbPath, t) // Using a global variable in main.go
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bss := bookShelfServer{}
			b := &gbookshelf.Book{
				Title:   tc.title,
				Page:    tc.page,
				Done:    tc.done,
				Current: tc.current,
			}
			u, err := bss.Update(context.Background(), b)
			fmt.Printf("updating... %s\n", b)
			if err != nil {
				if tc.err == "" {
					t.Fatalf("unexpected error: %v", err)
				}
				if tc.err != err.Error() {
					t.Errorf("expected error message %q; got %q", tc.err, err)
				}
				return
			}
			if u.Title != tc.title {
				t.Fatalf("updated book title should be %s; got %v", tc.title, u)
			}
			if u.Current != tc.current {
				t.Fatalf("updated current page position should be %d; got %v", tc.current, u)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tt := []struct {
		name  string
		title string
	}{
		{"list added book", "Designing Data-Intensive Applications"},
	}
	initBookshelf(dbPath, t) // Using a global variable in main.go
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bss := bookShelfServer{}
			b := &gbookshelf.Book{
				Title: tc.title,
			}
			r, err := bss.Remove(context.Background(), b)
			if err != nil {
				t.Fatalf("List should be succeeded; failed: %v", err)
			}
			if r.Title != tc.title {
				t.Fatalf("removed book title should be %s; got %v", tc.title, r)
			}
		})
	}
}
