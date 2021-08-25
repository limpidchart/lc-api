package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

type CreateChartRequest struct {
	*render.CreateChartRequest
}

func NewCreateChartRequest() *CreateChartRequest {
	return &CreateChartRequest{&render.CreateChartRequest{}}
}

func (req *CreateChartRequest) Unembed() *render.CreateChartRequest {
	return req.CreateChartRequest
}

func (req *CreateChartRequest) SetTitle() *CreateChartRequest {
	req.Title = "Chart"

	return req
}

func (req *CreateChartRequest) SetSizes() *CreateChartRequest {
	// nolint: gomnd
	req.Sizes = &render.ChartSizes{
		Width:  &wrapperspb.Int32Value{Value: 1000},
		Height: &wrapperspb.Int32Value{Value: 800},
	}

	return req
}

func (req *CreateChartRequest) SetBadSizes() *CreateChartRequest {
	// nolint: gomnd
	req.Sizes = &render.ChartSizes{
		Width:  &wrapperspb.Int32Value{Value: 999999},
		Height: nil,
	}

	return req
}

func (req *CreateChartRequest) SetMargins() *CreateChartRequest {
	// nolint: gomnd
	req.Margins = &render.ChartMargins{
		MarginTop:    &wrapperspb.Int32Value{Value: 10},
		MarginBottom: &wrapperspb.Int32Value{Value: 20},
		MarginLeft:   &wrapperspb.Int32Value{Value: 30},
		MarginRight:  &wrapperspb.Int32Value{Value: 40},
	}

	return req
}

func (req *CreateChartRequest) SetBadMargins() *CreateChartRequest {
	req.Margins = &render.ChartMargins{
		MarginTop:    nil,
		MarginBottom: nil,
		MarginLeft:   nil,
		MarginRight:  &wrapperspb.Int32Value{Value: -1000},
	}

	return req
}

func (req *CreateChartRequest) SetBandBottomAxis() *CreateChartRequest {
	axes := req.Axes
	if axes == nil {
		axes = &render.ChartAxes{}
	}

	axes.AxisBottom = NewBandChartScale().Unembed()
	req.Axes = axes

	return req
}

func (req *CreateChartRequest) SetBottomAxisLabel() *CreateChartRequest {
	axes := req.Axes
	if axes == nil {
		axes = &render.ChartAxes{}
	}

	axes.AxisBottomLabel = "Bottom Axis"
	req.Axes = axes

	return req
}

func (req *CreateChartRequest) SetLinearLeftAxis() *CreateChartRequest {
	axes := req.Axes
	if axes == nil {
		axes = &render.ChartAxes{}
	}

	axes.AxisLeft = NewLinearChartScale().Unembed()
	req.Axes = axes

	return req
}

func (req *CreateChartRequest) SetLeftAxisLabel() *CreateChartRequest {
	axes := req.Axes
	if axes == nil {
		axes = &render.ChartAxes{}
	}

	axes.AxisLeftLabel = "Left Axis"
	req.Axes = axes

	return req
}

func (req *CreateChartRequest) AddView(v *render.ChartView) *CreateChartRequest {
	req.Views = append(req.Views, v)

	return req
}

func (req *CreateChartRequest) AddAreaView() *CreateChartRequest {
	req.Views = append(req.Views, NewAreaView().Unembed())

	return req
}

func (req *CreateChartRequest) AddLineView() *CreateChartRequest {
	req.Views = append(req.Views, NewLineView().Unembed())

	return req
}

func (req *CreateChartRequest) AddVerticalBarView() *CreateChartRequest {
	req.Views = append(req.Views, NewVerticalBarView().Unembed())

	return req
}

func GetChartRequest(chartID string) *render.GetChartRequest {
	return &render.GetChartRequest{
		ChartId: chartID,
	}
}
