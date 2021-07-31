package renderer

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

const rendererServiceCfg = `{"loadBalancingPolicy":"round_robin"}`

var (
	// ErrCreateChartRequestCancelled contains error message about cancelled create chart request.
	ErrCreateChartRequestCancelled = errors.New("create chart request is cancelled")

	// ErrGenerateChartIDFailed contains error message about failed chart ID generation.
	ErrGenerateChartIDFailed = errors.New("unable to generate a random UUID for chart ID")
)

// NewConn creates a new lc-renderer connection.
func NewConn(ctx context.Context, rendererCfg config.RendererConfig) (*grpc.ClientConn, error) {
	rendererConnCtx, rendererConnCancel := context.WithTimeout(ctx, time.Second*time.Duration(rendererCfg.ConnTimeoutSeconds))
	defer rendererConnCancel()

	rendererConn, err := grpc.DialContext(
		rendererConnCtx,
		rendererCfg.Address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(rendererServiceCfg),
	)
	if err != nil {
		return nil, fmt.Errorf("lc-renderer gRPC dial failed: %w", err)
	}

	return rendererConn, nil
}

// CreateChartOpts represents options for CreateChart method.
type CreateChartOpts struct {
	RequestID      string
	Request        *render.CreateChartRequest
	RendererClient render.ChartRendererClient
	Timeout        time.Duration
}

// CreateChart converts render.CreateChartRequest and requests a chart rendering from lc-renderer.
//
// Note: tests are implemented in internal/servergrpc package.
func CreateChart(ctx context.Context, opts CreateChartOpts) (*render.ChartReply, error) {
	now := time.Now().UTC()

	renderChartReq, err := convert.CreateChartRequestToRenderChartRequest(opts.Request)
	if err != nil {
		return nil, err
	}

	renderChartReq.RequestId = opts.RequestID

	rendererCtx, rendererCancel := context.WithTimeout(ctx, opts.Timeout)
	defer rendererCancel()

	select {
	case <-ctx.Done():
		return nil, ErrCreateChartRequestCancelled
	case renderResult := <-renderChart(rendererCtx, opts.RendererClient, renderChartReq):
		switch {
		case isTimedOutErr(renderResult.err):
			return nil, ErrCreateChartRequestCancelled
		case renderResult.err != nil:
			return nil, renderResult.err
		default:
			chartID, err := uuid.NewRandom()
			if err != nil {
				return nil, ErrGenerateChartIDFailed
			}

			return convert.RenderChartReplyToAPIChartReply(opts.RequestID, chartID.String(), now, renderResult.reply), nil
		}
	}
}

type renderChartResult struct {
	reply *render.RenderChartReply
	err   error
}

func renderChart(ctx context.Context, client render.ChartRendererClient, req *render.RenderChartRequest) <-chan renderChartResult {
	result := make(chan renderChartResult)

	go func() {
		reply, err := client.RenderChart(ctx, req)
		result <- renderChartResult{reply, err}
		close(result)
	}()

	return result
}

func isTimedOutErr(err error) bool {
	if errors.Is(err, status.Error(codes.DeadlineExceeded, context.DeadlineExceeded.Error())) {
		return true
	}

	if errors.Is(err, status.Error(codes.Canceled, context.Canceled.Error())) {
		return true
	}

	return false
}
