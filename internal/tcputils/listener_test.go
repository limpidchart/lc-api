package tcputils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/tcputils"
)

func TestListener(t *testing.T) {
	t.Parallel()

	tcpListener, err := tcputils.Listener("127.0.0.1:0")

	assert.NotEmpty(t, tcpListener)
	assert.NoError(t, err)
}

func TestLocalListenerWithRandomPort(t *testing.T) {
	t.Parallel()

	tcpListener, err := tcputils.LocalListenerWithRandomPort()

	assert.NotEmpty(t, tcpListener)
	assert.NoError(t, err)
}
