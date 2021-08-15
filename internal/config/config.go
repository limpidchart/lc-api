package config

import (
	"os"
	"strconv"
)

const (
	lcRendererAddressDefault         = "dns:///localhost:54020"
	lcRendererConnTimeoutSecsDefault = 5
	lcRendererReqTimeoutSecsDefault  = 30

	gRPCAddressDefault             = "0.0.0.0:54010"
	gRPCShutdownTimeoutSecsDefault = 5

	gRPCHealthCheckAddressDefault = "0.0.0.0:54011"

	httpAddressDefault             = "0.0.0.0:54012"
	httpShutdownTimeoutSecsDefault = 5
	httpReadTimeoutSecsDefault     = 5
	httpWriteTimeoutSecsDefault    = 10
	httpIdleTimeoutSecsDefault     = 120

	metricsAddressDefault             = "0.0.0.0:54013"
	metricsShutdownTimeoutSecsDefault = 5
	metricsReadTimeoutSecsDefault     = 5
	metricsWriteTimeoutSecsDefault    = 10
	metricsIdleTimeoutSecsDefault     = 120
)

const (
	lcRendererAddressEnv         = "LC_API_RENDERER_ADDRESS"
	lcRendererConnTimeoutSecsEnv = "LC_API_RENDERER_CONN_TIMEOUT"
	lcRendererReqTimeoutSecsEnv  = "LC_API_RENDERER_REQUEST_TIMEOUT"

	gRPCAddressEnv             = "LC_API_GRPC_ADDRESS"
	gRPCShutdownTimeoutSecsEnv = "LC_API_GRPC_SHUTDOWN_TIMEOUT"

	gRPCHealthCheckAddressEnv = "LC_API_GRPC_HEALTH_CHECK_ADDRESS"

	httpAddressEnv             = "LC_API_HTTP_ADDRESS"
	httpShutdownTimeoutSecsEnv = "LC_API_HTTP_SHUTDOWN_TIMEOUT"
	httpReadTimeoutSecsEnv     = "LC_API_HTTP_READ_TIMEOUT"
	httpWriteTimeoutSecsEnv    = "LC_API_HTTP_WRITE_TIMEOUT"
	httpIdleTimeoutSecsEnv     = "LC_API_HTTP_IDLE_TIMEOUT"

	metricsAddressEnv             = "LC_METRICS_ADDRESS"
	metricsShutdownTimeoutSecsEnv = "LC_METRICS_SHUTDOWN_TIMEOUT"
	metricsReadTimeoutSecsEnv     = "LC_METRICS_READ_TIMEOUT"
	metricsWriteTimeoutSecsEnv    = "LC_METRICS_WRITE_TIMEOUT"
	metricsIdleTimeoutSecsEnv     = "LC_METRICS_IDLE_TIMEOUT"
)

// Config represents application config.
type Config struct {
	Renderer        RendererConfig
	GRPC            GRPCConfig
	GRPCHealthCheck GRPCHealthCheckConfig
	HTTP            HTTPConfig
	Metrics         MetricsConfig
}

// RendererConfig contains lc-renderer related configuration.
type RendererConfig struct {
	Address               string
	ConnTimeoutSeconds    int
	RequestTimeoutSeconds int
}

// GRPCConfig contains lc-api gRPC related configuration.
type GRPCConfig struct {
	Address                string
	ShutdownTimeoutSeconds int
}

// GRPCHealthCheckConfig contains lc-api gRPC health check related configuration.
type GRPCHealthCheckConfig struct {
	Address string
}

// HTTPConfig contains lc-api HTTP related configuration.
type HTTPConfig struct {
	Address                string
	ShutdownTimeoutSeconds int
	ReadTimeoutSeconds     int
	WriteTimeoutSeconds    int
	IdleTimeoutSeconds     int
}

// MetricsConfig contains lc-api metrics related configuration.
type MetricsConfig struct {
	Address                string
	ShutdownTimeoutSeconds int
	ReadTimeoutSeconds     int
	WriteTimeoutSeconds    int
	IdleTimeoutSeconds     int
}

// NewFromEnv creates a new Config from environment variables.
func NewFromEnv() Config {
	return Config{
		Renderer: RendererConfig{
			Address:               stringValFromEnvOrDefault(lcRendererAddressEnv, lcRendererAddressDefault),
			ConnTimeoutSeconds:    intValFromEnvOrDefault(lcRendererConnTimeoutSecsEnv, lcRendererConnTimeoutSecsDefault),
			RequestTimeoutSeconds: intValFromEnvOrDefault(lcRendererReqTimeoutSecsEnv, lcRendererReqTimeoutSecsDefault),
		},
		GRPC: GRPCConfig{
			Address:                stringValFromEnvOrDefault(gRPCAddressEnv, gRPCAddressDefault),
			ShutdownTimeoutSeconds: intValFromEnvOrDefault(gRPCShutdownTimeoutSecsEnv, gRPCShutdownTimeoutSecsDefault),
		},
		GRPCHealthCheck: GRPCHealthCheckConfig{
			Address: stringValFromEnvOrDefault(gRPCHealthCheckAddressEnv, gRPCHealthCheckAddressDefault),
		},
		HTTP: HTTPConfig{
			Address:                stringValFromEnvOrDefault(httpAddressEnv, httpAddressDefault),
			ShutdownTimeoutSeconds: intValFromEnvOrDefault(httpShutdownTimeoutSecsEnv, httpShutdownTimeoutSecsDefault),
			ReadTimeoutSeconds:     intValFromEnvOrDefault(httpReadTimeoutSecsEnv, httpReadTimeoutSecsDefault),
			WriteTimeoutSeconds:    intValFromEnvOrDefault(httpWriteTimeoutSecsEnv, httpWriteTimeoutSecsDefault),
			IdleTimeoutSeconds:     intValFromEnvOrDefault(httpIdleTimeoutSecsEnv, httpIdleTimeoutSecsDefault),
		},
		Metrics: MetricsConfig{
			Address:                stringValFromEnvOrDefault(metricsAddressEnv, metricsAddressDefault),
			ShutdownTimeoutSeconds: intValFromEnvOrDefault(metricsShutdownTimeoutSecsEnv, metricsShutdownTimeoutSecsDefault),
			ReadTimeoutSeconds:     intValFromEnvOrDefault(metricsReadTimeoutSecsEnv, metricsReadTimeoutSecsDefault),
			WriteTimeoutSeconds:    intValFromEnvOrDefault(metricsWriteTimeoutSecsEnv, metricsWriteTimeoutSecsDefault),
			IdleTimeoutSeconds:     intValFromEnvOrDefault(metricsIdleTimeoutSecsEnv, metricsIdleTimeoutSecsDefault),
		},
	}
}

func stringValFromEnvOrDefault(param, def string) string {
	val := os.Getenv(param)
	if val == "" {
		return def
	}

	return val
}

func intValFromEnvOrDefault(param string, def int) int {
	valRaw := os.Getenv(param)
	if valRaw == "" {
		return def
	}

	val, err := strconv.Atoi(valRaw)
	if err != nil {
		return def
	}

	return val
}
