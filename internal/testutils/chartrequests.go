package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
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

func JSONVerticalBarAndLineCreateChartRequest() *view.CreateChartRequest {
	//nolint: gomnd
	return &view.CreateChartRequest{
		Request: struct {
			Title   string             `json:"title"`
			Sizes   *view.ChartSizes   `json:"sizes"`
			Margins *view.ChartMargins `json:"margins"`
			Axes    *view.ChartAxes    `json:"axes"`
			Views   []*view.ChartView  `json:"views"`
		}{
			Title: "Vertical and line chart",
			Sizes: &view.ChartSizes{
				Width:  intToPtr(100),
				Height: intToPtr(200),
			},
			Margins: &view.ChartMargins{
				MarginTop:    intToPtr(10),
				MarginBottom: intToPtr(20),
				MarginLeft:   intToPtr(30),
				MarginRight:  intToPtr(40),
			},
			Axes: &view.ChartAxes{
				AxisTop:         nil,
				AxisTopLabel:    "",
				AxisBottom:      JSONBandChartScale(),
				AxisBottomLabel: "Categories",
				AxisLeft:        JSONLinearChartScale(),
				AxisLeftLabel:   "Values",
				AxisRight:       nil,
				AxisRightLabel:  "",
			},
			Views: []*view.ChartView{
				JSONVerticalBarView(),
				JSONLineView(),
			},
		},
	}
}

func AreaCreateChartRequest() *render.CreateChartRequest {
	//nolint: gomnd
	return &render.CreateChartRequest{
		Title: "Area chart",
		Sizes: &render.ChartSizes{
			Width:  &wrapperspb.Int32Value{Value: 1000},
			Height: &wrapperspb.Int32Value{Value: 800},
		},
		Margins: &render.ChartMargins{
			MarginTop:    &wrapperspb.Int32Value{Value: 70},
			MarginBottom: &wrapperspb.Int32Value{Value: 40},
			MarginLeft:   &wrapperspb.Int32Value{Value: 50},
			MarginRight:  &wrapperspb.Int32Value{Value: 30},
		},
		Axes: &render.ChartAxes{
			AxisTop:         nil,
			AxisTopLabel:    "",
			AxisBottom:      BandChartScale(),
			AxisBottomLabel: "",
			AxisLeft:        LinearChartScale(),
			AxisLeftLabel:   "",
			AxisRight:       nil,
			AxisRightLabel:  "",
		},
		Views: []*render.ChartView{
			AreaView(),
		},
	}
}

func BadSizesCreateChartRequest() *render.CreateChartRequest {
	res := AreaCreateChartRequest()

	//nolint: gomnd
	res.Sizes = &render.ChartSizes{
		Width:  &wrapperspb.Int32Value{Value: 999999},
		Height: nil,
	}

	return res
}

func BadMarginsCreateChartRequest() *render.CreateChartRequest {
	res := AreaCreateChartRequest()
	res.Margins = &render.ChartMargins{
		MarginTop:    nil,
		MarginBottom: nil,
		MarginLeft:   nil,
		MarginRight:  &wrapperspb.Int32Value{Value: -1000},
	}

	return res
}

func GetChartRequest(chartID string) *render.GetChartRequest {
	return &render.GetChartRequest{
		ChartId: chartID,
	}
}
