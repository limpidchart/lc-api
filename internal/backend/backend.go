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

// ConnSupervisor represents an entity that contains all needed backend connections,
// can report their health at can close them.
type ConnSupervisor interface {
	Shutdown()
	RendererClient() render.ChartRendererClient
	RendererRequestTimeout() time.Duration
	IsHealthy() bool
}

// Backend contains all backend connections needed for lc-api.
type Backend struct {
	rendererConn       *grpc.ClientConn
	rendererClient     render.ChartRendererClient
	rendererReqTimeout time.Duration
}

// NewBackend configures a new Backend.
func NewBackend(ctx context.Context, rendererCfg config.RendererConfig) (*Backend, error) {
	rendererConn, err := renderer.NewConn(ctx, rendererCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to lc-renderer: %w", err)
	}

	return &Backend{
		rendererConn:       rendererConn,
		rendererClient:     render.NewChartRendererClient(rendererConn),
		rendererReqTimeout: time.Duration(rendererCfg.RequestTimeoutSeconds) * time.Second,
	}, nil
}

// Shutdown closes all backend connections.
func (b *Backend) Shutdown() {
	b.rendererConn.Close()
}

// RendererClient returns configured render.ChartRendererClient.
func (b *Backend) RendererClient() render.ChartRendererClient {
	return b.rendererClient
}

// RendererRequestTimeout returns configured timeout for renderer requests.
func (b *Backend) RendererRequestTimeout() time.Duration {
	return b.rendererReqTimeout
}

// IsHealthy checks all backend connection and reports if Backend is healthy.
func (b *Backend) IsHealthy() bool {
	return b.rendererConn.GetState() == connectivity.Ready || b.rendererConn.GetState() == connectivity.Idle
}
