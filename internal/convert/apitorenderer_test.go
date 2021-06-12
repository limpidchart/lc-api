package convert_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/convert"
	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func testingLineView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_LINE,
		Values: &render.ChartView_ScalarValues{
			ScalarValues: &render.ChartViewScalarValues{
				Values: []float32{40, 50, 60},
			},
		},
		Colors: &render.ChartViewColors{
			FillColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#65c3ba",
				},
			},
			StrokeColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#009688",
				},
			},
			PointFillColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#54b2a9",
				},
			},
			PointStrokeColor: &render.ChartElementColor{
				ColorValue: &render.ChartElementColor_ColorHex{
					ColorHex: "#35a79c",
				},
			},
		},
		BarLabelVisible:    nil,
		BarLabelPosition:   0,
		PointVisible:       &wrapperspb.BoolValue{Value: true},
		PointType:          render.ChartView_X,
		PointLabelVisible:  &wrapperspb.BoolValue{Value: true},
		PointLabelPosition: render.ChartView_LEFT,
	}
}

func testingLineViewWithDefaults() *render.ChartView {
	res := testingLineView()
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
						Values: []float32{11, 22, 33},
						FillColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#fff4e6",
							},
						},
						StrokeColor: &render.ChartElementColor{
							ColorValue: &render.ChartElementColor_ColorHex{
								ColorHex: "#3c2f2f",
							},
						},
					},
				},
			},
		},
		Colors:             nil,
		BarLabelVisible:    &wrapperspb.BoolValue{Value: true},
		BarLabelPosition:   render.ChartView_END_INSIDE,
		PointVisible:       nil,
		PointType:          0,
		PointLabelVisible:  nil,
		PointLabelPosition: 0,
	}
}

func testingVerticalBarViewWithDefaults() *render.ChartView {
	res := testingVerticalBarView()
	res.Colors = &render.ChartViewColors{
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
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointType = render.ChartView_CIRCLE
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelPosition = render.ChartView_TOP

	return res
}

func testingBandChartScale() *render.ChartScale {
	return &render.ChartScale{
		Kind:       render.ChartScale_BAND,
		RangeStart: &wrapperspb.Int32Value{Value: 0},
		RangeEnd:   &wrapperspb.Int32Value{Value: 100},
		Domain: &render.ChartScale_DomainCategories{
			DomainCategories: &render.DomainCategories{Categories: []string{"A", "B", "C"}},
		},
		NoBoundariesOffset: false,
		InnerPadding:       &wrapperspb.FloatValue{Value: 0.2},
		OuterPadding:       &wrapperspb.FloatValue{Value: 0.2},
	}
}

func testingLinearChartScale() *render.ChartScale {
	return &render.ChartScale{
		Kind:       render.ChartScale_LINEAR,
		RangeStart: &wrapperspb.Int32Value{Value: 0},
		RangeEnd:   &wrapperspb.Int32Value{Value: 100},
		Domain: &render.ChartScale_DomainNumeric{
			DomainNumeric: &render.DomainNumeric{
				Start: 0,
				End:   100,
			},
		},
		NoBoundariesOffset: false,
		InnerPadding:       nil,
		OuterPadding:       nil,
	}
}

func testingLinearChartScaleWithDefaults() *render.ChartScale {
	res := testingLinearChartScale()
	res.InnerPadding = &wrapperspb.FloatValue{Value: 0.1}
	res.OuterPadding = &wrapperspb.FloatValue{Value: 0.1}

	return res
}

func testingVerticalBarAndLineChart() *render.CreateChartRequest {
	return &render.CreateChartRequest{
		Title: "Vertical and line chart",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 100},
			Height: &wrapperspb.Int32Value{Value: 200},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 10},
			MarginBottom: &wrapperspb.Int32Value{Value: 20},
			MarginLeft:   &wrapperspb.Int32Value{Value: 30},
			MarginRight:  &wrapperspb.Int32Value{Value: 40},
		},
		Axes: &render.ChartAxes{
			AxisTop:         nil,
			AxisTopLabel:    "",
			AxisBottom:      testingBandChartScale(),
			AxisBottomLabel: "Categories",
			AxisLeft:        testingLinearChartScale(),
			AxisLeftLabel:   "Values",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			testingVerticalBarView(),
			testingLineView(),
		},
	}
}

func TestCreateChartRequestToRenderChartRequest(t *testing.T) {
	t.Parallel()

	expected := &render.RenderChartRequest{
		RequestId: "",
		Title:     "Vertical and line chart",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 100},
			Height: &wrapperspb.Int32Value{Value: 200},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 10},
			MarginBottom: &wrapperspb.Int32Value{Value: 20},
			MarginLeft:   &wrapperspb.Int32Value{Value: 30},
			MarginRight:  &wrapperspb.Int32Value{Value: 40},
		},
		Axes: &render.ChartAxes{
			AxisTop:         nil,
			AxisTopLabel:    "",
			AxisBottom:      testingBandChartScale(),
			AxisBottomLabel: "Categories",
			AxisLeft:        testingLinearChartScaleWithDefaults(),
			AxisLeftLabel:   "Values",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			testingVerticalBarViewWithDefaults(),
			testingLineViewWithDefaults(),
		},
	}

	actual, err := convert.CreateChartRequestToRenderChartRequest(testingVerticalBarAndLineChart())
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
