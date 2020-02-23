package color

import (
	"fmt"
	"testing"
)

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
		if result != test.want {
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
		if result != test.want {
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
		if result != test.want {
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
		if result != test.want {
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
		if result != test.want {
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
		if result != test.want {
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
		if result != test.want {
			t.Errorf("HSV(%f, %f, %f) != %v", test.h, test.s, test.v, test.want)
			fmt.Println(result)
		}
	}
}

func TestLerp(t *testing.T) {
	result := Lerp(Color{0.0, 0.0, 0.0, 1.0}, Color{0.5, 1.0, 0.0, 1.0}, 0.5)
	want := Color{0.25, 0.5, 0.0, 1.0}
	if result != want {
		t.Errorf("%v != %v", result, want)
	}
}

func TestGrey(t *testing.T) {
	result := Grey(0)
	want := Color{0, 0, 0, 1}
	if result != want {
		t.Errorf("Grey(%f) != %v", 0.0, want)
	}

	result = Grey(1)
	want = Color{1, 1, 1, 1}
	if result != want {
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

func TestBlueviolet(t *testing.T) {
	value := Blueviolet()
	want := RGBHex(138, 43, 226)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBrown(t *testing.T) {
	value := Brown()
	want := RGBHex(165, 42, 42)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestAliceblue(t *testing.T) {
	value := Aliceblue()
	want := RGBHex(240, 248, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestAntiquewhite(t *testing.T) {
	value := Antiquewhite()
	want := RGBHex(250, 235, 215)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestAqua(t *testing.T) {
	value := Aqua()
	want := RGBHex(0, 255, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestAquamarine(t *testing.T) {
	value := Aquamarine()
	want := RGBHex(127, 255, 212)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestAzure(t *testing.T) {
	value := Azure()
	want := RGBHex(240, 255, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBeige(t *testing.T) {
	value := Beige()
	want := RGBHex(245, 245, 220)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBisque(t *testing.T) {
	value := Bisque()
	want := RGBHex(255, 228, 196)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBlack(t *testing.T) {
	value := Black()
	want := RGBHex(0, 0, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBlanchedalmond(t *testing.T) {
	value := Blanchedalmond()
	want := RGBHex(255, 235, 205)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBlue(t *testing.T) {
	value := Blue()
	want := RGBHex(0, 0, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestBurlywood(t *testing.T) {
	value := Burlywood()
	want := RGBHex(222, 184, 135)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCadetblue(t *testing.T) {
	value := Cadetblue()
	want := RGBHex(95, 158, 160)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestChartreuse(t *testing.T) {
	value := Chartreuse()
	want := RGBHex(127, 255, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestChocolate(t *testing.T) {
	value := Chocolate()
	want := RGBHex(210, 105, 30)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCoral(t *testing.T) {
	value := Coral()
	want := RGBHex(255, 127, 80)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCornflowerblue(t *testing.T) {
	value := Cornflowerblue()
	want := RGBHex(100, 149, 237)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCornsilk(t *testing.T) {
	value := Cornsilk()
	want := RGBHex(255, 248, 220)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCrimson(t *testing.T) {
	value := Crimson()
	want := RGBHex(220, 20, 60)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestCyan(t *testing.T) {
	value := Cyan()
	want := RGBHex(0, 255, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkblue(t *testing.T) {
	value := Darkblue()
	want := RGBHex(0, 0, 139)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkcyan(t *testing.T) {
	value := Darkcyan()
	want := RGBHex(0, 139, 139)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkgoldenrod(t *testing.T) {
	value := Darkgoldenrod()
	want := RGBHex(184, 134, 11)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkgray(t *testing.T) {
	value := Darkgray()
	want := RGBHex(169, 169, 169)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkgreen(t *testing.T) {
	value := Darkgreen()
	want := RGBHex(0, 100, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkgrey(t *testing.T) {
	value := Darkgrey()
	want := RGBHex(169, 169, 169)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkkhaki(t *testing.T) {
	value := Darkkhaki()
	want := RGBHex(189, 183, 107)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkmagenta(t *testing.T) {
	value := Darkmagenta()
	want := RGBHex(139, 0, 139)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkolivegreen(t *testing.T) {
	value := Darkolivegreen()
	want := RGBHex(85, 107, 47)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkorange(t *testing.T) {
	value := Darkorange()
	want := RGBHex(255, 140, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkorchid(t *testing.T) {
	value := Darkorchid()
	want := RGBHex(153, 50, 204)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkred(t *testing.T) {
	value := Darkred()
	want := RGBHex(139, 0, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarksalmon(t *testing.T) {
	value := Darksalmon()
	want := RGBHex(233, 150, 122)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkseagreen(t *testing.T) {
	value := Darkseagreen()
	want := RGBHex(143, 188, 143)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkslateblue(t *testing.T) {
	value := Darkslateblue()
	want := RGBHex(72, 61, 139)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkslategray(t *testing.T) {
	value := Darkslategray()
	want := RGBHex(47, 79, 79)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkslategrey(t *testing.T) {
	value := Darkslategrey()
	want := RGBHex(47, 79, 79)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkturquoise(t *testing.T) {
	value := Darkturquoise()
	want := RGBHex(0, 206, 209)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDarkviolet(t *testing.T) {
	value := Darkviolet()
	want := RGBHex(148, 0, 211)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDeeppink(t *testing.T) {
	value := Deeppink()
	want := RGBHex(255, 20, 147)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDeepskyblue(t *testing.T) {
	value := Deepskyblue()
	want := RGBHex(0, 191, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDimgray(t *testing.T) {
	value := Dimgray()
	want := RGBHex(105, 105, 105)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDimgrey(t *testing.T) {
	value := Dimgrey()
	want := RGBHex(105, 105, 105)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestDodgerblue(t *testing.T) {
	value := Dodgerblue()
	want := RGBHex(30, 144, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestFirebrick(t *testing.T) {
	value := Firebrick()
	want := RGBHex(178, 34, 34)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestFloralwhite(t *testing.T) {
	value := Floralwhite()
	want := RGBHex(255, 250, 240)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestForestgreen(t *testing.T) {
	value := Forestgreen()
	want := RGBHex(34, 139, 34)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestFuchsia(t *testing.T) {
	value := Fuchsia()
	want := RGBHex(255, 0, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGainsboro(t *testing.T) {
	value := Gainsboro()
	want := RGBHex(220, 220, 220)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGhostwhite(t *testing.T) {
	value := Ghostwhite()
	want := RGBHex(248, 248, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGold(t *testing.T) {
	value := Gold()
	want := RGBHex(255, 215, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGoldenrod(t *testing.T) {
	value := Goldenrod()
	want := RGBHex(218, 165, 32)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGray(t *testing.T) {
	value := Gray()
	want := RGBHex(128, 128, 128)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGreen(t *testing.T) {
	value := Green()
	want := RGBHex(0, 128, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestGreenyellow(t *testing.T) {
	value := Greenyellow()
	want := RGBHex(173, 255, 47)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestHoneydew(t *testing.T) {
	value := Honeydew()
	want := RGBHex(240, 255, 240)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestHotpink(t *testing.T) {
	value := Hotpink()
	want := RGBHex(255, 105, 180)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestIndianred(t *testing.T) {
	value := Indianred()
	want := RGBHex(205, 92, 92)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestIndigo(t *testing.T) {
	value := Indigo()
	want := RGBHex(75, 0, 130)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestIvory(t *testing.T) {
	value := Ivory()
	want := RGBHex(255, 255, 240)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestKhaki(t *testing.T) {
	value := Khaki()
	want := RGBHex(240, 230, 140)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLavender(t *testing.T) {
	value := Lavender()
	want := RGBHex(230, 230, 250)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLavenderblush(t *testing.T) {
	value := Lavenderblush()
	want := RGBHex(255, 240, 245)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLawngreen(t *testing.T) {
	value := Lawngreen()
	want := RGBHex(124, 252, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLemonchiffon(t *testing.T) {
	value := Lemonchiffon()
	want := RGBHex(255, 250, 205)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightblue(t *testing.T) {
	value := Lightblue()
	want := RGBHex(173, 216, 230)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightcoral(t *testing.T) {
	value := Lightcoral()
	want := RGBHex(240, 128, 128)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightcyan(t *testing.T) {
	value := Lightcyan()
	want := RGBHex(224, 255, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightgoldenrodyellow(t *testing.T) {
	value := Lightgoldenrodyellow()
	want := RGBHex(250, 250, 210)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightgray(t *testing.T) {
	value := Lightgray()
	want := RGBHex(211, 211, 211)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightgreen(t *testing.T) {
	value := Lightgreen()
	want := RGBHex(144, 238, 144)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightgrey(t *testing.T) {
	value := Lightgrey()
	want := RGBHex(211, 211, 211)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightpink(t *testing.T) {
	value := Lightpink()
	want := RGBHex(255, 182, 193)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightsalmon(t *testing.T) {
	value := Lightsalmon()
	want := RGBHex(255, 160, 122)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightseagreen(t *testing.T) {
	value := Lightseagreen()
	want := RGBHex(32, 178, 170)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightskyblue(t *testing.T) {
	value := Lightskyblue()
	want := RGBHex(135, 206, 250)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightslategray(t *testing.T) {
	value := Lightslategray()
	want := RGBHex(119, 136, 153)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightslategrey(t *testing.T) {
	value := Lightslategrey()
	want := RGBHex(119, 136, 153)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightsteelblue(t *testing.T) {
	value := Lightsteelblue()
	want := RGBHex(176, 196, 222)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLightyellow(t *testing.T) {
	value := Lightyellow()
	want := RGBHex(255, 255, 224)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLime(t *testing.T) {
	value := Lime()
	want := RGBHex(0, 255, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLimegreen(t *testing.T) {
	value := Limegreen()
	want := RGBHex(50, 205, 50)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestLinen(t *testing.T) {
	value := Linen()
	want := RGBHex(250, 240, 230)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMagenta(t *testing.T) {
	value := Magenta()
	want := RGBHex(255, 0, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMaroon(t *testing.T) {
	value := Maroon()
	want := RGBHex(128, 0, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumaquamarine(t *testing.T) {
	value := Mediumaquamarine()
	want := RGBHex(102, 205, 170)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumblue(t *testing.T) {
	value := Mediumblue()
	want := RGBHex(0, 0, 205)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumorchid(t *testing.T) {
	value := Mediumorchid()
	want := RGBHex(186, 85, 211)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumpurple(t *testing.T) {
	value := Mediumpurple()
	want := RGBHex(147, 112, 219)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumseagreen(t *testing.T) {
	value := Mediumseagreen()
	want := RGBHex(60, 179, 113)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumslateblue(t *testing.T) {
	value := Mediumslateblue()
	want := RGBHex(123, 104, 238)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumspringgreen(t *testing.T) {
	value := Mediumspringgreen()
	want := RGBHex(0, 250, 154)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumturquoise(t *testing.T) {
	value := Mediumturquoise()
	want := RGBHex(72, 209, 204)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMediumvioletred(t *testing.T) {
	value := Mediumvioletred()
	want := RGBHex(199, 21, 133)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMidnightblue(t *testing.T) {
	value := Midnightblue()
	want := RGBHex(25, 25, 112)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMintcream(t *testing.T) {
	value := Mintcream()
	want := RGBHex(245, 255, 250)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMistyrose(t *testing.T) {
	value := Mistyrose()
	want := RGBHex(255, 228, 225)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestMoccasin(t *testing.T) {
	value := Moccasin()
	want := RGBHex(255, 228, 181)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestNavajowhite(t *testing.T) {
	value := Navajowhite()
	want := RGBHex(255, 222, 173)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestNavy(t *testing.T) {
	value := Navy()
	want := RGBHex(0, 0, 128)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOldlace(t *testing.T) {
	value := Oldlace()
	want := RGBHex(253, 245, 230)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOlive(t *testing.T) {
	value := Olive()
	want := RGBHex(128, 128, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOlivedrab(t *testing.T) {
	value := Olivedrab()
	want := RGBHex(107, 142, 35)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOrange(t *testing.T) {
	value := Orange()
	want := RGBHex(255, 165, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOrangered(t *testing.T) {
	value := Orangered()
	want := RGBHex(255, 69, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestOrchid(t *testing.T) {
	value := Orchid()
	want := RGBHex(218, 112, 214)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPalegoldenrod(t *testing.T) {
	value := Palegoldenrod()
	want := RGBHex(238, 232, 170)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPalegreen(t *testing.T) {
	value := Palegreen()
	want := RGBHex(152, 251, 152)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPaleturquoise(t *testing.T) {
	value := Paleturquoise()
	want := RGBHex(175, 238, 238)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPalevioletred(t *testing.T) {
	value := Palevioletred()
	want := RGBHex(219, 112, 147)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPapayawhip(t *testing.T) {
	value := Papayawhip()
	want := RGBHex(255, 239, 213)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPeachpuff(t *testing.T) {
	value := Peachpuff()
	want := RGBHex(255, 218, 185)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPeru(t *testing.T) {
	value := Peru()
	want := RGBHex(205, 133, 63)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPink(t *testing.T) {
	value := Pink()
	want := RGBHex(255, 192, 203)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPlum(t *testing.T) {
	value := Plum()
	want := RGBHex(221, 160, 221)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPowderblue(t *testing.T) {
	value := Powderblue()
	want := RGBHex(176, 224, 230)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestPurple(t *testing.T) {
	value := Purple()
	want := RGBHex(128, 0, 128)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestRebeccapurple(t *testing.T) {
	value := Rebeccapurple()
	want := RGBHex(102, 51, 153)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestRed(t *testing.T) {
	value := Red()
	want := RGBHex(255, 0, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestRosybrown(t *testing.T) {
	value := Rosybrown()
	want := RGBHex(188, 143, 143)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestRoyalblue(t *testing.T) {
	value := Royalblue()
	want := RGBHex(65, 105, 225)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSaddlebrown(t *testing.T) {
	value := Saddlebrown()
	want := RGBHex(139, 69, 19)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSalmon(t *testing.T) {
	value := Salmon()
	want := RGBHex(250, 128, 114)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSandybrown(t *testing.T) {
	value := Sandybrown()
	want := RGBHex(244, 164, 96)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSeagreen(t *testing.T) {
	value := Seagreen()
	want := RGBHex(46, 139, 87)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSeashell(t *testing.T) {
	value := Seashell()
	want := RGBHex(255, 245, 238)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSienna(t *testing.T) {
	value := Sienna()
	want := RGBHex(160, 82, 45)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSilver(t *testing.T) {
	value := Silver()
	want := RGBHex(192, 192, 192)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSkyblue(t *testing.T) {
	value := Skyblue()
	want := RGBHex(135, 206, 235)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSlateblue(t *testing.T) {
	value := Slateblue()
	want := RGBHex(106, 90, 205)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSlategray(t *testing.T) {
	value := Slategray()
	want := RGBHex(112, 128, 144)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSlategrey(t *testing.T) {
	value := Slategrey()
	want := RGBHex(112, 128, 144)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSnow(t *testing.T) {
	value := Snow()
	want := RGBHex(255, 250, 250)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSpringgreen(t *testing.T) {
	value := Springgreen()
	want := RGBHex(0, 255, 127)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestSteelblue(t *testing.T) {
	value := Steelblue()
	want := RGBHex(70, 130, 180)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestTan(t *testing.T) {
	value := Tan()
	want := RGBHex(210, 180, 140)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestTeal(t *testing.T) {
	value := Teal()
	want := RGBHex(0, 128, 128)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestThistle(t *testing.T) {
	value := Thistle()
	want := RGBHex(216, 191, 216)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestTomato(t *testing.T) {
	value := Tomato()
	want := RGBHex(255, 99, 71)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestTurquoise(t *testing.T) {
	value := Turquoise()
	want := RGBHex(64, 224, 208)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestViolet(t *testing.T) {
	value := Violet()
	want := RGBHex(238, 130, 238)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestWheat(t *testing.T) {
	value := Wheat()
	want := RGBHex(245, 222, 179)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestWhite(t *testing.T) {
	value := White()
	want := RGBHex(255, 255, 255)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestWhitesmoke(t *testing.T) {
	value := Whitesmoke()
	want := RGBHex(245, 245, 245)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestYellow(t *testing.T) {
	value := Yellow()
	want := RGBHex(255, 255, 0)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}

func TestYellowgreen(t *testing.T) {
	value := Yellowgreen()
	want := RGBHex(154, 205, 50)
	if value != want {
		t.Errorf("%v != %v", value, want)
	}
}
