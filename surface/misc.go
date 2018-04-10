package surface

import "math"

///////////////////////////////////////
// all this will go elsewhere when the
// correct packages are created for it.
///////////////////////////////////////

// Point x y point
type Point struct {
	X float64
	Y float64
}

// TwoPi 2 pi
const TwoPi = math.Pi * 2

// HalfPi pi / 2
const HalfPi = math.Pi / 2

func clamp(value float64, min float64, max float64) float64 {
	return math.Min(max, math.Max(value, min))
}
