package metric

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/config"
)

// groupMetrics represents metrics routing group pattern.
const groupMetrics = "/metrics"

// Server implements HTTP metric server.
type Server struct {
	httpServer      *http.Server
	log             *zerolog.Logger
	shutdownTimeout time.Duration
}

// NewServer configures a new Server.
func NewServer(log *zerolog.Logger, metricCfg config.MetricsConfig, pRec PromRecorder) (*Server, error) {
	return &Server{
		// nolint: exhaustivestruct
		httpServer: &http.Server{
			Addr:         metricCfg.Address,
			ReadTimeout:  time.Duration(metricCfg.ReadTimeoutSeconds) * time.Second,
			WriteTimeout: time.Duration(metricCfg.WriteTimeoutSeconds) * time.Second,
			IdleTimeout:  time.Duration(metricCfg.IdleTimeoutSeconds) * time.Second,
			Handler:      routes(pRec),
		},
		log:             log,
		shutdownTimeout: time.Duration(metricCfg.ShutdownTimeoutSeconds) * time.Second,
	}, nil
}

// Serve start HTTP server to serve requests.
func (s *Server) Serve(ctx context.Context) error {
	serveErr := make(chan error)

	// Launch goroutine to serve HTTP requests.
	go func() {
		defer close(serveErr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			serveErr <- fmt.Errorf("unable to start metrics HTTP server: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.log.Info().
			Time(zerolog.TimestampFieldName, time.Now().UTC()).
			Msg("Trying to gracefully stop metrics HTTP server")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
		defer cancel()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			s.log.Warn().
				Time(zerolog.TimestampFieldName, time.Now().UTC()).
				Msg("Unable to gracefully stop metrics HTTP server, stopping it immediately")

			s.httpServer.Close()
		}

		return nil
	case err := <-serveErr:
		return err
	}
}

func routes(pRec PromRecorder) http.Handler {
	r := chi.NewRouter()

	r.Get(groupMetrics, pRec.HTTPHandler().ServeHTTP)

	return r
}
