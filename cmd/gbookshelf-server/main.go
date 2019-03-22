package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
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

	projectID               = os.Getenv("PROJECT_ID")
	optFirestoreCredentials = option.WithCredentialsFile(os.Getenv("FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH"))
	bookshelfCollection     = os.Getenv("GBOOKSHELF_BOOKSHELF")
	gbookshelfServerPort    = os.Getenv("GBOOKSHELF_SERVER_PORT")
	gbookshelfMetricsPort   = os.Getenv("GBOOKSHELF_METEICS_PORT")
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
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%v", gbookshelfMetricsPort)}

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
	l, err := net.Listen("tcp", ":"+gbookshelfServerPort)
	if err != nil {
		log.Fatalf("could not listen to :%v: %v", gbookshelfServerPort, err)
	}
	log.Fatal(srv.Serve(l))
}

func (bookShelfServer) List(ctx context.Context, void *gbookshelf.Void) (*gbookshelf.Books, error) {
	// Initialize Firestore client
	var client *firestore.Client
	client, err := initFirestoreClinet(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	bookshelf := client.Collection(bookshelfCollection)
	docs := bookshelf.Documents(ctx)
	defer docs.Stop()

	var books gbookshelf.Books
	bs, err := docs.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not get all books in bookshelf: %v", err)
	}

	var book *gbookshelf.Book
	for _, b := range bs {
		book = convertBookDocToMsg(b)
		books.Books = append(books.Books, book)
	}

	return &books, nil
}

func (bookShelfServer) Add(ctx context.Context, book *gbookshelf.Book) (*gbookshelf.Book, error) {
	// Initialize Firestore client
	var client *firestore.Client
	client, err := initFirestoreClinet(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Create a book document
	wRes, err := client.Doc(bookshelfCollection+"/"+book.Title).Create(ctx, map[string]interface{}{
		"title":   book.Title,
		"page":    book.Page,
		"done":    book.Done,
		"current": book.Current,
	})
	if err != nil {
		log.Fatalf("could not add a book: %v", err)
	}
	log.Printf("New book '%s' added successfully: %v", book.Title, wRes)
	promBookUpdateCounterMetric.WithLabelValues(book.Title).Inc()
	promCurrentPageGaugeMetric.WithLabelValues(book.Title).Set(float64(book.Current))

	return book, nil
}

func (bss bookShelfServer) Remove(ctx context.Context, b *gbookshelf.Book) (*gbookshelf.Book, error) {
	// Initialize Firestore client
	var client *firestore.Client
	client, err := initFirestoreClinet(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if _, err = client.Doc(bookshelfCollection + "/" + b.Title).Delete(ctx); err != nil {
		return nil, fmt.Errorf("cloud not remove a book '%s': %v", b.Title, err)
	}

	return b, nil
}

func (bss bookShelfServer) Update(ctx context.Context, b *gbookshelf.Book) (*gbookshelf.Book, error) {
	// Initialize Firestore client
	var client *firestore.Client
	client, err := initFirestoreClinet(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ny := client.Doc(bookshelfCollection + "/" + b.Title)
	err = client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ny)
		if err != nil {
			return err
		}

		var ud *gbookshelf.Book
		ud, err = validateBookStatus(convertBookDocToMsg(doc), b)
		if err != nil {
			return err
		}

		// Update a book document
		return tx.Update(ny, []firebase.Update{
			{Path: "title", Value: ud.Title},
			{Path: "page", Value: ud.Page},
			{Path: "done", Value: ud.Done},
			{Path: "current", Value: ud.Current},
		})
	})
	if err != nil {
		// TODO: Handle error.
	}
	if err != nil {
		return nil, fmt.Errorf("could not update a book '%s': %v", b.Title, err)
	}
	log.Printf("The book '%s' updated successfully", b.Title)

	promCurrentPageGaugeMetric.WithLabelValues(b.Title).Set(float64(b.Current))
	promBookUpdateCounterMetric.WithLabelValues(b.Title).Inc()

	return b, nil
}

func initFirestoreClinet(ctx context.Context) (*firestore.Client, error) {
	// TODO: make sure if it is necessary to Initialize client for each operation or not
	client, err := firebase.NewClient(ctx, projectID, optFirestoreCredentials)
	if err != nil {
		log.Printf("cloud not initialize new Firestore app: %v", err)
		return nil, fmt.Errorf("cloud not initialize new Firestore app: %v", err)
	}
	return client, nil
}

func convertBookDocToMsg(d *firestore.DocumentSnapshot) *gbookshelf.Book {
	title := d.Data()["title"].(string)
	page := d.Data()["page"].(int64)
	done := d.Data()["done"].(bool)
	current := d.Data()["current"].(int64)

	var book *gbookshelf.Book
	book = &gbookshelf.Book{Title: title,
		Page:    int32(page),
		Done:    done,
		Current: int32(current),
	}

	return book

}

func validateBookStatus(d *gbookshelf.Book, b *gbookshelf.Book) (*gbookshelf.Book, error) {
	// Keep page in DB if given page is default value (0)
	var p int32
	if b.Page == 0 {
		p = d.Page
	} else {
		p = b.Page
	}

	// Keep current page position in DB if given current page position is default value (0)
	var c int32
	if b.Current == 0 {
		c = d.Current
	} else {
		c = b.Current
	}

	if c > p {
		return nil, fmt.Errorf("The current page position (%d) can be not larger than the number of page (%d) of %s", c, p, d.Title)
	}

	var book *gbookshelf.Book
	book = &gbookshelf.Book{
		Title:   d.Title, // must equals to b.Title
		Page:    p,
		Done:    b.Done, // use given status
		Current: c,
	}
	log.Printf("Update %v-> %v\n", d, book)

	return book, nil
}
