package middleware_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/serverhttp/v0/middleware"
)

// nolint: paralleltest
func TestRecoverer(t *testing.T) {
	logFile, err := ioutil.TempFile(os.TempDir(), "test_recovery.log")
	if err != nil {
		t.Fatalf("unable to create temp log file: %s", err)
	}

	defer os.Remove(logFile.Name())
	defer logFile.Close()

	logger := zerolog.New(logFile)
	router := chi.NewRouter()
	router.Use(middleware.Recover(&logger))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		panic("some nil point reference or whatever")
	})

	w := httptest.NewRecorder()

	r, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("unable to make a test request: %s", err)
	}

	router.ServeHTTP(w, r)

	expectedMessages := []string{
		`"stack":"goroutine`,
		`"error":"some nil point reference or whatever"`,
		`"message":"HTTP server panicked"`,
	}

	contents, err := ioutil.ReadFile(logFile.Name())
	if err != nil {
		t.Logf("unable to read temp log file: %s", err)
	}

	for _, expectedMsg := range expectedMessages {
		msgInLog := strings.Contains(string(contents), expectedMsg)
		assert.True(t, msgInLog)
	}
}
