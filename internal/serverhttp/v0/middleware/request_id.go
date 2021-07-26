package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// SetRequestID request ID and saves it into the context.
func SetRequestID(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID, err := uuid.NewRandom()
			if err != nil {
				log.Error().Err(err).Msg("unable to generate a random UUID for request ID")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

				return
			}

			ctx := context.WithValue(r.Context(), ctxRequestID, reqID.String())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetRequestID returns request ID or empty string if it's not set.
func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(ctxRequestID).(string); ok {
		return reqID
	}

	return ""
}
