package apitorenderer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/validate/apitorenderer"
)

func testingColorsDefault() *render.ChartViewColors {
	return &render.ChartViewColors{
		FillColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#71c7ec",
			},
		},
		StrokeColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#005073",
			},
		},
		PointFillColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#71c7ec",
			},
		},
		PointStrokeColor: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#005073",
			},
		},
	}
}

func testingChartUnspecifiedKind() *render.ChartView {
	res := testingHorizontalBarView()
	res.Kind = render.ChartView_UNSPECIFIED_KIND

	return res
}

func testingChartNoValues() *render.ChartView {
	res := testingHorizontalBarView()
	res.Values = nil

	return res
}

func testingHorizontalBarView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_HORIZONTAL_BAR,
		Values: &render.ChartView_BarsValues{
			BarsValues: &render.ChartViewBarsValues{
				BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
					{
						Values: []float32{10, 20},
						FillColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#66b2b2",
							},
						},
						StrokeColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#004c4c",
							},
						},
					},
				},
			},
		},
		Colors:             nil,
		BarLabelVisible:    &wrapperspb.BoolValue{Value: false},
		BarLabelPosition:   render.ChartView_END_OUTSIDE,
		PointVisible:       nil,
		PointType:          0,
		PointLabelVisible:  nil,
		PointLabelPosition: 0,
	}
}

func testingHorizontalBarViewWithDefaults() *render.ChartView {
	res := testingHorizontalBarView()
	res.Colors = testingColorsDefault()
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointType = render.ChartView_CIRCLE
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelPosition = render.ChartView_TOP

	return res
}

func testingAreaView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_AREA,
		Values: &render.ChartView_ScalarValues{
			ScalarValues: &render.ChartViewScalarValues{
				Values: []float32{1010, 2100},
			},
		},
		Colors: &render.ChartViewColors{
			FillColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#71c7ee",
				},
			},
			StrokeColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#005072",
				},
			},
			PointFillColor:   nil,
			PointStrokeColor: nil,
		},
		BarLabelVisible:    nil,
		BarLabelPosition:   0,
		PointVisible:       &wrapperspb.BoolValue{Value: false},
		PointType:          render.ChartView_X,
		PointLabelVisible:  &wrapperspb.BoolValue{Value: false},
		PointLabelPosition: render.ChartView_TOP,
	}
}

func testingAreaViewWithDefaults() *render.ChartView {
	res := testingAreaView()
	res.Colors = testingColorsDefault()
	res.Colors.FillColor = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: "#71c7ee",
		},
	}
	res.Colors.StrokeColor = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: "#005072",
		},
	}
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_CENTER

	return res
}

func testingAreaViewBadRGBColor() *render.ChartView {
	res := testingAreaView()
	res.Colors = testingColorsDefault()
	res.Colors.FillColor = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorRgb{
			ColorRgb: &render.ChartElementColor_RGB{
				R: 1,
				G: 320,
				B: 2,
			},
		},
	}

	return res
}

func testingLineView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_LINE,
		Values: &render.ChartView_ScalarValues{
			ScalarValues: &render.ChartViewScalarValues{
				Values: []float32{100, 200},
			},
		},
		Colors:             nil,
		BarLabelVisible:    nil,
		BarLabelPosition:   0,
		PointVisible:       &wrapperspb.BoolValue{Value: true},
		PointType:          render.ChartView_SQUARE,
		PointLabelVisible:  &wrapperspb.BoolValue{Value: true},
		PointLabelPosition: render.ChartView_BOTTOM_RIGHT,
	}
}

func testingLineViewWithDefaults() *render.ChartView {
	res := testingLineView()
	res.Colors = testingColorsDefault()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_CENTER

	return res
}

