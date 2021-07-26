package jsontoapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/jsontoapi"
)

func TestChartViewFromJSON(t *testing.T) {
	t.Parallel()

	// nolint: govet
	tt := []struct {
		name         string
		chartView    *view.ChartView
		expectedView *render.ChartView
		expectedErr  error
	}{
		{
			"horizontal_bar",
			testutils.NewJSONHorizontalBarView().Unembed(),
			testutils.NewHorizontalBarView().SetDefaultPointBools().Unembed(),
			nil,
		},
		{
			"area",
			testutils.NewJSONAreaView().Unembed(),
			testutils.NewAreaView().SetDefaultBarBools().Unembed(),
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
			testutils.NewJSONHorizontalBarView().SetScalarValues().Unembed(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"scalar_and_points_values",
			testutils.NewJSONHorizontalBarView().UnsetValues().SetScalarValues().SetPointsValues().Unembed(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"points_and_bars_values",
			testutils.NewJSONHorizontalBarView().SetPointsValues().Unembed(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"all_values",
			testutils.NewJSONHorizontalBarView().SetScalarValues().SetPointsValues().Unembed(),
			nil,
			jsontoapi.ErrOnlyOneOfValuesKindShouldBeSpecified,
		},
		{
			"bad_points_count",
			testutils.NewJSONAreaView().SetBadPointsCount().Unembed(),
			nil,
			jsontoapi.ErrBadPointValuesCount,
		},
		{
			"no_values",
			testutils.NewJSONAreaView().UnsetValues().Unembed(),
			nil,
			jsontoapi.ErrValuesShouldBeSpecified,
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
