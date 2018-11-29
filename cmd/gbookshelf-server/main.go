package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	firebase "cloud.google.com/go/firestore"
	"github.com/doi-t/gbookshelf/pkg/apis/gbookshelf"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/api/option"
	grpc "google.golang.org/grpc"
)

type bookShelfServer struct{}

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	promBookUpdateCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gbookshelf_book_update_count",
		Help: "Total number of book update.",
	}, []string{"book_title"})

	// Create a customized counter metric.
	promCurrentPageGaugeMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gbookshelf_book_current_page",
		Help: "The current page position of book.",
	}, []string{"book_title"})

	projectID      = os.Getenv("PROJECT_ID")
	optCredentials = option.WithCredentialsFile(os.Getenv("GCLOUD_CRENTIAL_FILE_PATH"))
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(
		grpcMetrics,
		promBookUpdateCounterMetric,
		promCurrentPageGaugeMetric,
	)
}

// FIXME: Graceful shutdown is missing.
func main() {
	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 2112)}

	// Create a gRPC Server with gRPC interceptor.
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)
	// Register gbookshelf-server gRPC service implementations.
	var bookshelf bookShelfServer
	gbookshelf.RegisterBookShelfServer(srv, bookshelf)

	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(srv)

	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	// Start http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Unable to start a http server for Prometheus: %v", err)
		}
	}()

	// Start gbookshelf service
	l, err := net.Listen("tcp", ":8888") // TODO: make port number environment variable
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

func (bookShelfServer) List(ctx context.Context, void *gbookshelf.Void) (*gbookshelf.Books, error) {
	// Initialize Firestore client
	client, err := firebase.NewClient(ctx, projectID, optCredentials)
	if err != nil {
		return nil, fmt.Errorf("cloud not Initialize new Firestore app: %v", err)
	}
	defer client.Close()

	bookshelf := client.Collection("bookShelf")
	docs := bookshelf.Documents(ctx)
	defer docs.Stop()

	var books gbookshelf.Books
	bs, err := docs.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not get all books in bookshelf: %v", err)
	}
	for _, b := range bs {
		// FIXME: Now: all ok of type assertion is ignored.
		// TODO: Don't want to map variables by myself
		title := b.Data()["title"].(string)
		page := b.Data()["page"].(int64)
		done := b.Data()["done"].(bool)
		current := b.Data()["current"].(int64)

		var book gbookshelf.Book
		book = gbookshelf.Book{Title: title, Page: int32(page), Done: done, Current: int32(current)}

		books.Books = append(books.Books, &book)
	}

	return &books, nil
}

func (bookShelfServer) Add(ctx context.Context, book *gbookshelf.Book) (*gbookshelf.Book, error) {
	// Initialize Firestore client
	client, err := firebase.NewClient(ctx, projectID, optCredentials)
	if err != nil {
		return nil, fmt.Errorf("cloud not Initialize new Firestore app: %v", err)
	}
	defer client.Close()

	// Create a book document
	wRes, err := client.Doc("bookShelf/"+book.Title).Create(ctx, map[string]interface{}{
		"title":   book.Title,
		"page":    book.Page,
		"done":    book.Done,
		"current": book.Current,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	log.Printf("New book '%s' added successfully: %v", book.Title, wRes)
	promBookUpdateCounterMetric.WithLabelValues(book.Title).Inc()
	promCurrentPageGaugeMetric.WithLabelValues(book.Title).Set(float64(book.Current))

	return book, nil
}

func (bss bookShelfServer) Remove(ctx context.Context, rb *gbookshelf.Book) (*gbookshelf.Book, error) {
	return nil, nil
}

func (bss bookShelfServer) Update(ctx context.Context, b *gbookshelf.Book) (*gbookshelf.Book, error) {
	// FIXME It is not necessary anymore
	l, err := bss.List(ctx, &gbookshelf.Void{})
	if err != nil {
		return nil, err
	}

	newList, err := updateBookList(l, b)
	if err != nil {
		return nil, err
	}

	for _, book := range newList.Books {
		if book.Title == b.Title {
			b = book
		}
	}

	// Initialize Firestore client
	client, err := firebase.NewClient(ctx, projectID, optCredentials)
	if err != nil {
		return nil, fmt.Errorf("cloud not Initialize new Firestore app: %v", err)
	}
	defer client.Close()

	// Update a book document
	wRes, err := client.Doc("bookShelf/"+b.Title).Update(ctx, []firebase.Update{
		{Path: "title", Value: b.Title},
		{Path: "page", Value: b.Page},
		{Path: "done", Value: b.Done},
		{Path: "current", Value: b.Current},
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	log.Printf("New book '%s' added successfully: %v", b.Title, wRes)

	promCurrentPageGaugeMetric.WithLabelValues(b.Title).Set(float64(b.Current))
	promBookUpdateCounterMetric.WithLabelValues(b.Title).Inc()

	return b, nil
}

func updateBookList(l *gbookshelf.Books, b *gbookshelf.Book) (*gbookshelf.Books, error) {
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

	return &newList, nil
}
