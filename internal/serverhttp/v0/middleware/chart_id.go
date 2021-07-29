package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

// RequireChartID checks that chart_id parameter can be used and stores it in the context.
func RequireChartID(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawChartID := chi.URLParam(r, view.ParamChartID)
			chartID, err := validateUUID(rawChartID)
			if err != nil {
				msg := fmt.Sprintf("%s value is bad: %s", view.ParamChartID, err)
				log := log.With().Str(RequestIDLogKey, GetRequestID(r.Context())).Logger()

				log.Warn().Msg(msg)

				w.WriteHeader(http.StatusBadRequest)
				MarshalJSON(w, view.NewError(msg))

				return
			}

			ctx := context.WithValue(r.Context(), ctxChartID, chartID.String())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetChartID retrieves chart_id value from context.
func GetChartID(ctx context.Context) string {
	if chartID, ok := ctx.Value(ctxChartID).(string); ok {
		return chartID
	}

	return ""
}

func validateUUID(raw string) (uuid.UUID, error) {
	u, err := uuid.Parse(raw)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("unable to parse %s as UUID: %w", raw, err)
	}

	return u, nil
}
