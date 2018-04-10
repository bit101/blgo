package main

import (
	"bitlib/color"
	"bitlib/surface"
)

func main() {
	surface := surface.NewBitSurface(500, 500)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			xx := float64(x)
			yy := float64(y)
			surface.SetSourceColor(color.HSV(xx/100*630, 1, 1-yy/100))
			surface.FillRectangle(xx*5, yy*5, 5, 5)
		}
	}
	surface.WriteToPNG("test.png")

}
