package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/hiroyuki-takayama-RAIX/config"
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

	/* flag is not necessary because port number is define with docker-compose.yml
	p := flag.String("port", "18080", "string flag")
	flag.Parse()
	port := *p
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Failed to listen port %s: %v", port, err)
	}
	*/

	/* delete to make run() independent
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen port %s: %v", port, err)
	}
	*/

	if err := run(context.Background()); err != nil {
		log.Printf("Failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// generate Config instance
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// make net.Listener to set port number
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Printf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("https://%s", l.Addr().String())
	log.Printf("start with: %v", url)

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
