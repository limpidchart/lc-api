package testutils

import (
	"time"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

// EmptyBackend represents a backend.Backend that doesn't have real connections.
type EmptyBackend struct {
	healthy bool
}

// NewEmptyBackend returns a new TestingBackend.
func NewEmptyBackend(healthy bool) backend.Backend {
	return &EmptyBackend{healthy}
}

func (b *EmptyBackend) Shutdown() {}

func (b *EmptyBackend) RendererClient() render.ChartRendererClient {
	return nil
}

func (b *EmptyBackend) RendererRequestTimeout() time.Duration {
	return time.Second
}

func (b *EmptyBackend) IsHealthy() bool {
	return b.healthy
}
