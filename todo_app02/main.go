package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: "localhost:18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}

		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}

func main() {
	err := http.ListenAndServe(":18080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	}))
	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
