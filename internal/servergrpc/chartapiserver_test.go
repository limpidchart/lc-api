package servergrpc_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/servergrpc"
	"github.com/limpidchart/lc-api/internal/testutils"
)

const (
	testingChartAPIEnvTimeoutSecs         = 5
	testingChartAPIEnvRendererTimeoutSecs = 2
)

type testingChartAPIEnv struct {
	chartRendererServer     *testutils.TestingChartRendererServer
	chartRendererServerConn *grpc.ClientConn
	chartAPIServer          *servergrpc.Server
	chartAPIServerConn      *grpc.ClientConn
}

type testingChartAPIEnvOpts struct {
	rendererFailMsg   string
	rendererChartData []byte
}

func newTestingChartAPIEnv(ctx context.Context, t *testing.T, opts testingChartAPIEnvOpts) (*testingChartAPIEnv, error) {
	t.Helper()

	chartRendererServer, err := testutils.NewTestingChartRendererServer(testutils.RendererServerOpts{
		ChartData: opts.rendererChartData,
		FailMsg:   opts.rendererFailMsg,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to start testing renderer server: %w", err)
	}

	go func() {
		if serveErr := chartRendererServer.Serve(); serveErr != nil {
			t.Errorf("unable to serve testing renderer server: %w", serveErr)

			return
		}
	}()

	chartRendererServerConn, err := grpc.DialContext(ctx, chartRendererServer.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to testing chart renderer server: %w", err)
	}

	listener, err := testutils.LocalListener()
	if err != nil {
		return nil, fmt.Errorf("unable to configure local listener: %w", err)
	}

	chartAPIServer := servergrpc.NewServer(servergrpc.Opts{
		GRPCServer:      grpc.NewServer(),
		Listener:        listener,
		RendererClient:  render.NewChartRendererClient(chartRendererServerConn),
		RendererTimeout: time.Second * testingChartAPIEnvRendererTimeoutSecs,
	})

	go func() {
		if serveErr := chartAPIServer.Serve(); serveErr != nil {
			t.Errorf("unable to serve testing chart API server: %s", serveErr)

			return
		}
	}()

	chartAPIServerConn, err := grpc.DialContext(ctx, chartAPIServer.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("unable to create connection to testing chart API server: %w", err)
	}

	return &testingChartAPIEnv{
		chartRendererServer:     chartRendererServer,
		chartRendererServerConn: chartRendererServerConn,
		chartAPIServer:          chartAPIServer,
		chartAPIServerConn:      chartAPIServerConn,
	}, nil
}

func (te *testingChartAPIEnv) Stop() {
	te.chartAPIServerConn.Close()
	te.chartAPIServer.GracefulStop()
	te.chartRendererServerConn.Close()
	te.chartRendererServer.GracefulStop()
}

func TestCreateChart_OK(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	chartData := []byte("chart svg")

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: chartData,
		rendererFailMsg:   "",
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	defer testingChartAPIEnv.Stop()

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)

	createChartReply, createChartErr := chartAPIClient.CreateChart(ctx, testutils.AreaCreateChartRequest())

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
			testutils.BadSizesCreateChartRequest(),
			status.Errorf(codes.InvalidArgument, "unable to validates chart sizes: chart size max width is 100000"),
		},
		{
			"bad_margins",
			testutils.BadMarginsCreateChartRequest(),
			status.Errorf(codes.InvalidArgument, "unable to validates chart margins: chart min right margin is 0"),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingChartAPIEnvTimeoutSecs)
	defer cancel()

	testingChartAPIEnv, err := newTestingChartAPIEnv(ctx, t, testingChartAPIEnvOpts{
		rendererChartData: nil,
		rendererFailMsg:   "",
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	defer testingChartAPIEnv.Stop()

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
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	defer testingChartAPIEnv.Stop()

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)

	actualReply, actualErr := chartAPIClient.CreateChart(ctx, testutils.AreaCreateChartRequest())

	assert.Equal(t, expectedErr.Error(), actualErr.Error())
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
	})
	if err != nil {
		t.Fatalf("unable to start testing chart API environment: %s", err)
	}

	defer testingChartAPIEnv.Stop()

	chartAPIClient := render.NewChartAPIClient(testingChartAPIEnv.chartAPIServerConn)

	actualReply, actualErr := chartAPIClient.GetChart(ctx, testutils.GetChartRequest(chartID))

	assert.Equal(t, expectedErr.Error(), actualErr.Error())
	assert.Empty(t, actualReply)
}
