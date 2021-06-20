package config

import (
	"os"
	"strconv"
)

const (
	lcAPIAddressDefault             = "0.0.0.0:54010"
	lcAPIShutdownTimeoutSecsDefault = 5

	lcRendererAddressDefault         = "dns:///localhost:54020"
	lcRendererConnTimeoutSecsDefault = 5
	lcRendererReqTimeoutSecsDefault  = 30
)

const (
	lcAPIAddressEnv             = "LC_API_ADDRESS"
	lcAPIShutdownTimeoutSecsEnv = "LC_API_SHUTDOWN_TIMEOUT"

	lcRendererAddressEnv         = "LC_API_RENDERER_ADDRESS"
	lcRendererConnTimeoutSecsEnv = "LC_API_RENDERER_CONN_TIMEOUT"
	lcRendererReqTimeoutSecsEnv  = "LC_API_RENDERER_REQUEST_TIMEOUT"
)

// Config represents application config.
type Config struct {
	API      APIConfig
	Renderer RendererConfig
}

// APIConfig contains lc-api related configuration.
type APIConfig struct {
	Address                string
	ShutdownTimeoutSeconds int
}

// RendererConfig contains lc-renderer related configuration.
type RendererConfig struct {
	Address               string
	ConnTimeoutSeconds    int
	RequestTimeoutSeconds int
}

// NewFromEnv creates a new Config from environment variables.
func NewFromEnv() Config {
	return Config{
		API: APIConfig{
			Address:                stringValFromEnvOrDefault(lcAPIAddressEnv, lcAPIAddressDefault),
			ShutdownTimeoutSeconds: intValFromEnvOrDefault(lcAPIShutdownTimeoutSecsEnv, lcAPIShutdownTimeoutSecsDefault),
		},
		Renderer: RendererConfig{
			Address:               stringValFromEnvOrDefault(lcRendererAddressEnv, lcRendererAddressDefault),
			ConnTimeoutSeconds:    intValFromEnvOrDefault(lcRendererConnTimeoutSecsEnv, lcRendererConnTimeoutSecsDefault),
			RequestTimeoutSeconds: intValFromEnvOrDefault(lcRendererReqTimeoutSecsEnv, lcRendererReqTimeoutSecsDefault),
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
