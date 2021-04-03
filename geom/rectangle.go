package geom

import "github.com/bit101/blgo/blmath"

type Rectangle struct {
	X, Y, W, H float64
}

func NewRectangle(x, y, w, h float64) *Rectangle {
	return &Rectangle{x, y, w, h}
}

func MapRectangle(x, y float64, sRect, dRect *Rectangle) (float64, float64) {
	xx := blmath.Map(x, sRect.X, sRect.X+sRect.W, dRect.X, dRect.X+dRect.W)
	yy := blmath.Map(y, sRect.Y, sRect.Y+sRect.H, dRect.Y, dRect.Y+dRect.H)
	return xx, yy
}
