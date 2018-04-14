package color

import (
	"testing"
)

func equalsColor(cA Color, cB Color) bool {
	return cA.R == cB.R && cA.G == cB.G && cA.B == cB.B && cA.A == cB.A
}

func TestRGB(t *testing.T) {
	var tests = []struct {
		r    float64
		g    float64
		b    float64
		want Color
	}{
		{0, 0, 0, Color{0, 0, 0, 1}},
		{1, 0, 0, Color{1, 0, 0, 1}},
		{0, 1, 0, Color{0, 1, 0, 1}},
		{0, 0, 1, Color{0, 0, 1, 1}},
		{0, 0.5, 1, Color{0, 0.5, 1, 1}},
	}
	for _, test := range tests {
		result := RGB(test.r, test.g, test.b)
		if !equalsColor(result, test.want) {
			t.Errorf("colorRGB(%f, %f, %f) != %v", test.r, test.g, test.b, test.want)
		}
	}
}

func TestRGBA(t *testing.T) {
	var tests = []struct {
		r    float64
		g    float64
		b    float64
		a    float64
		want Color
	}{
		{0, 0, 0, 0, Color{0, 0, 0, 0}},
		{1, 0, 0, 0.5, Color{1, 0, 0, 0.5}},
		{0, 1, 0, 0.75, Color{0, 1, 0, 0.75}},
		{0, 0, 1, 1, Color{0, 0, 1, 1}},
		{0, 0.5, 1, 0.2, Color{0, 0.5, 1, 0.2}},
	}
	for _, test := range tests {
		result := RGBA(test.r, test.g, test.b, test.a)
		if !equalsColor(result, test.want) {
			t.Errorf("colorRGBA(%f, %f, %f, %f) != %v", test.r, test.g, test.b, test.a, test.want)
		}
	}
}

func TestNumber(t *testing.T) {
	var tests = []struct {
		number int
		want   Color
	}{
		{0x000000, Color{0, 0, 0, 1}},
		{0xff0000, Color{1, 0, 0, 1}},
		{0x00ff00, Color{0, 1, 0, 1}},
		{0x0000ff, Color{0, 0, 1, 1}},
		{0xff00ff, Color{1, 0, 1, 1}},
		{0x1100ff, Color{0.06666666666666667, 0, 1, 1}},
	}
	for _, test := range tests {
		result := Number(test.number)
		if !equalsColor(result, test.want) {
			t.Errorf("Number(%d) != %v", test.number, test.want)
		}
	}
}

func TestNumberWithAlpha(t *testing.T) {
	var tests = []struct {
		number int
		want   Color
	}{
		{0x00000000, Color{0, 0, 0, 0}},
		{0xffffffff, Color{1, 1, 1, 1}},
		{0xffff0000, Color{1, 0, 0, 1}},
		{0x0000ff00, Color{0, 1, 0, 0}},
		{0xff0000ff, Color{0, 0, 1, 1}},
		{0x00ff00ff, Color{1, 0, 1, 0}},
		{0x110000ff, Color{0, 0, 1, 0.06666666666666667}},
	}
	for _, test := range tests {
		result := NumberWithAlpha(test.number)
		if !equalsColor(result, test.want) {
			t.Errorf("NumberWithAlpha(%d) != %v", test.number, test.want)
		}
	}
}
