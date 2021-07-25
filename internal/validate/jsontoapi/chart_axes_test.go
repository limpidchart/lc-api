package jsontoapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/jsontoapi"
)

func TestChartScaleFromJSON(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name          string
		chartScale    *view.ChartScale
		expectedScale *render.ChartScale
		expectedErr   error
	}{
		{
			"empty",
			nil,
			nil,
			nil,
		},
		{
			"band_scale",
			testutils.NewJSONBandChartScale().Unembed(),
			testutils.NewBandChartScale().Unembed(),
			nil,
		},
		{
			"linear_scale",
			testutils.NewJSONLinearChartScale().Unembed(),
			testutils.NewLinearChartScale().Unembed(),
			nil,
		},
		{
			"band_scale_with_no_boundaries_offset",
			testutils.NewJSONBandChartScale().SetNoBoundariesOffset().Unembed(),
			testutils.NewBandChartScale().SetNoBoundariesOffset().Unembed(),
			nil,
		},
		{
			"band_scale_with_two_domains",
			testutils.NewJSONBandChartScale().SetTwoDomains().Unembed(),
			nil,
			jsontoapi.ErrOnlyOneOfDomainsShouldBeSpecified,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualScale, actualErr := jsontoapi.ChartScaleFromJSON(tc.chartScale)
			assert.Equal(t, actualScale, tc.expectedScale)
			assert.Equal(t, actualErr, tc.expectedErr)
		})
	}
}
