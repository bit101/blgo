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

// FromPolar creates a new point from and angle and radius.
func FromPolar(angle float64, radius float64) Point {
	return Point{math.Cos(angle) * radius, math.Sin(angle) * radius}
}

// Distance between this point and another point
func (p0 *Point) Distance(p1 *Point) float64 {
	return math.Hypot(p0.X-p1.X, p0.Y-p1.Y)
}

// Magnitude is distance from origin to this point
func (p *Point) Magnitude() float64 {
	return math.Hypot(p.X, p.Y)
}

func (p *Point) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

func (p *Point) Translate(x float64, y float64) {
	p.X += x
	p.Y += y
}

func (p *Point) Scale(scaleX float64, scaleY float64) {
	p.X *= scaleX
	p.Y *= scaleY
}

func (p *Point) Rotate(angle float64) {
	x := p.X*math.Cos(angle) + p.Y*math.Sin(angle)
	y := p.Y*math.Cos(angle) - p.X*math.Sin(angle)
	p.X = x
	p.Y = y
}
