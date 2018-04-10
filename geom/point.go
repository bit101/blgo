package geom

import "math"

// Point 2d point
type Point struct {
	X float64
	Y float64
}

// NewPoint creates a new 2d point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// Distance between this point and another point
func (p0 *Point) Distance(p1 *Point) float64 {
	return math.Hypot(p0.X-p1.X, p0.Y-p1.Y)
}

// Magnitude is distance from origin to this point
func (p *Point) Magnitude() float64 {
	return math.Hypot(p.X, p.Y)
}
