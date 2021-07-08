package jsontoapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
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
			testutils.JSONBandChartScale(),
			testutils.BandChartScale(),
			nil,
		},
		{
			"linear_scale",
			testutils.JSONLinearChartScale(),
			testutils.LinearChartScale(),
			nil,
		},
		{
			"band_scale_with_no_boundaries_offset",
			testutils.JSONBandChartScaleWithNoBoundariesOffset(),
			testutils.BandChartScaleWithNoBoundariesOffset(),
			nil,
		},
		{
			"band_scale_with_two_domains",
			testutils.JSONBandChartScaleTwoDomains(),
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
