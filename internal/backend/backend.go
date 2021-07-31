package backend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"

	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/renderer"
)

// Backend contains all backend connections needed for lc-api.
type Backend interface {
	Shutdown()
	RendererClient() render.ChartRendererClient
	RendererRequestTimeout() time.Duration
	IsHealthy() bool
}

type appBackend struct {
	rendererConn       *grpc.ClientConn
	rendererClient     render.ChartRendererClient
	rendererReqTimeout time.Duration
}

// NewBackend configures a new Backend.
func NewBackend(ctx context.Context, rendererCfg config.RendererConfig) (Backend, error) {
	rendererConn, err := renderer.NewConn(ctx, rendererCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to lc-renderer: %w", err)
	}

	return &appBackend{
		rendererConn:       rendererConn,
		rendererClient:     render.NewChartRendererClient(rendererConn),
		rendererReqTimeout: time.Duration(rendererCfg.RequestTimeoutSeconds) * time.Second,
	}, nil
}

// Shutdown closes all backend connections.
func (b *appBackend) Shutdown() {
	b.rendererConn.Close()
}

// RendererClient returns configured render.ChartRendererClient.
func (b *appBackend) RendererClient() render.ChartRendererClient {
	return b.rendererClient
}

// RendererRequestTimeout returns configured timeout for renderer requests.
func (b *appBackend) RendererRequestTimeout() time.Duration {
	return b.rendererReqTimeout
}

// IsHealthy checks all backend connection and reports if Backend is healthy.
func (b *appBackend) IsHealthy() bool {
	return b.rendererConn.GetState() == connectivity.Ready || b.rendererConn.GetState() == connectivity.Idle
}
