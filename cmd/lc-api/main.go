//go:generate swagger generate spec -i ../../api/tags.yaml -o ../../api/lc-api-swagger.yaml
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/metric"
	"github.com/limpidchart/lc-api/internal/servergrpc"
	"github.com/limpidchart/lc-api/internal/servergrpchealthcheck"
	"github.com/limpidchart/lc-api/internal/serverhttp"
)

// Version contains lc-api version.
// It should be provided via build tags:
//   git_tag=$(git describe --tags --abbrev=0)
//   version=${git_tag#v}
//   CGO_ENABLED=0 go build -o ./bin/lc-api -ldflags="-X main.Version=${version}" -v ./cmd/lc-api/main.go
// nolint: gochecknoglobals
var Version string

func main() {
	cfg := config.NewFromEnv()
	errs := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	zerolog.DurationFieldUnit = time.Second
	log := zerolog.New(os.Stderr)

	catchSignals(ctx, &log, cancel)

	mrec, err := metric.NewRecorder()
	if err != nil {
		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to configure metric recorder")

		os.Exit(1)
	}

	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Msg("Initializing backend connections")

	b, err := backend.NewBackend(ctx, cfg.Renderer)
	if err != nil {
		cancel()

		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to initialize backend connections")

		os.Exit(1)
	}

	defer b.Shutdown()

	startMetricsServer(ctx, &log, cfg.Metrics, mrec, errs)
	startGRPCServer(ctx, cancel, &log, b, cfg.GRPC, mrec, errs)
	startHCServer(ctx, cancel, &log, b, cfg.GRPCHealthCheck, errs)
	startHTTPServer(ctx, &log, b, cfg.HTTP, mrec, errs)

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

func startGRPCServer(ctx context.Context, cancel context.CancelFunc, log *zerolog.Logger, b backend.Backend, gRPCCfg config.GRPCConfig, mrec metric.Recorder, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("address", gRPCCfg.Address).
		Msg("Starting gRPC server")

	gRPCServer, err := servergrpc.NewServer(log, b, gRPCCfg, mrec)
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

func startHCServer(ctx context.Context, cancel context.CancelFunc, log *zerolog.Logger, b backend.Backend, hcCfg config.GRPCHealthCheckConfig, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("address", hcCfg.Address).
		Msg("Starting gRPC health check server")

	hcServer, err := servergrpchealthcheck.NewServer(log, b, hcCfg)
	if err != nil {
		cancel()

		log.Error().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Err(err).
			Msg("Unable to configure gRPC health check server")

		os.Exit(1)
	}

	go func() {
		if err := hcServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()
}

func startHTTPServer(ctx context.Context, log *zerolog.Logger, b backend.Backend, httpCfg config.HTTPConfig, mrec metric.Recorder, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("address", httpCfg.Address).
		Msg("Starting HTTP server")

	httpServer, err := serverhttp.NewServer(log, b, httpCfg, mrec)
	if err != nil {
		errs <- err

		return
	}

	go func() {
		if err := httpServer.Serve(ctx); err != nil {
			errs <- err
		}
	}()
}

func startMetricsServer(ctx context.Context, log *zerolog.Logger, metricsCfg config.MetricsConfig, mrec metric.Recorder, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("address", metricsCfg.Address).
		Msg("Starting metrics server")

	metricsServer, err := metric.NewServer(log, metricsCfg, mrec)
	if err != nil {
		errs <- err

		return
	}

	go func() {
		if err := metricsServer.Serve(ctx); err != nil {
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
