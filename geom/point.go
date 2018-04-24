package geom

import (
	"math"

	"github.com/bit101/bitlibgo/random"
)

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

// RandomPoint returns a point within the rectangle defined in the params x, y, w, h.
func RandomPoint(x, y, w, h float64) *Point {
	return &Point{
		X: random.FloatRange(x, x+w),
		Y: random.FloatRange(y, y+h),
	}
}

// FromPolar creates a new point from and angle and radius.
func FromPolar(angle float64, radius float64) Point {
	return Point{math.Cos(angle) * radius, math.Sin(angle) * radius}
}

// Distance between this point and another point
func (p *Point) Distance(p1 *Point) float64 {
	return math.Hypot(p.X-p1.X, p.Y-p1.Y)
}

// Magnitude is distance from origin to this point
func (p *Point) Magnitude() float64 {
	return math.Hypot(p.X, p.Y)
}

// Angle returns the angle from the origin to this point.
func (p *Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

// Translate moves this point on the x and y axes.
func (p *Point) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}

// Scale scales this point on the x and y axes.
func (p *Point) Scale(scaleX float64, scaleY float64) {
	p.X *= scaleX
	p.Y *= scaleY
}

// Rotate rotates this point around the origin.
func (p *Point) Rotate(angle float64) {
	x := p.X*math.Cos(angle) + p.Y*math.Sin(angle)
	y := p.Y*math.Cos(angle) - p.X*math.Sin(angle)
	p.X = x
	p.Y = y
}