func testingScatterView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_SCATTER,
		Values: &render.ChartView_PointsValues{
			PointsValues: &render.ChartViewPointsValues{
				Points: []*render.ChartViewPointsValues_Point{
					{
						X: 321,
						Y: 8741,
					},
					{
						X: 23,
						Y: 85,
					},
				},
			},
		},
		Colors:             nil,
		BarLabelVisible:    nil,
		BarLabelPosition:   0,
		PointVisible:       &wrapperspb.BoolValue{Value: true},
		PointType:          render.ChartView_CIRCLE,
		PointLabelVisible:  &wrapperspb.BoolValue{Value: true},
		PointLabelPosition: render.ChartView_TOP_LEFT,
	}
}

func testingScatterViewWithDefaults() *render.ChartView {
	res := testingScatterView()
	res.Colors = testingColorsDefault()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_CENTER

	return res
}

func testingVerticalBarView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_VERTICAL_BAR,
		Values: &render.ChartView_BarsValues{
			BarsValues: &render.ChartViewBarsValues{
				BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
					{
						Values: []float32{20, 30},
						FillColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#bf0000",
							},
						},
						StrokeColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#400000",
							},
						},
					},
				},
			},
		},
		Colors:             nil,
		BarLabelVisible:    &wrapperspb.BoolValue{Value: false},
		BarLabelPosition:   render.ChartView_END_OUTSIDE,
		PointVisible:       nil,
		PointType:          0,
		PointLabelVisible:  nil,
		PointLabelPosition: 0,
	}
}

func testingVerticalBarViewWithDefaults() *render.ChartView {
	res := testingVerticalBarView()
	res.Colors = testingColorsDefault()
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointType = render.ChartView_CIRCLE
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelPosition = render.ChartView_TOP

	return res
}

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
			[]*render.ChartView{testingVerticalBarView(), testingLineView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testingVerticalBarViewWithDefaults(), testingLineViewWithDefaults()},
			2,
			nil,
		},
		{
			"area",
			[]*render.ChartView{testingAreaView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testingAreaViewWithDefaults()},
			2,
			nil,
		},
		{
			"horizontal_bar",
			[]*render.ChartView{testingHorizontalBarView()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			[]*render.ChartView{testingHorizontalBarViewWithDefaults()},
			0,
			nil,
		},
		{
			"scatter",
			[]*render.ChartView{testingScatterView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			[]*render.ChartView{testingScatterViewWithDefaults()},
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
			[]*render.ChartView{testingChartUnspecifiedKind()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartViewKindIsUnknown,
		},
		{
			"view_without_values",
			[]*render.ChartView{testingChartNoValues()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartViewValuesShouldBeSpecified,
		},
		{
			"bad_categories_count",
			[]*render.ChartView{testingAreaView()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			1,
			apitorenderer.ErrChartViewValuesCountShouldBeEqualOrLessOfCategoriesCount,
		},
		{
			"bad_rgb_value",
			[]*render.ChartView{testingAreaViewBadRGBColor()},
			render.ChartScale_BAND,
			render.ChartScale_LINEAR,
			nil,
			2,
			apitorenderer.ErrChartElementColorRGBBadValue,
		},
		{
			"area_with_bad_scales",
			[]*render.ChartView{testingAreaView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForAreaViewAreBad,
		},
		{
			"horizontal_bar_with_bad_scales",
			[]*render.ChartView{testingHorizontalBarView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForHorizontalBarViewAreBad,
		},
		{
			"line_with_bad_scales",
			[]*render.ChartView{testingLineView()},
			render.ChartScale_LINEAR,
			render.ChartScale_LINEAR,
			nil,
			0,
			apitorenderer.ErrChartScalesForLineViewAreBad,
		},
		{
			"scatter_with_bad_scales",
			[]*render.ChartView{testingScatterView()},
			render.ChartScale_LINEAR,
			render.ChartScale_BAND,
			nil,
			0,
			apitorenderer.ErrChartScalesForScatterViewAreBad,
		},
		{
			"vertical_bar_with_bad_scales",
			[]*render.ChartView{testingVerticalBarView()},
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
