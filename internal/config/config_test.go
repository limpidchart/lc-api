package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/config"
)

// nolint: paralleltest
func TestNewFromEnv(t *testing.T) {
	tt := []struct {
		name           string
		setEnvFuncs    []func() error
		unsetEnvFuncs  []func() error
		expectedConfig config.Config
	}{
		{
			"all_is_set",
			[]func() error{
				setEnvVar(t, "LC_API_ADDRESS", "localhost:63010"),
				setEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT", "10"),
				setEnvVar(t, "LC_API_RENDERER_ADDRESS", "localhost:63020"),
				setEnvVar(t, "LC_API_RENDERER_CONN_TIMEOUT", "44"),
				setEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT", "120"),
			},
			[]func() error{
				unsetEnvVar(t, "LC_API_ADDRESS"),
				unsetEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT"),
				unsetEnvVar(t, "LC_API_RENDERER_ADDRESS"),
				unsetEnvVar(t, "LC_API_RENDERER_CONN_TIMEOUT"),
				unsetEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT"),
			},
			config.Config{
				API: config.APIConfig{
					Address:                "localhost:63010",
					ShutdownTimeoutSeconds: 10,
				},
				Renderer: config.RendererConfig{
					Address:               "localhost:63020",
					ConnTimeoutSeconds:    44,
					RequestTimeoutSeconds: 120,
				},
			},
		},
		{
			"lc_api_is_set",
			[]func() error{
				setEnvVar(t, "LC_API_ADDRESS", "localhost:63010"),
				setEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT", "40"),
			},
			[]func() error{
				unsetEnvVar(t, "LC_API_ADDRESS"),
				unsetEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT"),
			},
			config.Config{
				API: config.APIConfig{
					Address:                "localhost:63010",
					ShutdownTimeoutSeconds: 40,
				},
				Renderer: config.RendererConfig{
					Address:               "dns:///localhost:54020",
					ConnTimeoutSeconds:    5,
					RequestTimeoutSeconds: 30,
				},
			},
		},
		{
			"lc_renderer_is_set",
			[]func() error{
				setEnvVar(t, "LC_API_RENDERER_ADDRESS", "localhost:63040"),
				setEnvVar(t, "LC_API_RENDERER_CONN_TIMEOUT", "250"),
				setEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT", "300"),
			},
			[]func() error{
				unsetEnvVar(t, "LC_API_RENDERER_ADDRESS"),
				unsetEnvVar(t, "LC_API_RENDERER_CONN_TIMEOUT"),
				unsetEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT"),
			},
			config.Config{
				API: config.APIConfig{
					Address:                "0.0.0.0:54010",
					ShutdownTimeoutSeconds: 5,
				},
				Renderer: config.RendererConfig{
					Address:               "localhost:63040",
					ConnTimeoutSeconds:    250,
					RequestTimeoutSeconds: 300,
				},
			},
		},
		{
			"nothing_is_set",
			nil,
			nil,
			config.Config{
				API: config.APIConfig{
					Address:                "0.0.0.0:54010",
					ShutdownTimeoutSeconds: 5,
				},
				Renderer: config.RendererConfig{
					Address:               "dns:///localhost:54020",
					ConnTimeoutSeconds:    5,
					RequestTimeoutSeconds: 30,
				},
			},
		},
		{
			"bad_lc_renderer_req_timeout",
			[]func() error{
				setEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT", "11"),
				setEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT", "twenty"),
			},
			[]func() error{
				unsetEnvVar(t, "LC_API_SHUTDOWN_TIMEOUT"),
				unsetEnvVar(t, "LC_API_RENDERER_REQUEST_TIMEOUT"),
			},
			config.Config{
				API: config.APIConfig{
					Address:                "0.0.0.0:54010",
					ShutdownTimeoutSeconds: 11,
				},
				Renderer: config.RendererConfig{
					Address:               "dns:///localhost:54020",
					ConnTimeoutSeconds:    5,
					RequestTimeoutSeconds: 30,
				},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for _, setEnvFunc := range tc.setEnvFuncs {
				if err := setEnvFunc(); err != nil {
					t.Fatal(err)
				}
			}

			assert.Equal(t, tc.expectedConfig, config.NewFromEnv())

			for _, unsetEnvFunc := range tc.unsetEnvFuncs {
				if err := unsetEnvFunc(); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func setEnvVar(t *testing.T, name, value string) func() error {
	t.Helper()

	return func() error {
		if err := os.Setenv(name, value); err != nil {
			return fmt.Errorf("unable to set env var: %w", err)
		}

		return nil
	}
}

func unsetEnvVar(t *testing.T, name string) func() error {
	t.Helper()

	return func() error {
		if err := os.Unsetenv(name); err != nil {
			return fmt.Errorf("unable to unset env var: %w", err)
		}

		return nil
	}
}
