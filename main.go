package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"golang.org/x/sync/errgroup"
)

func main() {
	/* select port number with command line arguments
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	port := os.Args[1]
	*/

	p := flag.String("port", "18080", "string flag")
	flag.Parse()
	port := *p
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Failed to listen port %s: %v", port, err)
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen port %s: %v", port, err)
	}

	if err := run(context.Background(), l); err != nil {
		log.Printf("Failed to terminate server: %v", err)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		// Addr: ":18080",  <= not used
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		err := s.Serve(l)
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to close: %+v", err)
			return err
		} else {
			return nil
		}
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("Failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
