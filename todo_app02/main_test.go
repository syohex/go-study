package main

import (
	"context"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	res, err := http.Get("http://localhost:18080" + "/in")
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}

	defer res.Body.Close()

	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read body: %+v", err)
	}

	want := "Hello, in"
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
	cancel()

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
