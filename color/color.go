package color

import (
	"math"
	"math/rand"
)

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func RGB(r float64, g float64, b float64) Color {
	return Color{r, g, b, 1.0}
}

func RGBA(r float64, g float64, b float64, a float64) Color {
	return Color{r, g, b, a}
}

func Number(value int) Color {
	r := (value >> 16 & 0xff)
	g := (value >> 8 & 0xff)
	b := (value & 0xff)
	return RGBHex(r, g, b)
}

func NumberWithAlpha(value int) Color {
	a := (value >> 24)
	r := (value >> 16 & 0xff)
	g := (value >> 8 & 0xff)
	b := (value & 0xff)
	return RGBAHex(r, g, b, a)
}

func RGBHex(r int, g int, b int) Color {
	return RGBAHex(r, g, b, 255)
}

func RGBAHex(r int, g int, b int, a int) Color {
	return RGBA(float64(r)/255.0, float64(g)/255.0, float64(b)/255.0, float64(a)/255.0)
}

func RandomRGB() Color {
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	return RGB(r, g, b)
}

func HSV(h float64, s float64, v float64) Color {
	h = h / 360.0
	i := math.Floor(h * 6.0)
	f := h*6.0 - i
	p := v * (1.0 - s)
	q := v * (1.0 - f*s)
	t := v * (1.0 - (1.0-f)*s)
	switch int(i) % 6 {
	case 0:
		return RGB(v, t, p)
	case 1:
		return RGB(q, v, p)
	case 2:
		return RGB(p, v, t)
	case 3:
		return RGB(p, q, v)
	case 4:
		return RGB(t, p, v)
	case 5:
		return RGB(v, p, q)
	default:
		return RGB(0.0, 0.0, 0.0)
	}
}

func Grey(shade float64) Color {
	return RGB(shade, shade, shade)
}

func GreyHex(shade int) Color {
	return Grey(float64(shade) / 255.0)
}

func RandomGrey() Color {
	return RandomGreyRange(0.0, 1.0)
}

func RandomGreyRange(min float64, max float64) Color {
	return Grey(min + rand.Float64()*(max-min))
}

