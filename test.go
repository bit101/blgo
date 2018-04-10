package main

import "bitlib/surface"

func main() {
	surface := surface.NewBitSurface(500, 500)
	surface.ClearRGB(1, 0, 0)
	surface.Line(0, 0, 500, 500)
	surface.Line(500, 0, 0, 500)
	surface.Grid(0, 0, 500, 500, 10, 40)
	surface.WriteToPNG("test.png")

}
