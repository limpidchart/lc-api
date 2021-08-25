package chart_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
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

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/metric"
	"github.com/limpidchart/lc-api/internal/serverhttp"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/resource/chart"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/testutils"
)

const testingRendererEnvTimeoutSecs = 5

type testingRendererEnv struct {
	chartRendererServer *testutils.TestingChartRendererServer
}

type testingRendererEnvOpts struct {
	rendererFailMsg   string
	rendererChartData []byte
	rendererLatency   time.Duration
}

func newTestingRendererEnv(ctx context.Context, t *testing.T, opts testingRendererEnvOpts) *testingRendererEnv {
	t.Helper()

	chartRendererServer, err := testutils.NewTestingChartRendererServer(testutils.Opts{
		ChartData: opts.rendererChartData,
		FailMsg:   opts.rendererFailMsg,
		Latency:   opts.rendererLatency,
	})
	if err != nil {
		t.Fatalf("unable to configure testing lc-renderer server: %s", err)
	}

	go func() {
		if serveErr := chartRendererServer.Serve(ctx); serveErr != nil {
			t.Errorf("unable to start testing lc-renderer server: %s", serveErr)

			return
		}
	}()

	return &testingRendererEnv{chartRendererServer}
}

func (tre *testingRendererEnv) address() string {
	return tre.chartRendererServer.Address()
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
	chartDataEncoded := "PHN2Zz52ZXJ0aWNhbF9hbmRfbGluZTwvc3ZnPg=="

	tre := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Millisecond * 100,
	})

	b, err := backend.NewBackend(ctx, config.RendererConfig{
		Address:               tre.address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	})
	if err != nil {
		t.Fatalf("unable to configure backend: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, b, metric.NewEmptyRecorder()))
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
	assert.Equal(t, chartDataEncoded, respBody.Chart.ChartData)
}

func TestCreateChart_VerticalAndLineOKGZIP(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte(`<svg>vertical_and_line</svg>`)
	chartDataEncoded := "PHN2Zz52ZXJ0aWNhbF9hbmRfbGluZTwvc3ZnPg=="

	tre := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Millisecond * 100,
	})

	b, err := backend.NewBackend(ctx, config.RendererConfig{
		Address:               tre.address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	})
	if err != nil {
		t.Fatalf("unable to configure backend: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, b, metric.NewEmptyRecorder()))
	})

	w := httptest.NewRecorder()
	url := strings.Join([]string{serverhttp.GroupV0, serverhttp.GroupCharts}, "")

	r, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewReader(verticalAndLineChartRequest(t)))
	if err != nil {
		t.Fatalf("unable to prepare HTTP request: %s", err)
	}

	r.Header.Set("Accept-Encoding", "gzip")

	router.ServeHTTP(w, r)

	resp := w.Result()

	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "gzip", resp.Header.Get("Content-Encoding"))
	assert.Empty(t, resp.Header.Get("Content-Length"))

	type respChart struct {
		Chart *view.ChartReply `json:"chart"`
	}

	respBody := respChart{}

	respReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Fatalf("unable to create new GZIP reader: %s", err)
	}

	defer respReader.Close()

	if err := json.NewDecoder(respReader).Decode(&respBody); err != nil {
		t.Fatalf("unable to unmarshal the response body: %s", err)
	}

	assert.NotEmpty(t, respBody.Chart.RequestID)
	assert.NotEmpty(t, respBody.Chart.ChartID)
	assert.NotEmpty(t, respBody.Chart.CreatedAt)
	assert.NotEmpty(t, respBody.Chart.DeletedAt)
	assert.Equal(t, respBody.Chart.CreatedAt, respBody.Chart.DeletedAt) // equal until the storage backend is implemented
	assert.Equal(t, view.ChartStatusCreated.String(), respBody.Chart.ChartStatus)
	assert.Equal(t, chartDataEncoded, respBody.Chart.ChartData)
}

func TestCreateChart_ErrTimeout(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("wont_happen")

	tre := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})

	b, err := backend.NewBackend(ctx, config.RendererConfig{
		Address:               tre.address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	})
	if err != nil {
		t.Fatalf("unable to configure backend: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, b, metric.NewEmptyRecorder()))
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
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"error":{"message":"Renderer request timed-out"}}`+"\n", string(body))
}

func TestCreateChart_ErrNoAxes(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("bad_axes")

	tre := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})

	b, err := backend.NewBackend(ctx, config.RendererConfig{
		Address:               tre.address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	})
	if err != nil {
		t.Fatalf("unable to configure backend: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, b, metric.NewEmptyRecorder()))
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
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"error":{"message":"Unable to render a chart: unable to validate chart axes: chart axes are not specified"}}`+"\n", string(body))
}

func TestCreateChart_ErrBadJSON(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingRendererEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("bad_json")

	tre := newTestingRendererEnv(ctx, t, testingRendererEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Minute,
	})

	b, err := backend.NewBackend(ctx, config.RendererConfig{
		Address:               tre.address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	})
	if err != nil {
		t.Fatalf("unable to configure backend: %s", err)
	}

	log := zerolog.New(os.Stderr)
	router := chi.NewRouter()
	router.Route(serverhttp.GroupV0, func(router chi.Router) {
		router.Mount(serverhttp.GroupCharts, chart.Routes(&log, b, metric.NewEmptyRecorder()))
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
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"error":{"message":"Unable to decode create chart JSON: unexpected EOF"}}`+"\n", string(body))
}
