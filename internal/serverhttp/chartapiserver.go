package serverhttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
)

const (
	// GroupV0 represents routing group pattern of the v0 API version.
	GroupV0 = "/v0"

	// GroupCharts represents routing group pattern of the charts API.
	GroupCharts = "/charts"
)

// Server implements HTTP server.
type Server struct {
	httpServer      *http.Server
	log             *zerolog.Logger
	shutdownTimeout time.Duration
}

// NewServer configures a new Server.
func NewServer(log *zerolog.Logger, b backend.Backend, httpCfg config.HTTPConfig) *Server {
	return &Server{
		// nolint: exhaustivestruct
		httpServer: &http.Server{
			Addr:         httpCfg.Address,
			ReadTimeout:  time.Duration(httpCfg.ReadTimeoutSeconds) * time.Second,
			WriteTimeout: time.Duration(httpCfg.WriteTimeoutSeconds) * time.Second,
			IdleTimeout:  time.Duration(httpCfg.IdleTimeoutSeconds) * time.Second,
			Handler:      routes(log, b),
		},
		log:             log,
		shutdownTimeout: time.Duration(httpCfg.ShutdownTimeoutSeconds) * time.Second,
	}
}

// Serve start HTTP server to serve requests.
func (s *Server) Serve(ctx context.Context) error {
	serveErr := make(chan error)

	// Launch goroutine to serve HTTP requests.
	go func() {
		defer close(serveErr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			serveErr <- fmt.Errorf("unable to start lc-api HTTP server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.log.Info().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Msg("Trying to gracefully stop lc-api HTTP server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			s.log.Warn().
				Time(zerolog.TimestampFieldName, time.Now().UTC()).
				Msg("Unable to gracefully stop lc-api HTTP server, stopping it immediately")

			s.httpServer.Close()
		}

		return nil
	case err := <-serveErr:
		return err
	}
}

func routes(log *zerolog.Logger, b backend.Backend) chi.Router {
	r := chi.NewRouter()

	r.Route(GroupV0, func(r chi.Router) {
		r.Mount(GroupCharts, chart.Routes(log, b))
	})

	return r
}
