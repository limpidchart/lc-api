package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/servergrpc"
)

// Version contains lc-api version.
// It should be provided via build tags:
//   git_tag=$(git describe --tags --abbrev=0)
//   version=${git_tag#v}
//   CGO_ENABLED=0 go build -o ./bin/lc-api -ldflags="-X main.Version=${version}" -v ./cmd/lc-api/main.go
//nolint: gochecknoglobals
var Version string

func main() {
	cfg := config.NewFromEnv()
	errs := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	log := zerolog.New(os.Stderr)

	catchSignals(ctx, &log, cancel)

	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Msg("Starting lc-api")

	chartAPIServer, err := servergrpc.NewServer(ctx, &log, cfg.GRPC, cfg.Renderer)
	if err != nil {
		cancel()

		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to configure lc-api gRPC server")

		os.Exit(1)
	}

	go func() {
		if err := chartAPIServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()

	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("addr", cfg.GRPC.Address).
		Msg("Server is started")

	select {
	case <-ctx.Done():
		return
	case err := <-errs:
		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to start lc-api")

		cancel()
	}
}

func catchSignals(ctx context.Context, log *zerolog.Logger, cancel context.CancelFunc) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-done:
				log.Info().
					Time(zerolog.TimestampFieldName, time.Now().UTC()).
					Msg("Got signal, exiting")

				cancel()
			}
		}
	}()
}
