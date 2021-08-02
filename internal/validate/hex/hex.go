package hex

import (
	"errors"
	"fmt"
	"strings"

	"github.com/limpidchart/lc-api/internal/render/github.com/limpidchart/lc-proto/render/v0"
	"github.com/limpidchart/lc-api/internal/serverhttp/v0/view"
)

const (
	hexLen = 7

	hexStart = '#'
)

// nolint: gochecknoglobals
var hexTable = map[uint8]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
	'a': {},
	'b': {},
	'c': {},
	'd': {},
	'e': {},
	'f': {},
}

var (
	// ErrHexDoesntStartWithHash contains error message about bad starting value.
	ErrHexDoesntStartWithHash = errors.New("hex value should start with #")

	// ErrHexContainsUnexpectedSymbol contains error message about bad hex symbol.
	ErrHexContainsUnexpectedSymbol = errors.New("hex value contains unexpected symbol")

	// ErrHexBadValueLen contains error message about bad hex length.
	ErrHexBadValueLen = fmt.Errorf("hex value len should be %d (including # symbol) if set", hexLen)
)

// ValidateChartElementColor parses and validates chart element hex color.
func ValidateChartElementColor(chartElementColor *render.ChartElementColor) (*render.ChartElementColor, error) {
	colorHex := chartElementColor.GetColorHex()

	if colorHex == "" {
		return &render.ChartElementColor{
			ColorValue: &render.ChartElementColor_ColorHex{
				ColorHex: "",
			},
		}, nil
	}

	h, err := hexColorValue(chartElementColor.GetColorHex())
	if err != nil {
		return nil, err
	}

	return &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: h,
		},
	}, nil
}

// ValidateChartElementColorJSON parses and validates chart element hex color JSON representation.
func ValidateChartElementColorJSON(color *view.ChartElementColor) (*render.ChartElementColor, error) {
	h, err := hexColorValue(color.Hex)
	if err != nil {
		return nil, err
	}

	return &render.ChartElementColor{
		ColorValue: &render.ChartElementColor_ColorHex{
			ColorHex: h,
		},
	}, nil
}

func hexColorValue(h string) (string, error) {
	if len(h) == 0 {
		return "", nil
	}

	if len(h) != hexLen {
		return "", ErrHexBadValueLen
	}

	if h[0] != hexStart {
		return "", ErrHexDoesntStartWithHash
	}

	h = strings.ToLower(h)

	for i := 1; i < len(h); i++ {
		if _, ok := hexTable[h[i]]; !ok {
			return "", ErrHexContainsUnexpectedSymbol
		}
	}

	return h, nil
}
