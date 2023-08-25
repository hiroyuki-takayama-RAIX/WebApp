package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	//skip test
	t.Skip("refactering...")

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return run(ctx)
	})

	// send a request to the server
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	t.Logf("try request to %v", url)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("Failed to get: %+v", err)
	}

	// close rsp.Body to
	defer rsp.Body.Close()

	// extract http body as byte[]
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("Failed to read body; %v", err)
	}

	// compare want and got
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want '%v', but got '%v'", want, string(got))
	}

	cancel()
	// when ctx is canceled,
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
