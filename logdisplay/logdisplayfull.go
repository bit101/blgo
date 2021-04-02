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
func (d *LogDisplayFilled) Set(value float64, x, y int) {
	if x >= 0 && x < d.width && y >= 0 && y < d.height {
		index := x + y*d.width
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
func (d *LogDisplayFilled) Get(x, y int) float64 {
	value := d.values[x+y*d.width]
	return math.Log(value-d.min) / math.Log(d.max-d.min)
}
