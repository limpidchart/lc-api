package chart_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/renderer"
	"github.com/limpidchart/lc-api/internal/serverhttp"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/testutils"
)

const testingRendererEnvTimeoutSecs = 5

type testingRendererEnv struct {
	rendererClient render.ChartRendererClient
}

type testingRendererEnvOpts struct {
	rendererFailMsg   string
	rendererChartData []byte
	rendererLatency   time.Duration
}

func newTestingRendererEnv(ctx context.Context, t *testing.T, opts testingRendererEnvOpts) (*testingRendererEnv, error) {
	t.Helper()

	chartRendererServer, err := testutils.NewTestingChartRendererServer(testutils.Opts{
		ChartData: opts.rendererChartData,
		FailMsg:   opts.rendererFailMsg,
		Latency:   opts.rendererLatency,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to configure testing lc-renderer server: %w", err)
	}

	go func() {
		if serveErr := chartRendererServer.Serve(ctx); serveErr != nil {
			t.Errorf("unable to start testing lc-renderer server: %s", serveErr)

			return
		}
	}()

	// nolint: exhaustivestruct
	cfg := config.Config{
		Renderer: config.RendererConfig{
			Address:               chartRendererServer.Address(),
			ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
			RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
		},
	}

	rendererConn, err := renderer.NewConn(context.Background(), cfg.Renderer)
	if err != nil {
		return nil, fmt.Errorf("unable to create a connection to testing lc-renderer: %w", err)
	}

	return &testingRendererEnv{render.NewChartRendererClient(rendererConn)}, nil
}

func verticalAndLineChartRequest(t *testing.T) []byte {
	t.Helper()

	reqBody := testutils.NewJSONCreateChartRequest().
		SetTitle().
		SetSizes().
		SetMargins().
		SetBandBottomAxis().
		SetLinearLeftAxis().
		AddVerticalBarView().
		AddLineView().
		Unembed()

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("unable to marshal request body: %s", err)
	}

	return reqBodyJSON
}

func noAxesChartRequest(t *testing.T) []byte {
	t.Helper()

	reqBody := testutils.NewJSONCreateChartRequest().
		SetTitle().
		SetSizes().
		SetMargins().
		AddVerticalBarView().
		AddLineView().
		Unembed()

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("unable to marshal request body: %s", err)
	}

	return reqBodyJSON
}

func TestCreateChart_VerticalAndLineOK(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte(`<svg>vertical_and_line</svg>`)

	testingRendererEnv, err := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Millisecond * 100,
	})
	if err != nil {
		t.Fatalf("unable to start testing renderer environment: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, testingRendererEnv.rendererClient, time.Second*time.Duration(testutils.RendererRequestTimeoutSecs)))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader(verticalAndLineChartRequest(t)))
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

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	type respChart struct {
		Chart *view.ChartReply `json:"chart"`
	}

	// nolint: exhaustivestruct
	respBody := respChart{}

	if err = json.Unmarshal(body, &respBody); err != nil {
		t.Fatalf("unable to unmarshal the response body: %s", err)
	}

	assert.NotEmpty(t, respBody.Chart.RequestID)
	assert.NotEmpty(t, respBody.Chart.ChartID)
	assert.NotEmpty(t, respBody.Chart.CreatedAt)
	assert.NotEmpty(t, respBody.Chart.DeletedAt)
	assert.Equal(t, respBody.Chart.CreatedAt, respBody.Chart.DeletedAt) // equal until the storage backend is implemented
	assert.Equal(t, view.ChartStatusCreated.String(), respBody.Chart.ChartStatus)
	assert.Equal(t, string(chartData), respBody.Chart.ChartData)
}

func TestCreateChart_ErrTimeout(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("wont_happen")

	testingRendererEnv, err := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})
	if err != nil {
		t.Fatalf("unable to start testing renderer environment: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, testingRendererEnv.rendererClient, time.Second*time.Duration(testutils.RendererRequestTimeoutSecs)))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader(verticalAndLineChartRequest(t)))
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

	assert.Equal(t, http.StatusRequestTimeout, resp.StatusCode)
	assert.Equal(t, testutils.EncodeToJSON(t, view.NewError("Renderer request timed-out")), string(body))
}

func TestCreateChart_ErrNoAxes(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("bad_axes")

	testingRendererEnv, err := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})
	if err != nil {
		t.Fatalf("unable to start testing renderer environment: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, testingRendererEnv.rendererClient, time.Second*time.Duration(testutils.RendererRequestTimeoutSecs)))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader(noAxesChartRequest(t)))
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
	assert.Equal(t, testutils.EncodeToJSON(t, view.NewError("Unable to render a chart: unable to validate chart axes: chart axes are not specified")), string(body))
}

func TestCreateChart_ErrBadJSON(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("bad_json")

	testingRendererEnv, err := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})
	if err != nil {
		t.Fatalf("unable to start testing renderer environment: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, testingRendererEnv.rendererClient, time.Second*time.Duration(testutils.RendererRequestTimeoutSecs)))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader([]byte(`{"chart":{}`)))
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
	assert.Equal(t, testutils.EncodeToJSON(t, view.NewError("Unable to decode create chart JSON: unexpected EOF")), string(body))
}
