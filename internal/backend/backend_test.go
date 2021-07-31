package backend_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestBackend(t *testing.T) {
	t.Parallel()

	chartRendererServer, err := testutils.NewTestingChartRendererServer(testutils.Opts{
		ChartData: nil,
		FailMsg:   "",
		Latency:   time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unable to configure testing lc-renderer server: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	go func() {
		if serveErr := chartRendererServer.Serve(ctx); serveErr != nil {
			t.Errorf("unable to start testing lc-renderer server: %s", serveErr)

			return
		}
	}()

	rendererCfg := config.RendererConfig{
		Address:               chartRendererServer.Address(),
		ConnTimeoutSeconds:    testutils.RendererConnTimeoutSecs,
		RequestTimeoutSeconds: testutils.RendererRequestTimeoutSecs,
	}

	b, err := backend.NewBackend(context.Background(), rendererCfg)
	assert.NoError(t, err)
	assert.NotEmpty(t, b.RendererClient())
	assert.True(t, b.IsHealthy())
	assert.Equal(t, rendererCfg.RequestTimeoutSeconds, int(b.RendererRequestTimeout().Seconds()))
}
