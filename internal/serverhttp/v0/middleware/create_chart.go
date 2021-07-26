package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

// RequireCreateChartParams checks if provided create chart parameters body can be used and stores it in the context.
func RequireCreateChartParams() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// nolint: exhaustivestruct
			createOptsJSON := view.CreateChartRequest{}

			err := json.NewDecoder(r.Body).Decode(&createOptsJSON)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				MarshalJSON(w, view.NewError("Create chart body is not a valid JSON"))

				return
			}

			createChartRequest, err := convert.JSONToCreateChartRequest(&createOptsJSON)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				MarshalJSON(w, view.NewError(fmt.Sprintf("Unable to use the provided create chart parameters: %s", err)))

				return
			}

			ctx := context.WithValue(r.Context(), ctxCreateChartRequest, createChartRequest)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetCreateChartRequest retrieves create chart request from context.
func GetCreateChartRequest(ctx context.Context) *render.CreateChartRequest {
	v, ok := ctx.Value(ctxCreateChartRequest).(*render.CreateChartRequest)
	if !ok {
		return nil
	}

	return v
}
