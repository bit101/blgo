package logdisplay

import "math"

// LogDisplayFilled represents a bit array. All elements are filled with arbitrary values.
// Max and min values are kept track of.
// The value of each pixel is then log10(value - min) / log10(max - min).
type LogDisplayFilled struct {
	width, height int
	values        []float64
	max, min      float64
}

// NewLogDisplayFilled creates a new LogDisplayFilled struct.
func NewLogDisplayFilled(width, height int) *LogDisplayFilled {
	return &LogDisplayFilled{
		width:  width,
		height: height,
		max:    0.0,
		min:    math.MaxFloat64,
		values: make([]float64, width*height),
	}
}

// Set sets the raw value of a pixel.
func (d *LogDisplayFilled) Set(value, x, y float64) {
	xx, yy := int(x), int(y)
	if xx >= 0 && xx < d.width && yy >= 0 && yy < d.height {
		index := xx + yy*d.width
		d.values[index] = value
		if value > d.max {
			d.max = value
		}
		if value < d.min {
			d.min = value
		}
	}
}

// Get calculates the logarithmic value of the pixel.
func (d *LogDisplayFilled) Get(x, y float64) float64 {
	xx, yy := int(x), int(y)
	value := d.values[xx+yy*d.width]
	return math.Log(value-d.min) / math.Log(d.max-d.min)
}
