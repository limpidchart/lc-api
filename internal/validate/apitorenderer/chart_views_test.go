package apitorenderer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
	"github.com/limpidchart/lc-api/internal/validate/rgb"
)

func TestValidateChartViews(t *testing.T) {
	t.Parallel()

	// nolint: govet
	tt := []struct {
		name               string
		chartViews         []*render.ChartView
		hScaleKind         render.ChartScale_ChartScaleKind
		vScaleKind         render.ChartScale_ChartScaleKind
		expectedChartViews []*render.ChartView
		categoriesCount    int
		expectedErr        error
	}{
		{
			"vertical_bar_and_line",
			[]*render.ChartView{testutils.NewVerticalBarView().Unembed(), testutils.NewLineView().Unembed()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{
				testutils.NewVerticalBarView().SetDefaultColors().SetDefaultPointParams().Unembed(),
				testutils.NewLineView().SetDefaultColors().SetDefaultBarParams().Unembed(),
			},
			3,
			nil,
		},
		{
			"area",
			[]*render.ChartView{testutils.NewAreaView().Unembed()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testutils.NewAreaView().SetDefaultColors().SetDefaultBarParams().Unembed()},
			2,
			nil,
		},
		{
			"horizontal_bar",
			[]*render.ChartView{testutils.NewHorizontalBarView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			[]*render.ChartView{testutils.NewHorizontalBarView().SetDefaultColors().SetDefaultPointParams().Unembed()},
			0,
			nil,
		},
		{
			"scatter",
			[]*render.ChartView{testutils.NewScatterView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testutils.NewScatterView().SetDefaultColors().SetDefaultBarParams().Unembed()},
			0,
			nil,
		},
		{
			"no_views",
			[]*render.ChartView{},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartViewsAreNotSpecified,
		},
		{
			"unknown_view_kind",
			[]*render.ChartView{testutils.NewHorizontalBarView().UnsetKind().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartViewKindIsUnknown,
		},
		{
			"view_without_values",
			[]*render.ChartView{testutils.NewHorizontalBarView().UnsetValues().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartViewValuesShouldBeSpecified,
		},
		{
			"bad_categories_count",
			[]*render.ChartView{testutils.NewAreaView().Unembed()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			1,
			apitorenderer.ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount,
		},
		{
			"bad_rgb_value",
			[]*render.ChartView{testutils.NewAreaView().SetBadFillRGBColor().Unembed()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			2,
			fmt.Errorf("unable to validate rgb chart element color: %w", rgb.ErrChartElementColorRGBBadValue),
		},
		{
			"area_with_bad_scales",
			[]*render.ChartView{testutils.NewAreaView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForAreaViewAreBad,
		},
		{
			"horizontal_bar_with_bad_scales",
			[]*render.ChartView{testutils.NewHorizontalBarView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForHorizontalBarViewAreBad,
		},
		{
			"line_with_bad_scales",
			[]*render.ChartView{testutils.NewLineView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForLineViewAreBad,
		},
		{
			"scatter_with_bad_scales",
			[]*render.ChartView{testutils.NewScatterView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartScalesForScatterViewAreBad,
		},
		{
			"vertical_bar_with_bad_scales",
			[]*render.ChartView{testutils.NewVerticalBarView().Unembed()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartScalesForVerticalBarViewAreBad,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actualChartViews, actualErr := apitorenderer.ValidateChartViews(tc.chartViews, tc.categoriesCount, tc.hScaleKind, tc.vScaleKind)
			if tc.expectedChartViews != nil {
				assert.Equal(t, tc.expectedChartViews, actualChartViews)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
