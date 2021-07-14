package testutils

import "github.com/limpidchart/lc-api/internal/serverrest/view/v0"

type JSONChartView struct {
	*view.ChartView
}

func NewJSONVerticalBarView() *JSONChartView {
	return &JSONChartView{
		&view.ChartView{
			Kind: "vertical_bar",
			BarsValues: []*view.BarsValues{
				{
					Values: []float32{11, 22, 33},
					FillColor: &view.ChartElementColor{
						Hex: "#fff4e6",
						RGB: nil,
					},
					StrokeColor: &view.ChartElementColor{
						Hex: "#3c2f2f",
						RGB: nil,
					},
				},
			},
			PointsValues:       nil,
			ScalarValues:       nil,
			Colors:             nil,
			BarLabelVisible:    boolToPtr(true),
			BarLabelPosition:   "end_inside",
			PointVisible:       nil,
			PointType:          "",
			PointLabelVisible:  nil,
			PointLabelPosition: "",
		},
	}
}

func NewJSONLineView() *JSONChartView {
	return &JSONChartView{
		&view.ChartView{
			Kind:         "line",
			BarsValues:   nil,
			PointsValues: nil,
			ScalarValues: &view.ScalarValues{
				Values: []float32{40, 50, 60},
			},
			Colors:             nil,
			BarLabelVisible:    nil,
			BarLabelPosition:   "",
			PointVisible:       boolToPtr(true),
			PointType:          "x",
			PointLabelVisible:  boolToPtr(true),
			PointLabelPosition: "left",
		},
	}
}

func NewJSONAreaView() *JSONChartView {
	return &JSONChartView{
		&view.ChartView{
			Kind:         "area",
			BarsValues:   nil,
			PointsValues: nil,
			ScalarValues: &view.ScalarValues{
				Values: []float32{1010, 2100},
			},
			Colors:             nil,
			BarLabelVisible:    nil,
			BarLabelPosition:   "",
			PointVisible:       boolToPtr(false),
			PointType:          "x",
			PointLabelVisible:  boolToPtr(false),
			PointLabelPosition: "top",
		},
	}
}

func NewJSONHorizontalBarView() *JSONChartView {
	return &JSONChartView{
		&view.ChartView{
			Kind: "horizontal_bar",
			BarsValues: []*view.BarsValues{
				{
					Values: []float32{10, 20},
					FillColor: &view.ChartElementColor{
						Hex: "#66b2b2",
						RGB: nil,
					},
					StrokeColor: &view.ChartElementColor{
						Hex: "#004c4c",
						RGB: nil,
					},
				},
			},
			PointsValues:       nil,
			ScalarValues:       nil,
			Colors:             nil,
			BarLabelVisible:    boolToPtr(false),
			BarLabelPosition:   "end_outside",
			PointVisible:       nil,
			PointType:          "",
			PointLabelVisible:  nil,
			PointLabelPosition: "",
		},
	}
}

func (v *JSONChartView) Unembed() *view.ChartView {
	return v.ChartView
}

func (v *JSONChartView) SetScalarValues() *JSONChartView {
	v.ScalarValues = &view.ScalarValues{
		Values: []float32{10, 20},
	}

	return v
}

func (v *JSONChartView) SetPointsValues() *JSONChartView {
	v.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20}},
	}

	return v
}

func (v *JSONChartView) UnsetValues() *JSONChartView {
	v.BarsValues = nil
	v.PointsValues = nil
	v.ScalarValues = nil

	return v
}

func (v *JSONChartView) SetBadPointsCount() *JSONChartView {
	v.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20, 30}, {40}},
	}

	return v
}
