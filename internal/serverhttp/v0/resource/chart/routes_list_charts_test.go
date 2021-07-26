package chart_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/serverhttp"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestListCharts_Unimplemented(t *testing.T) {
	t.Parallel()

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, nil, 0))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("unable to prepare HTTP request: %s", err)
	}

	router.ServeHTTP(w, r)

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %s", err)
	}

	resp.Body.Close()

	assert.Equal(t, http.StatusNotImplemented, resp.StatusCode)
	assert.Equal(t, testutils.EncodeToJSON(t, view.NewError("Unable to get list of charts since auth is not implemented yet")), string(body))
}
