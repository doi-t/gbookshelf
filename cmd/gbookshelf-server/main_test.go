package main

import (
	"context"
	"testing"

	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
)

// TODO: figure out the best way to organize the order of unit tests that depends on shared data in database

func TestAdd(t *testing.T) {
	tt := []struct {
		name    string
		title   string
		page    int32
		current int32
	}{
		{"add a book", "Designing Data-Intensive Applications", 624, 20},
		{"add a book with default values", "zero page", 0, 0},
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

// test data in db depends on what above unit tests do. How is it supposed to be?
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

// test data in db depends on what above unit tests do. How should it supposed to be?
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
		{name: "update zero page book", title: "zero page", page: 111},
		{name: "finish to read", title: "Designing Data-Intensive Applications", done: true},
		{name: "invalid current page position", title: "Designing Data-Intensive Applications", current: 9999, err: "The current page position (9999) can be not larger than the number of page (624) of Designing Data-Intensive Applications"},
		{name: "unknown titlel", title: "Unknown book", err: "could not find a book title: Unknown book"},
	}
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
			if err != nil {
				if tc.err != err.Error() {
					t.Errorf("expected error message %q; got %q", err, tc.err)
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

// test data in db depends on what above unit tests do. How is it supposed to be?
func TestRemove(t *testing.T) {
	tt := []struct {
		name  string
		title string
	}{
		{"list added book", "Designing Data-Intensive Applications"},
	}
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
