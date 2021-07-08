package testutils

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverrest/view/v0"
)

func ColorsDefault() *render.ChartViewColors {
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

func HorizontalBarView() *render.ChartView {
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

func JSONHorizontalBarView() *view.ChartView {
	return &view.ChartView{
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
	}
}

func JSONHorizontalBarViewWithScalarAndBarsValues() *view.ChartView {
	res := JSONHorizontalBarView()
	res.ScalarValues = &view.ScalarValues{
		Values: []float32{10, 20},
	}

	return res
}

func JSONHorizontalBarViewWithScalarAndPointsValues() *view.ChartView {
	res := JSONHorizontalBarView()
	res.BarsValues = nil
	res.ScalarValues = &view.ScalarValues{
		Values: []float32{10, 20},
	}
	res.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20}},
	}

	return res
}

func JSONHorizontalBarViewWithPointsAndBarsValues() *view.ChartView {
	res := JSONHorizontalBarView()
	res.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20}},
	}

	return res
}

func JSONHorizontalBarViewWithAllValues() *view.ChartView {
	res := JSONHorizontalBarView()
	res.ScalarValues = &view.ScalarValues{
		Values: []float32{10, 20},
	}
	res.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20}},
	}

	return res
}

func HorizontalBarViewWithDefaults() *render.ChartView {
	res := HorizontalBarView()
	res.Colors = ColorsDefault()
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointType = render.ChartView_CIRCLE
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelPosition = render.ChartView_TOP

	return res
}

func HorizontalBarViewWithBoolDefaults() *render.ChartView {
	res := HorizontalBarView()
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}

	return res
}

func AreaView() *render.ChartView {
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

func JSONAreaView() *view.ChartView {
	return &view.ChartView{
		Kind:         "area",
		BarsValues:   nil,
		PointsValues: nil,
		ScalarValues: &view.ScalarValues{
			Values: []float32{1010, 2100},
		},
		Colors: &view.ChartViewColors{
			FillColor: &view.ChartElementColor{
				Hex: "#71c7ee",
				RGB: nil,
			},
			StrokeColor: &view.ChartElementColor{
				Hex: "#005072",
				RGB: nil,
			},
			PointFillColor:   nil,
			PointStrokeColor: nil,
		},
		BarLabelVisible:    nil,
		BarLabelPosition:   "",
		PointVisible:       boolToPtr(false),
		PointType:          "x",
		PointLabelVisible:  boolToPtr(false),
		PointLabelPosition: "top",
	}
}

func JSONAreaViewBadPointsCount() *view.ChartView {
	res := JSONAreaView()
	res.ScalarValues = nil
	res.PointsValues = &view.PointsValues{
		Values: [][]float32{{10, 20, 30}, {40}},
	}

	return res
}

func AreaViewWithDefaults() *render.ChartView {
	res := AreaView()
	res.Colors = ColorsDefault()
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

func AreaViewWithBoolDefaults() *render.ChartView {
	res := AreaView()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}

	return res
}

func AreaViewBadRGBColor() *render.ChartView {
	res := AreaView()
	res.Colors = ColorsDefault()

	//nolint: gomnd
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

func ScatterView() *render.ChartView {
	//nolint: gomnd
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

func ScatterViewWithDefaults() *render.ChartView {
	res := ScatterView()
	res.Colors = ColorsDefault()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_CENTER

	return res
}

func LineView() *render.ChartView {
	return &render.ChartView{
		Kind: render.ChartView_LINE,
		Values: &render.ChartView_ScalarValues{
			ScalarValues: &render.ChartViewScalarValues{
				Values: []float32{40, 50, 60},
			},
		},
		Colors:             ColorsDefault(),
		BarLabelVisible:    nil,
		BarLabelPosition:   0,
		PointVisible:       &wrapperspb.BoolValue{Value: true},
		PointType:          render.ChartView_X,
		PointLabelVisible:  &wrapperspb.BoolValue{Value: true},
		PointLabelPosition: render.ChartView_LEFT,
	}
}

func JSONLineView() *view.ChartView {
	return &view.ChartView{
		Kind:         "line",
		BarsValues:   nil,
		PointsValues: nil,
		ScalarValues: &view.ScalarValues{
			Values: []float32{40, 50, 60},
		},
		Colors: &view.ChartViewColors{
			FillColor: &view.ChartElementColor{
				Hex: "#71c7ec",
				RGB: nil,
			},
			StrokeColor: &view.ChartElementColor{
				Hex: "#005073",
				RGB: nil,
			},
			PointFillColor: &view.ChartElementColor{
				Hex: "#71c7ec",
				RGB: nil,
			},
			PointStrokeColor: &view.ChartElementColor{
				Hex: "#005073",
				RGB: nil,
			},
		},
		BarLabelVisible:    nil,
		BarLabelPosition:   "",
		PointVisible:       boolToPtr(true),
		PointType:          "x",
		PointLabelVisible:  boolToPtr(true),
		PointLabelPosition: "left",
	}
}

func LineViewWithDefaults() *render.ChartView {
	res := LineView()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_CENTER

	return res
}

func LineViewWithBoolDefaults() *render.ChartView {
	res := LineView()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}

	return res
}

func VerticalBarView() *render.ChartView {
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

func JSONVerticalBarView() *view.ChartView {
	return &view.ChartView{
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
	}
}

func VerticalBarViewWithDefaults() *render.ChartView {
	res := VerticalBarView()
	res.Colors = ColorsDefault()
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointType = render.ChartView_CIRCLE
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelPosition = render.ChartView_TOP

	return res
}

func VerticalBarViewWithBoolDefaultsAndEndInsideLabel() *render.ChartView {
	res := VerticalBarView()
	res.BarLabelVisible = &wrapperspb.BoolValue{Value: true}
	res.BarLabelPosition = render.ChartView_END_INSIDE
	res.PointVisible = &wrapperspb.BoolValue{Value: true}
	res.PointLabelVisible = &wrapperspb.BoolValue{Value: true}

	return res
}

func UnspecifiedKindView() *render.ChartView {
	res := HorizontalBarView()
	res.Kind = render.ChartView_UNSPECIFIED_KIND

	return res
}

func HorizontalBarViewWithoutValues() *render.ChartView {
	res := HorizontalBarView()
	res.Values = nil

	return res
}
