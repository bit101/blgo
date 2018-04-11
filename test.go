package main

import (
	"bitlib/bitmath"
	"bitlib/color"
	"bitlib/surface"
	"math"
)

func main() {
	surface := surface.NewBitSurface(500, 500)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			xx := float64(x)
			yy := float64(y)
			dist := math.Hypot(xx, yy)
			surface.SetSourceColor(color.HSV(bitmath.SinRange(dist*0.1, 0, 360), 1, 1))
			surface.FillRectangle(xx*5, yy*5, 5, 5)
		}
	}
	surface.WriteToPNG("test.png")

}
