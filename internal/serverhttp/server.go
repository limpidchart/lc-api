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
	"github.com/limpidchart/lc-api/internal/metric"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
)

const (
	// GroupV0 represents routing group pattern of the v0 API version.
	GroupV0 = "/v0"

	// GroupCharts represents routing group pattern of the charts API.
	GroupCharts = "/charts"
)

const name = "HTTP API"

// Server implements HTTP server.
type Server struct {
	httpServer      *http.Server
	log             *zerolog.Logger
	shutdownTimeout time.Duration
}

// NewServer configures a new Server.
func NewServer(log *zerolog.Logger, bCon backend.ConnSupervisor, httpCfg config.HTTPConfig, pRec metric.PromRecorder) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         httpCfg.Address,
			ReadTimeout:  time.Duration(httpCfg.ReadTimeoutSeconds) * time.Second,
			WriteTimeout: time.Duration(httpCfg.WriteTimeoutSeconds) * time.Second,
			IdleTimeout:  time.Duration(httpCfg.IdleTimeoutSeconds) * time.Second,
			Handler:      routes(log, bCon, pRec),
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

// Address returns server address.
func (s *Server) Address() string {
	return s.httpServer.Addr
}

// Name returns server name.
func (s *Server) Name() string {
	return name
}

func routes(log *zerolog.Logger, bCon backend.ConnSupervisor, pRec metric.PromRecorder) chi.Router {
	r := chi.NewRouter()

	r.Route(GroupV0, func(r chi.Router) {
		r.Mount(GroupCharts, chart.Routes(log, bCon, pRec))
	})

	return r
}
