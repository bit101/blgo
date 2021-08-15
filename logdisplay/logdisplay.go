package logdisplay

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/color"
)

// LogDisplay represents a bit array. Pixel values are incremented or set directly.
// Max value is kept track of
// The value of each pixel is then log10(value) / log10(max).
type LogDisplay struct {
	surface       *blgo.Surface
	width, height int
	values        []float64
	max           float64
}

// NewLogDisplay creates a new LogDisplay struct.
func NewLogDisplay(surface *blgo.Surface) *LogDisplay {
	width := int(surface.Width)
	height := int(surface.Height)
	return &LogDisplay{
		surface: surface,
		width:   width,
		height:  height,
		max:     0.0,
		values:  make([]float64, width*height),
	}
}

// Inc increments the raw value of a pixel by 1.
func (d *LogDisplay) Inc(x, y float64) {
	xx, yy := int(x), int(y)
	if xx >= 0 && xx < d.width && yy >= 0 && yy < d.height {
		index := xx + yy*d.width
		value := d.values[index] + 1
		d.values[index] = value
		if value > d.max {
			d.max = value
		}
	}
}

// Get calculates the logarithmic value of the pixel.
func (d *LogDisplay) Get(x, y int) float64 {
	xx, yy := x, y
	value := d.values[xx+yy*d.width]
	return math.Log(value) / math.Log(d.max)
}

func (d *LogDisplay) RenderGrey() {
	d.Render(color.Grey)
}

func (d *LogDisplay) RenderHSV(min, max float64) {
	d.Render(func(g float64) color.Color {
		h := g*(max-min) + min
		return color.HSV(h, 1, g)
	})
}

func (d *LogDisplay) Render(colorFunc func(float64) color.Color) {
	for x := 0; x < d.width; x++ {
		for y := 0; y < d.height; y++ {
			g := d.Get(x, y)
			d.surface.SetSourceColor(colorFunc(g))
			d.surface.FillRectangle(float64(x), float64(y), 1, 1)
		}
	}
}
