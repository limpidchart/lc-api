package servergrpc_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/servergrpc"
	"github.com/limpidchart/lc-api/internal/tcputils"
	"github.com/limpidchart/lc-api/internal/testutils"
)

const (
	testingChartAPIEnvTimeoutSecs  = 5
	testingChartAPIEnvShutdownSecs = 1

	testingChartAPIEnvRendererConnTimeoutSecs = 1
	testingChartAPIEnvRendererReqTimeoutSecs  = 1
)

type testingChartAPIEnv struct {
	chartAPIServerConn *grpc.ClientConn
}

type testingChartAPIEnvOpts struct {
	rendererFailMsg   string
	rendererChartData []byte
	rendererLatency   time.Duration
}

func newTestingChartAPIEnv(ctx context.Context, t *testing.T, opts testingChartAPIEnvOpts) (*testingChartAPIEnv, error) {
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

	log := zerolog.New(os.Stderr)

	chartAPIServer, err := servergrpc.NewServer(
		ctx,
		&log,
		config.GRPCConfig{
			Address:                tcputils.LocalhostWithRandomPort,
			ShutdownTimeoutSeconds: testingChartAPIEnvShutdownSecs,
		},
		config.RendererConfig{
			Address:               chartRendererServer.Address(),
			ConnTimeoutSeconds:    testingChartAPIEnvRendererConnTimeoutSecs,
			RequestTimeoutSeconds: testingChartAPIEnvRendererReqTimeoutSecs,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to configure testing lc-api server: %w", err)
	}

	go func() {
		if serveErr := chartAPIServer.Serve(ctx); serveErr != nil {
			t.Errorf("unable to serve testing chart API server: %s", serveErr)

			return
		}
	}()

	chartAPIServerConn, err := grpc.DialContext(ctx, chartAPIServer.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to testing lc-api server: %w", err)
	}

	return &testingChartAPIEnv{
		chartAPIServerConn: chartAPIServerConn,
	}, nil
}

func TestCreateChart_OK(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("chart svg")

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
		rendererLatency:   time.Millisecond * 100,
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)
	req := testutils.NewCreateChartRequest().
		SetSizes().
		SetBandBottomAxis().
		SetLinearLeftAxis().
		AddAreaView().
		Unembed()

	createChartReply, createChartErr := chartAPIClient.CreateChart(ctx, req)

	assert.NoError(t, createChartErr)
	assert.NotEmpty(t, createChartReply.RequestId)
	assert.NotEmpty(t, createChartReply.ChartId)
	assert.Equal(t, render.ChartStatus_CREATED, createChartReply.ChartStatus)
	assert.NotEmpty(t, createChartReply.CreatedAt)
	assert.NotEmpty(t, createChartReply.DeletedAt)
	assert.Equal(t, createChartReply.CreatedAt, createChartReply.DeletedAt) // should be equal when storage is not used
	assert.Equal(t, chartData, createChartReply.ChartData)
}

// nolint: paralleltest
func TestCreateChart_ConvertErrs(t *testing.T) {
	//nolint: govet
	tt := []struct {
		name        string
		request     *render.CreateChartRequest
		expectedErr error
	}{
		{
			"bad_sizes",
			testutils.NewCreateChartRequest().SetBadSizes().Unembed(),
			status.Errorf(codes.InvalidArgument, "unable to validate chart sizes: chart size max width is 100000"),
		},
		{
			"bad_margins",
			testutils.NewCreateChartRequest().SetBadMargins().Unembed(),
			status.Errorf(codes.InvalidArgument, "unable to validate chart margins: chart min right margin is 0"),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: nil,
		rendererFailMsg:   "",
		rendererLatency:   time.Millisecond * 200,
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualReply, actualErr := chartAPIClient.CreateChart(ctx, tc.request)
			assert.Equal(t, tc.expectedErr.Error(), actualErr.Error())
			assert.Empty(t, actualReply)
		})
	}
}

func TestCreateChart_RendererFailed(t *testing.T) {
	t.Parallel()

	errMsg := "unable to render your stuff!"
	expectedErr := status.Errorf(codes.InvalidArgument, "rpc error: code = InvalidArgument desc = %s", errMsg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: nil,
		rendererFailMsg:   errMsg,
		rendererLatency:   time.Millisecond * 500,
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)
	req := testutils.NewCreateChartRequest().
		SetSizes().
		SetBandBottomAxis().
		SetLinearLeftAxis().
		AddAreaView().
		Unembed()

	actualReply, actualErr := chartAPIClient.CreateChart(ctx, req)

	assert.Equal(t, expectedErr.Error(), actualErr.Error())
	assert.Empty(t, actualReply)
}

func TestCreateChart_RendererTooLong(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: nil,
		rendererFailMsg:   "",
		rendererLatency:   time.Hour,
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)
	req := testutils.NewCreateChartRequest().
		SetSizes().
		SetBandBottomAxis().
		SetLinearLeftAxis().
		AddLineView().
		Unembed()

	actualReply, actualErr := chartAPIClient.CreateChart(ctx, req)

	assert.Contains(t, actualErr.Error(), "code = InvalidArgument desc = rpc error")
	assert.Empty(t, actualReply)
}

func TestGetChart_NotFound(t *testing.T) {
	t.Parallel()

	chartID := "uuid_1"
	errMsg := "chart %s is not found"
	expectedErr := status.Errorf(codes.NotFound, errMsg, chartID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: nil,
		rendererFailMsg:   "",
		rendererLatency:   time.Second,
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)

	actualReply, actualErr := chartAPIClient.GetChart(ctx, testutils.GetChartRequest(chartID))

	assert.Equal(t, expectedErr.Error(), actualErr.Error())
	assert.Empty(t, actualReply)
}
