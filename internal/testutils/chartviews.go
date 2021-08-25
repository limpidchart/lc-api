package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
)

func ColorsDefault() *render.ChartViewColors {
	return &render.ChartViewColors{
		Fill: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#71c7ec",
			},
		},
		Stroke: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#005073",
			},
		},
		PointFill: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#71c7ec",
			},
		},
		PointStroke: &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "#005073",
			},
		},
	}
}

type ChartView struct {
	*render.ChartView
}

func NewHorizontalBarView() *ChartView {
	return &ChartView{
		&render.ChartView{
			Kind: render.ChartView_HORIZONTAL_BAR,
			Values: &render.ChartView_BarsValues{
				BarsValues: &render.ChartViewBarsValues{
					BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
						{
							Values: []float32{10, 20},
							Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
								Fill: &render.ChartElementColor{
									ColorValue: &render.ChartElementColor_ColorHex{
										ColorHex: "#66b2b2",
									},
								},
								Stroke: &render.ChartElementColor{
									ColorValue: &render.ChartElementColor_ColorHex{
										ColorHex: "#004c4c",
									},
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
		},
	}
}

func NewAreaView() *ChartView {
	return &ChartView{
		&render.ChartView{
			Kind: render.ChartView_AREA,
			Values: &render.ChartView_ScalarValues{
				ScalarValues: &render.ChartViewScalarValues{
					Values: []float32{1010, 2100},
				},
			},
			Colors:             nil,
			BarLabelVisible:    nil,
			BarLabelPosition:   0,
			PointVisible:       &wrapperspb.BoolValue{Value: false},
			PointType:          render.ChartView_X,
			PointLabelVisible:  &wrapperspb.BoolValue{Value: false},
			PointLabelPosition: render.ChartView_TOP,
		},
	}
}

func NewScatterView() *ChartView {
	// nolint: gomnd
	return &ChartView{
		&render.ChartView{
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
		},
	}
}

func NewLineView() *ChartView {
	return &ChartView{
		&render.ChartView{
			Kind: render.ChartView_LINE,
			Values: &render.ChartView_ScalarValues{
				ScalarValues: &render.ChartViewScalarValues{
					Values: []float32{40, 50, 60},
				},
			},
			Colors:             nil,
			BarLabelVisible:    nil,
			BarLabelPosition:   0,
			PointVisible:       &wrapperspb.BoolValue{Value: true},
			PointType:          render.ChartView_X,
			PointLabelVisible:  &wrapperspb.BoolValue{Value: true},
			PointLabelPosition: render.ChartView_LEFT,
		},
	}
}

func NewVerticalBarView() *ChartView {
	return &ChartView{
		&render.ChartView{
			Kind: render.ChartView_VERTICAL_BAR,
			Values: &render.ChartView_BarsValues{
				BarsValues: &render.ChartViewBarsValues{
					BarsDatasets: []*render.ChartViewBarsValues_BarsDataset{
						{
							Values: []float32{11, 22, 33},
							Colors: &render.ChartViewBarsValues_ChartViewBarsColors{
								Fill: &render.ChartElementColor{
									ColorValue: &render.ChartElementColor_ColorHex{
										ColorHex: "#fff4e6",
									},
								},
								Stroke: &render.ChartElementColor{
									ColorValue: &render.ChartElementColor_ColorHex{
										ColorHex: "#3c2f2f",
									},
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
		},
	}
}

func (v *ChartView) Unembed() *render.ChartView {
	return v.ChartView
}

func (v *ChartView) SetDefaultColors() *ChartView {
	v.Colors = ColorsDefault()

	return v
}

func (v *ChartView) SetFillAndStrokeColor() *ChartView {
	colors := v.Colors
	if colors == nil {
		colors = &render.ChartViewColors{}
	}

	colors.Fill = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: "#410037",
		},
	}
	colors.Stroke = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: "#270021",
		},
	}
	v.Colors = colors

	return v
}

func (v *ChartView) SetBadFillRGBColor() *ChartView {
	colors := v.Colors
	if colors == nil {
		colors = &render.ChartViewColors{}
	}

	// nolint: gomnd
	colors.Fill = &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorRgb{
			ColorRgb: &render.ChartElementColor_RGB{
				R: 1,
				G: 320,
				B: 2,
			},
		},
	}
	v.Colors = colors

	return v
}

func (v *ChartView) SetBarLabelPosition() *ChartView {
	v.BarLabelPosition = render.ChartView_END_INSIDE

	return v
}

func (v *ChartView) SetDefaultPointParams() *ChartView {
	v.PointVisible = &wrapperspb.BoolValue{Value: true}
	v.PointType = render.ChartView_CIRCLE
	v.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	v.PointLabelPosition = render.ChartView_TOP

	return v
}

func (v *ChartView) SetDefaultPointBools() *ChartView {
	v.PointVisible = &wrapperspb.BoolValue{Value: true}
	v.PointLabelVisible = &wrapperspb.BoolValue{Value: true}

	return v
}

func (v *ChartView) SetDefaultBarParams() *ChartView {
	v.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	v.BarLabelPosition = render.ChartView_CENTER

	return v
}

func (v *ChartView) SetDefaultBarBools() *ChartView {
	v.BarLabelVisible = &wrapperspb.BoolValue{Value: true}

	return v
}

func (v *ChartView) UnsetKind() *ChartView {
	v.Kind = render.ChartView_UNSPECIFIED_KIND

	return v
}

func (v *ChartView) UnsetValues() *ChartView {
	v.Values = nil

	return v
}
