//go:generate swagger generate spec -i ../../api/tags.yaml -o ../../api/lc-api-swagger.yaml
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/metric"
	"github.com/limpidchart/lc-api/internal/servergrpc"
	"github.com/limpidchart/lc-api/internal/servergrpchc"
	"github.com/limpidchart/lc-api/internal/serverhttp"
	"github.com/limpidchart/lc-api/internal/tcputils"
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

	rec, err := metric.NewRecorder()
	if err != nil {
		log.Error().Time(zerolog.TimestampFieldName, time.Now().UTC()).Err(err).Msg("Unable to configure metric recorder")
		os.Exit(1)
	}

	hcListener, err := tcputils.Listener(cfg.GRPCHealthCheck.Address)
	if err != nil {
		cancel()
		log.Error().Time(zerolog.TimestampFieldName, time.Now().UTC()).Err(err).Msg("Unable to create TCP listener for healthcheck server")
		os.Exit(1)
	}

	gRPCListener, err := tcputils.Listener(cfg.GRPC.Address)
	if err != nil {
		cancel()
		log.Error().Time(zerolog.TimestampFieldName, time.Now().UTC()).Err(err).Msg("Unable to create TCP listener for gRPC API server")
		os.Exit(1)
	}

	b, err := backend.NewBackend(ctx, cfg.Renderer)
	if err != nil {
		cancel()
		log.Error().Time(zerolog.TimestampFieldName, time.Now().UTC()).Err(err).Msg("Unable to create backend connections")
		os.Exit(1)
	}

	defer b.Shutdown()

	startServer(ctx, &log, metric.NewServer(&log, cfg.Metrics, rec), errs)
	startServer(ctx, &log, servergrpc.NewServer(&log, gRPCListener, b, cfg.GRPC, rec), errs)
	startServer(ctx, &log, serverhttp.NewServer(&log, b, cfg.HTTP, rec), errs)
	startServer(ctx, &log, servergrpchc.NewServer(&log, hcListener, b), errs)

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

type server interface {
	Name() string
	Address() string
	Serve(ctx context.Context) error
}

func startServer(ctx context.Context, log *zerolog.Logger, s server, errs chan<- error) {
	log.Info().
		Time(zerolog.TimestampFieldName, time.Now().UTC()).
		Str("version", Version).
		Str("address", s.Address()).
		Msg(fmt.Sprintf("Starting %s server", s.Name()))

	go func() {
		if err := s.Serve(ctx); err != nil {
			errs <- fmt.Errorf("unable to start %s server: %w", s.Name(), err)
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
