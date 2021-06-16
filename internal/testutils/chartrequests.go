package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func VerticalBarAndLineCreateChartRequest() *render.CreateChartRequest {
	//nolint: gomnd
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
			AxisBottom:      BandChartScale(),
			AxisBottomLabel: "Categories",
			AxisLeft:        LinearChartScale(),
			AxisLeftLabel:   "Values",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			VerticalBarView(),
			LineView(),
		},
	}
}
