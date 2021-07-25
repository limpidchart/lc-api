package testutils

import "github.com/limpidchart/lc-api/internal/serverhttp/v0/view"

type JSONCreateChartRequest struct {
	*view.CreateChartRequest
}

func NewJSONCreateChartRequest() *JSONCreateChartRequest {
	// nolint: exhaustivestruct
	return &JSONCreateChartRequest{&view.CreateChartRequest{}}
}

func (req *JSONCreateChartRequest) Unembed() *view.CreateChartRequest {
	return req.CreateChartRequest
}

func (req *JSONCreateChartRequest) SetTitle() *JSONCreateChartRequest {
	req.Chart.Title = "Chart"

	return req
}

func (req *JSONCreateChartRequest) SetSizes() *JSONCreateChartRequest {
	// nolint: gomnd
	req.Chart.Sizes = &view.ChartSizes{
		Width:  intToPtr(100),
		Height: intToPtr(200),
	}

	return req
}

func (req *JSONCreateChartRequest) SetMargins() *JSONCreateChartRequest {
	// nolint: gomnd
	req.Chart.Margins = &view.ChartMargins{
		MarginTop:    intToPtr(10),
		MarginBottom: intToPtr(20),
		MarginLeft:   intToPtr(30),
		MarginRight:  intToPtr(40),
	}

	return req
}

func (req *JSONCreateChartRequest) SetBandBottomAxis() *JSONCreateChartRequest {
	axes := req.Chart.Axes
	if axes == nil {
		// nolint: exhaustivestruct
		axes = &view.ChartAxes{}
	}

	axes.AxisBottom = NewJSONBandChartScale().Unembed()
	req.Chart.Axes = axes

	return req
}

func (req *JSONCreateChartRequest) SetBottomAxisLabel() *JSONCreateChartRequest {
	axes := req.Chart.Axes
	if axes == nil {
		// nolint: exhaustivestruct
		axes = &view.ChartAxes{}
	}

	axes.AxisBottomLabel = "Bottom Axis"
	req.Chart.Axes = axes

	return req
}

func (req *JSONCreateChartRequest) SetLinearLeftAxis() *JSONCreateChartRequest {
	axes := req.Chart.Axes
	if axes == nil {
		// nolint: exhaustivestruct
		axes = &view.ChartAxes{}
	}

	axes.AxisLeft = NewJSONLinearChartScale().Unembed()
	req.Chart.Axes = axes

	return req
}

func (req *JSONCreateChartRequest) SetLeftAxisLabel() *JSONCreateChartRequest {
	axes := req.Chart.Axes
	if axes == nil {
		// nolint: exhaustivestruct
		axes = &view.ChartAxes{}
	}

	axes.AxisLeftLabel = "Left Axis"
	req.Chart.Axes = axes

	return req
}

func (req *JSONCreateChartRequest) AddLineView() *JSONCreateChartRequest {
	req.Chart.Views = append(req.Chart.Views, NewJSONLineView().Unembed())

	return req
}

func (req *JSONCreateChartRequest) AddVerticalBarView() *JSONCreateChartRequest {
	req.Chart.Views = append(req.Chart.Views, NewJSONVerticalBarView().Unembed())

	return req
}
