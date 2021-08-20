package middleware

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

// BackendCheck checks if backend is healthy and returns http.StatusServiceUnavailable if it's not.
func BackendCheck(log *zerolog.Logger, b backend.Backend) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !b.IsHealthy() {
				log.Error().Msg("Backend connections are not healthy")

				MarshalJSON(w, http.StatusServiceUnavailable, view.NewError(http.StatusText(http.StatusServiceUnavailable)))

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
