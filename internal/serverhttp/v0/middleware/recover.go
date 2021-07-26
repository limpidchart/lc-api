package middleware

import (
	"net/http"
	"runtime"

	"github.com/rs/zerolog"
)

const stackSize = 8192

// Recover catches panics and returns the Internal Server Error back to the caller.
func Recover(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					stack := make([]byte, stackSize)
					stack = stack[:runtime.Stack(stack, false)]

					log.Error().Bytes(zerolog.ErrorStackFieldName, stack).Interface(zerolog.ErrorFieldName, rec).Msg("HTTP server panicked")

					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

					return
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
