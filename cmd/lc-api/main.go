//go:generate swagger generate spec -i ../../api/tags.yaml -o ../../api/lc-api-swagger.yaml
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/renderer"
	"github.com/limpidchart/lc-api/internal/servergrpc"
	"github.com/limpidchart/lc-api/internal/serverhttp"
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
		Msg("Creating a connection to lc-renderer")

	rendererConn, err := renderer.NewConn(ctx, cfg.Renderer)
	if err != nil {
		cancel()

		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to create a connection to lc-renderer")

		os.Exit(1)
	}

	defer rendererConn.Close()

	rendererClient := render.NewChartRendererClient(rendererConn)

	startGRPCServer(ctx, cancel, &log, cfg.GRPC, rendererClient, cfg.Renderer.RequestTimeoutSeconds, errs)
	startHTTPServer(ctx, &log, cfg.HTTP, rendererClient, cfg.Renderer.RequestTimeoutSeconds, errs)

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

func startGRPCServer(ctx context.Context, cancel context.CancelFunc, log *zerolog.Logger, cfg config.GRPCConfig, rendererClient render.ChartRendererClient, rendererReqTimeout int, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("addr", cfg.Address).
		Msg("Starting gRPC server")

	gRPCServer, err := servergrpc.NewServer(log, cfg, rendererClient, rendererReqTimeout)
	if err != nil {
		cancel()

		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to configure gRPC server")

		os.Exit(1)
	}

	go func() {
		if err := gRPCServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()
}

func startHTTPServer(ctx context.Context, log *zerolog.Logger, cfg config.HTTPConfig, rendererClient render.ChartRendererClient, rendererReqTimeout int, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("addr", cfg.Address).
		Msg("Starting HTTP server")

	httpServer := serverhttp.NewServer(log, cfg, rendererClient, rendererReqTimeout)

	go func() {
		if err := httpServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()
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
