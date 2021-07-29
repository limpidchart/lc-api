package chart_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/serverhttp"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestGetChart_NotFound(t *testing.T) {
	t.Parallel()

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, nil, 0))
	})

	w := httptest.NewRecorder()
	chartID := testutils.RandomUUID(t).String()
	url := fmt.Sprintf("%s%s/%s", serverhttp.GroupV0, serverhttp.GroupCharts, chartID)

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

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(`{"error":{"id":"%s","message":"chart not found"}}`+"\n", chartID), string(body))
}

func TestGetChart_BadChartID(t *testing.T) {
	t.Parallel()

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, nil, 0))
	})

	w := httptest.NewRecorder()
	url := fmt.Sprintf("%s%s/%s", serverhttp.GroupV0, serverhttp.GroupCharts, "mychart")

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

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, `{"error":{"message":"chart_id value is bad: unable to parse mychart as UUID: invalid UUID length: 7"}}`+"\n", string(body))
}
