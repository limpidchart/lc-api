package middleware_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/serverhttp/v0/middleware"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestBackendCheck_OK(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stdout)
	router := chi.NewRouter()
	router.Use(middleware.BackendCheck(&logger, testutils.NewEmptyBackend(true)))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	w := httptest.NewRecorder()

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("unable to make a test request: %s", err)
	}

	router.ServeHTTP(w, r)

	resp := w.Result()

	resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestBackendCheck_Err(t *testing.T) {
	t.Parallel()

	logger := zerolog.New(os.Stdout)
	router := chi.NewRouter()
	router.Use(middleware.BackendCheck(&logger, testutils.NewEmptyBackend(false)))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	w := httptest.NewRecorder()

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("unable to make a test request: %s", err)
	}

	router.ServeHTTP(w, r)

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %s", err)
	}

	resp.Body.Close()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, `{"error":{"message":"Service Unavailable"}}`+"\n", string(body))
}
