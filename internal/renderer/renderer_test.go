package renderer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/testutils"
)

func TestNewConn(t *testing.T) {
	chartRendererServer, err := testutils.NewTestingChartRendererServer(testutils.Opts{
		ChartData: nil,
		FailMsg:   "",
		Latency:   time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unable to start testing lc-renderer server: %s", err)
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

	chartRendererConn, err := NewConn(context.Background(), rendererCfg)
	assert.NoError(t, err)
	assert.NotEmpty(t, chartRendererConn)
}
