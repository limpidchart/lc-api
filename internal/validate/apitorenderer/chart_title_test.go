package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartTitle(t *testing.T) {
	t.Parallel()

	// nolint: govet
	tt := []struct {
		name        string
		title       string
		expectedErr error
	}{
		{
			"empty",
			"",
			nil,
		},
		{
			"small len",
			"my chart",
			nil,
		},
		{
			"too big len",
			testutils.RandomString(1025),
			apitorenderer.ErrTitleMaxLen,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualErr := apitorenderer.ValidateChartTitle(tc.title)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
