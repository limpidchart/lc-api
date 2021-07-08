package jsontoapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/jsontoapi"
)

func TestChartViewFromJSON(t *testing.T) {
	t.Parallel()

	//nolint: govet
	tt := []struct {
		name         string
		chartView    *view.ChartView
		expectedView *render.ChartView
		expectedErr  error
	}{
		{
			"horizontal_bar",
			testutils.JSONHorizontalBarView(),
			testutils.HorizontalBarViewWithBoolDefaults(),
			nil,
		},
		{
			"area",
			testutils.JSONAreaView(),
			testutils.AreaViewWithBoolDefaults(),
			nil,
		},
		{
			"empty",
			nil,
			nil,
			jsontoapi.ErrViewIsEmpty,
		},
		{
			"scalar_and_bars_values",
			testutils.JSONHorizontalBarViewWithScalarAndBarsValues(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"scalar_and_points_values",
			testutils.JSONHorizontalBarViewWithScalarAndPointsValues(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"points_and_bars_values",
			testutils.JSONHorizontalBarViewWithPointsAndBarsValues(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"all_values",
			testutils.JSONHorizontalBarViewWithAllValues(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"bad_points_count",
			testutils.JSONAreaViewBadPointsCount(),
			nil,
			jsontoapi.ErrBadPointValuesCount,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actualView, actualErr := jsontoapi.ChartViewFromJSON(tc.chartView)
			assert.Equal(t, tc.expectedView, actualView)
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
