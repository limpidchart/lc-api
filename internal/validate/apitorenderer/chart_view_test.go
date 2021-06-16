package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/testutils"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func TestValidateChartViews(t *testing.T) {
	t.Parallel()

	//nolint: govet
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
			[]*render.ChartView{testutils.VerticalBarView(), testutils.LineView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testutils.VerticalBarViewWithDefaults(), testutils.LineViewWithDefaults()},
			3,
			nil,
		},
		{
			"area",
			[]*render.ChartView{testutils.AreaView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testutils.AreaViewWithDefaults()},
			2,
			nil,
		},
		{
			"horizontal_bar",
			[]*render.ChartView{testutils.HorizontalBarView()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			[]*render.ChartView{testutils.HorizontalBarViewWithDefaults()},
			0,
			nil,
		},
		{
			"scatter",
			[]*render.ChartView{testutils.ScatterView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testutils.ScatterViewWithDefaults()},
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
			[]*render.ChartView{testutils.UnspecifiedKindView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartViewKindIsUnknown,
		},
		{
			"view_without_values",
			[]*render.ChartView{testutils.HorizontalBarViewWithoutValues()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartViewValuesShouldBeSpecified,
		},
		{
			"bad_categories_count",
			[]*render.ChartView{testutils.AreaView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			1,
			apitorenderer.ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount,
		},
		{
			"bad_rgb_value",
			[]*render.ChartView{testutils.AreaViewBadRGBColor()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			2,
			apitorenderer.ErrChartElementColorRGBBadValue,
		},
		{
			"area_with_bad_scales",
			[]*render.ChartView{testutils.AreaView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForAreaViewAreBad,
		},
		{
			"horizontal_bar_with_bad_scales",
			[]*render.ChartView{testutils.HorizontalBarView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForHorizontalBarViewAreBad,
		},
		{
			"line_with_bad_scales",
			[]*render.ChartView{testutils.LineView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForLineViewAreBad,
		},
		{
			"scatter_with_bad_scales",
			[]*render.ChartView{testutils.ScatterView()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartScalesForScatterViewAreBad,
		},
		{
			"vertical_bar_with_bad_scales",
			[]*render.ChartView{testutils.VerticalBarView()},
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
				assert.ElementsMatch(t, tc.expectedChartViews, actualChartViews)
			}
			assert.Equal(t, tc.expectedErr, actualErr)
		})
	}
}
