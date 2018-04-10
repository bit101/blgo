package bitmath

import (
	"math"
)

// TwoPi 2 pi
const TwoPi = math.Pi * 2

// HalfPi pi / 2
const HalfPi = math.Pi / 2

func Norm(value float64, min float64, max float64) float64 {
	return (value - min) / (max - min)
}

func Lerp(min float64, max float64, t float64) float64 {
	return min + (max-min)*t
}

func MapTo(srcValue float64, srcMin float64, srcMax float64, dstMin float64, dstMax float64) float64 {
	norm := Norm(srcValue, srcMin, srcMax)
	return Lerp(dstMin, dstMax, norm)
}

func Clamp(value float64, min float64, max float64) float64 {
	// let min and max be reversed and still work.
	realMin := min
	realMax := max
	if min > max {
		realMin = max
		realMax = min
	}
	result := value
	if value < realMin {
		result = realMin
	}
	if value > realMax {
		result = realMax
	}
	return result
}