func Blueviolet() Color           { return RGBHex(138, 43, 226) }
func Brown() Color                { return RGBHex(165, 42, 42) }
func Aliceblue() Color            { return RGBHex(240, 248, 255) }
func Antiquewhite() Color         { return RGBHex(250, 235, 215) }
func Aqua() Color                 { return RGBHex(0, 255, 255) }
func Aquamarine() Color           { return RGBHex(127, 255, 212) }
func Azure() Color                { return RGBHex(240, 255, 255) }
func Beige() Color                { return RGBHex(245, 245, 220) }
func Bisque() Color               { return RGBHex(255, 228, 196) }
func Black() Color                { return RGBHex(0, 0, 0) }
func Blanchedalmond() Color       { return RGBHex(255, 235, 205) }
func Blue() Color                 { return RGBHex(0, 0, 255) }
func Burlywood() Color            { return RGBHex(222, 184, 135) }
func Cadetblue() Color            { return RGBHex(95, 158, 160) }
func Chartreuse() Color           { return RGBHex(127, 255, 0) }
func Chocolate() Color            { return RGBHex(210, 105, 30) }
func Coral() Color                { return RGBHex(255, 127, 80) }
func Cornflowerblue() Color       { return RGBHex(100, 149, 237) }
func Cornsilk() Color             { return RGBHex(255, 248, 220) }
func Crimson() Color              { return RGBHex(220, 20, 60) }
func Cyan() Color                 { return RGBHex(0, 255, 255) }
func Darkblue() Color             { return RGBHex(0, 0, 139) }
func Darkcyan() Color             { return RGBHex(0, 139, 139) }
func Darkgoldenrod() Color        { return RGBHex(184, 134, 11) }
func Darkgray() Color             { return RGBHex(169, 169, 169) }
func Darkgreen() Color            { return RGBHex(0, 100, 0) }
func Darkgrey() Color             { return RGBHex(169, 169, 169) }
func Darkkhaki() Color            { return RGBHex(189, 183, 107) }
func Darkmagenta() Color          { return RGBHex(139, 0, 139) }
func Darkolivegreen() Color       { return RGBHex(85, 107, 47) }
func Darkorange() Color           { return RGBHex(255, 140, 0) }
func Darkorchid() Color           { return RGBHex(153, 50, 204) }
func Darkred() Color              { return RGBHex(139, 0, 0) }
func Darksalmon() Color           { return RGBHex(233, 150, 122) }
func Darkseagreen() Color         { return RGBHex(143, 188, 143) }
func Darkslateblue() Color        { return RGBHex(72, 61, 139) }
func Darkslategray() Color        { return RGBHex(47, 79, 79) }
func Darkslategrey() Color        { return RGBHex(47, 79, 79) }
func Darkturquoise() Color        { return RGBHex(0, 206, 209) }
func Darkviolet() Color           { return RGBHex(148, 0, 211) }
func Deeppink() Color             { return RGBHex(255, 20, 147) }
func Deepskyblue() Color          { return RGBHex(0, 191, 255) }
func Dimgray() Color              { return RGBHex(105, 105, 105) }
func Dimgrey() Color              { return RGBHex(105, 105, 105) }
func Dodgerblue() Color           { return RGBHex(30, 144, 255) }
func Firebrick() Color            { return RGBHex(178, 34, 34) }
func Floralwhite() Color          { return RGBHex(255, 250, 240) }
func Forestgreen() Color          { return RGBHex(34, 139, 34) }
func Fuchsia() Color              { return RGBHex(255, 0, 255) }
func Gainsboro() Color            { return RGBHex(220, 220, 220) }
func Ghostwhite() Color           { return RGBHex(248, 248, 255) }
func Gold() Color                 { return RGBHex(255, 215, 0) }
func Goldenrod() Color            { return RGBHex(218, 165, 32) }
func Gray() Color                 { return RGBHex(128, 128, 128) }
func Green() Color                { return RGBHex(0, 128, 0) }
func Greenyellow() Color          { return RGBHex(173, 255, 47) }
func Honeydew() Color             { return RGBHex(240, 255, 240) }
func Hotpink() Color              { return RGBHex(255, 105, 180) }
func Indianred() Color            { return RGBHex(205, 92, 92) }
func Indigo() Color               { return RGBHex(75, 0, 130) }
func Ivory() Color                { return RGBHex(255, 255, 240) }
func Khaki() Color                { return RGBHex(240, 230, 140) }
func Lavender() Color             { return RGBHex(230, 230, 250) }
func Lavenderblush() Color        { return RGBHex(255, 240, 245) }
func Lawngreen() Color            { return RGBHex(124, 252, 0) }
func Lemonchiffon() Color         { return RGBHex(255, 250, 205) }
func Lightblue() Color            { return RGBHex(173, 216, 230) }
func Lightcoral() Color           { return RGBHex(240, 128, 128) }
func Lightcyan() Color            { return RGBHex(224, 255, 255) }
func Lightgoldenrodyellow() Color { return RGBHex(250, 250, 210) }
func Lightgray() Color            { return RGBHex(211, 211, 211) }
func Lightgreen() Color           { return RGBHex(144, 238, 144) }
func Lightgrey() Color            { return RGBHex(211, 211, 211) }
func Lightpink() Color            { return RGBHex(255, 182, 193) }
func Lightsalmon() Color          { return RGBHex(255, 160, 122) }
func Lightseagreen() Color        { return RGBHex(32, 178, 170) }
func Lightskyblue() Color         { return RGBHex(135, 206, 250) }
func Lightslategray() Color       { return RGBHex(119, 136, 153) }
func Lightslategrey() Color       { return RGBHex(119, 136, 153) }
func Lightsteelblue() Color       { return RGBHex(176, 196, 222) }
func Lightyellow() Color          { return RGBHex(255, 255, 224) }
func Lime() Color                 { return RGBHex(0, 255, 0) }
func Limegreen() Color            { return RGBHex(50, 205, 50) }
func Linen() Color                { return RGBHex(250, 240, 230) }
func Magenta() Color              { return RGBHex(255, 0, 255) }
func Maroon() Color               { return RGBHex(128, 0, 0) }
func Mediumaquamarine() Color     { return RGBHex(102, 205, 170) }
func Mediumblue() Color           { return RGBHex(0, 0, 205) }
func Mediumorchid() Color         { return RGBHex(186, 85, 211) }
func Mediumpurple() Color         { return RGBHex(147, 112, 219) }
func Mediumseagreen() Color       { return RGBHex(60, 179, 113) }
func Mediumslateblue() Color      { return RGBHex(123, 104, 238) }
func Mediumspringgreen() Color    { return RGBHex(0, 250, 154) }
func Mediumturquoise() Color      { return RGBHex(72, 209, 204) }
func Mediumvioletred() Color      { return RGBHex(199, 21, 133) }
func Midnightblue() Color         { return RGBHex(25, 25, 112) }
func Mintcream() Color            { return RGBHex(245, 255, 250) }
func Mistyrose() Color            { return RGBHex(255, 228, 225) }
func Moccasin() Color             { return RGBHex(255, 228, 181) }
func Navajowhite() Color          { return RGBHex(255, 222, 173) }
func Navy() Color                 { return RGBHex(0, 0, 128) }
func Oldlace() Color              { return RGBHex(253, 245, 230) }
func Olive() Color                { return RGBHex(128, 128, 0) }
func Olivedrab() Color            { return RGBHex(107, 142, 35) }
func Orange() Color               { return RGBHex(255, 165, 0) }
func Orangered() Color            { return RGBHex(255, 69, 0) }
func Orchid() Color               { return RGBHex(218, 112, 214) }
func Palegoldenrod() Color        { return RGBHex(238, 232, 170) }
func Palegreen() Color            { return RGBHex(152, 251, 152) }
func Paleturquoise() Color        { return RGBHex(175, 238, 238) }
func Palevioletred() Color        { return RGBHex(219, 112, 147) }
func Papayawhip() Color           { return RGBHex(255, 239, 213) }
func Peachpuff() Color            { return RGBHex(255, 218, 185) }
func Peru() Color                 { return RGBHex(205, 133, 63) }
func Pink() Color                 { return RGBHex(255, 192, 203) }
func Plum() Color                 { return RGBHex(221, 160, 221) }
func Powderblue() Color           { return RGBHex(176, 224, 230) }
func Purple() Color               { return RGBHex(128, 0, 128) }
func Rebeccapurple() Color        { return RGBHex(102, 51, 153) }
func Red() Color                  { return RGBHex(255, 0, 0) }
func Rosybrown() Color            { return RGBHex(188, 143, 143) }
func Royalblue() Color            { return RGBHex(65, 105, 225) }
func Saddlebrown() Color          { return RGBHex(139, 69, 19) }
func Salmon() Color               { return RGBHex(250, 128, 114) }
func Sandybrown() Color           { return RGBHex(244, 164, 96) }
func Seagreen() Color             { return RGBHex(46, 139, 87) }
func Seashell() Color             { return RGBHex(255, 245, 238) }
func Sienna() Color               { return RGBHex(160, 82, 45) }
func Silver() Color               { return RGBHex(192, 192, 192) }
func Skyblue() Color              { return RGBHex(135, 206, 235) }
func Slateblue() Color            { return RGBHex(106, 90, 205) }
func Slategray() Color            { return RGBHex(112, 128, 144) }
func Slategrey() Color            { return RGBHex(112, 128, 144) }
func Snow() Color                 { return RGBHex(255, 250, 250) }
func Springgreen() Color          { return RGBHex(0, 255, 127) }
func Steelblue() Color            { return RGBHex(70, 130, 180) }
func Tan() Color                  { return RGBHex(210, 180, 140) }
func Teal() Color                 { return RGBHex(0, 128, 128) }
func Thistle() Color              { return RGBHex(216, 191, 216) }
func Tomato() Color               { return RGBHex(255, 99, 71) }
func Turquoise() Color            { return RGBHex(64, 224, 208) }
func Violet() Color               { return RGBHex(238, 130, 238) }
func Wheat() Color                { return RGBHex(245, 222, 179) }
func White() Color                { return RGBHex(255, 255, 255) }
func Whitesmoke() Color           { return RGBHex(245, 245, 245) }
func Yellow() Color               { return RGBHex(255, 255, 0) }
func Yellowgreen() Color          { return RGBHex(154, 205, 50) }
