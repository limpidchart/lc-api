package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/servergrpc"
)

// Version contains lc-api version.
// It should be provided via build tags:
//   git_tag=$(git describe --tags --abbrev=0)
//   version=${git_tag#v}
//   CGO_ENABLED=0 go build -o ./bin/lc-api -ldflags="-X main.Version=${version}" -v ./cmd/lc-api/lc-api.go
//nolint: gochecknoglobals
var Version string

func main() {
	cfg := config.NewFromEnv()
	errs := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	catchSignals(ctx, cancel)

	log.Println("Starting lc-api")

	chartAPIServer, err := servergrpc.NewServer(ctx, cfg.API, cfg.Renderer)
	if err != nil {
		cancel()
		log.Fatalf("Unable to configure lc-api gRPC server: %s", err)
	}

	go func() {
		if err := chartAPIServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()

	log.Printf("Server is started, version: %s, addr: %s", Version, cfg.API.Address)

	select {
	case <-ctx.Done():
		return
	case err := <-errs:
		cancel()
		log.Fatalf("Unable to start lc-api: %s", err)
	}
}

func catchSignals(ctx context.Context, cancel context.CancelFunc) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-done:
				log.Println("Got signal, exiting")
				cancel()
			}
		}
	}()
}
