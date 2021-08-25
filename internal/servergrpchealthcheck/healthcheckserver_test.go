package servergrpchealthcheck_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/limpidchart/lc-api/internal/backend"
	"github.com/limpidchart/lc-api/internal/config"
	"github.com/limpidchart/lc-api/internal/servergrpchealthcheck"
	"github.com/limpidchart/lc-api/internal/tcputils"
)

const testingHCEnvTimeoutSecs = 5

type testingHCEnv struct {
	hcServerConn *grpc.ClientConn
}

func newTestingHC(ctx context.Context, t *testing.T, healthy bool) *testingHCEnv {
	t.Helper()

	log := zerolog.New(os.Stderr)
	b := backend.NewEmptyBackend(healthy)
	hcCfg := config.GRPCHealthCheckConfig{
		Address: tcputils.LocalhostWithRandomPort,
	}

	hcServer, err := servergrpchealthcheck.NewServer(&log, b, hcCfg)
	if err != nil {
		t.Fatalf("unable to configure testing lc-api gRPC health check server: %s", err)
	}

	go func() {
		if serveErr := hcServer.Serve(ctx); serveErr != nil {
			t.Errorf("unable to start testing lc-api gRPC health check server: %s", serveErr)

			return
		}
	}()

	hcServerConn, err := grpc.DialContext(ctx, hcServer.Address(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("unable to create connection to testing lc-api gRPC health check server: %s", err)
	}

	return &testingHCEnv{hcServerConn}
}

func TestCheck_OK(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingHCEnvTimeoutSecs)
	defer cancel()

	testingHCEnv := newTestingHC(ctx, t, true)

	hcClient := grpc_health_v1.NewHealthClient(testingHCEnv.hcServerConn)

	// nolint: exhaustivestruct
	hcReply, hcErr := hcClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})

	assert.NoError(t, hcErr)
	assert.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING, hcReply.Status)
}

func TestCheck_Err(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*testingHCEnvTimeoutSecs)
	defer cancel()

	testingHCEnv := newTestingHC(ctx, t, false)

	hcClient := grpc_health_v1.NewHealthClient(testingHCEnv.hcServerConn)

	// nolint: exhaustivestruct
	hcReply, hcErr := hcClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})

	assert.NoError(t, hcErr)
	assert.Equal(t, grpc_health_v1.HealthCheckResponse_NOT_SERVING, hcReply.Status)
}
