package color

import (
	"fmt"
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
			t.Errorf("RGB(%f, %f, %f) != %v", test.r, test.g, test.b, test.want)
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
			t.Errorf("RGBA(%f, %f, %f, %f) != %v", test.r, test.g, test.b, test.a, test.want)
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

func TestRGBHex(t *testing.T) {
	var tests = []struct {
		r    int
		g    int
		b    int
		want Color
	}{
		{0, 0, 0, Color{0, 0, 0, 1}},
		{255, 0, 0, Color{1, 0, 0, 1}},
		{0, 255, 0, Color{0, 1, 0, 1}},
		{0, 0, 255, Color{0, 0, 1, 1}},
		{0, 17, 255, Color{0, 0.06666666666666667, 1, 1}},
	}
	for _, test := range tests {
		result := RGBHex(test.r, test.g, test.b)
		if !equalsColor(result, test.want) {
			t.Errorf("RGBHex(%d, %d, %d) != %v", test.r, test.g, test.b, test.want)
		}
	}
}

func TestRGBAHex(t *testing.T) {
	var tests = []struct {
		r    int
		g    int
		b    int
		a    int
		want Color
	}{
		{0, 0, 0, 255, Color{0, 0, 0, 1}},
		{255, 0, 0, 0, Color{1, 0, 0, 0}},
		{0, 255, 0, 255, Color{0, 1, 0, 1}},
		{0, 0, 255, 0, Color{0, 0, 1, 0}},
		{0, 17, 255, 17, Color{0, 0.06666666666666667, 1, 0.06666666666666667}},
	}
	for _, test := range tests {
		result := RGBAHex(test.r, test.g, test.b, test.a)
		if !equalsColor(result, test.want) {
			t.Errorf("RGBAHex(%d, %d, %d, %d) != %v", test.r, test.g, test.b, test.a, test.want)
		}
	}
}

func TestRandomRGB(t *testing.T) {
	result := RandomRGB()
	if !(result.R >= 0.0 && result.R < 1.0 &&
		result.G >= 0.0 && result.G < 1.0 &&
		result.B >= 0.0 && result.B < 1.0) {
		t.Errorf("RandomRGB() rgb values not between 0.0 and 1.0")
	}
}

func TestHSV(t *testing.T) {
	var tests = []struct {
		h    float64
		s    float64
		v    float64
		want Color
	}{
		{0, 0, 0, Color{0, 0, 0, 1}},
		{0, 0, 1, Color{1, 1, 1, 1}},

		{0, 1, 1, Color{1, 0, 0, 1}},
		{30, 1, 1, Color{1, 0.5, 0, 1}},
		{60, 1, 1, Color{1, 1, 0, 1}},
		{90, 1, 1, Color{0.5, 1, 0, 1}},
		{120, 1, 1, Color{0, 1, 0, 1}},
		{180, 1, 1, Color{0, 1, 1, 1}},
		{240, 1, 1, Color{0, 0, 1, 1}},
		{300, 1, 1, Color{1, 0, 1, 1}},
		{360, 1, 1, Color{1, 0, 0, 1}},
	}
	for _, test := range tests {
		result := HSV(test.h, test.s, test.v)
		if !equalsColor(result, test.want) {
			t.Errorf("HSV(%f, %f, %f) != %v", test.h, test.s, test.v, test.want)
			fmt.Println(result)
		}
	}
}

func TestGrey(t *testing.T) {
	result := Grey(0)
	want := Color{0, 0, 0, 1}
	if !equalsColor(result, want) {
		t.Errorf("Grey(%f) != %v", 0.0, want)
	}

	result = Grey(1)
	want = Color{1, 1, 1, 1}
	if !equalsColor(result, want) {
		t.Errorf("Grey(%f) != %v", 1.0, want)
	}

	result = Grey(0.5)
	if !(result.R == result.G && result.G == result.B) {
		t.Errorf("Grey(0.5) rgb values are not equal.")
	}

}

func TestRandomGrey(t *testing.T) {
	result := RandomGrey()
	if !(result.R >= 0.0 && result.R < 1.0) {
		t.Errorf("RandomGrey() rgb values not between 0.0 and 1.0")
	}
	if !(result.R == result.G && result.G == result.B) {
		t.Errorf("RandomGrey() rgb values not equal")
	}
}

func TestRandomGreyRange(t *testing.T) {
	result := RandomGreyRange(0.0, 1.0)
	if !(result.R >= 0.0 && result.R < 1.0) {
		t.Errorf("RandomGreyRange(0.0, 1.0) rgb values not between 0.0 and 1.0")
	}
	if !(result.R == result.G && result.G == result.B) {
		t.Errorf("RandomGreyRange(0.0, 1.0) rgb values not equal")
	}

	result = RandomGreyRange(0.0, 0.5)
	if !(result.R >= 0.0 && result.R < 0.5) {
		t.Errorf("RandomGreyRange(0.0, 0.5) rgb values not between 0.0 and 0.5")
	}

	result = RandomGreyRange(0.5, 1.0)
	if !(result.R >= 0.5 && result.R < 1.0) {
		t.Errorf("RandomGreyRange(0.5, 1.0) rgb values not between 0.5 and 1.0")
	}
}
