package geom

import (
	"bitlib/bitmath"
	"errors"
	"math"
)

func DotProduct(p0 *Point, p1 *Point, p2 *Point, p3 *Point) float64 {
	dx0 := p1.X - p0.X
	dy0 := p1.Y - p0.Y
	dx1 := p3.X - p2.X
	dy1 := p3.Y - p2.Y
	return dx0*dx1 + dy0*dy1
}

func AngleBetween(p0 *Point, p1 *Point, p2 *Point, p3 *Point) float64 {
	dp := DotProduct(p0, p1, p2, p3)
	mag0 := dist(p0, p1)
	mag1 := dist(p2, p3)
	return math.Acos(dp / mag0 / mag1)
}

func dist(p0 *Point, p1 *Point) float64 {
	dx := p1.X - p0.X
	dy := p1.Y - p0.Y
	return math.Hypot(dx, dy)
}

func LerpPoint(p0 Point, p1 Point, t float64) Point {
	return Point{
		bitmath.Lerp(p0.X, p1.X, t),
		bitmath.Lerp(p0.Y, p1.Y, t),
	}
}

func BezierPoint(p0 Point, p1 Point, p2 Point, p3 Point, t float64) Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT * oneMinusT
	m1 := 3.0 * oneMinusT * oneMinusT * t
	m2 := 3.0 * oneMinusT * t * t
	m3 := t * t * t
	return Point{
		m0*p0.X + m1*p1.X + m2*p2.X + m3*p3.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y + m3*p3.Y,
	}
}

func QuadraticPoint(p0 Point, p1 Point, p2 Point, t float64) Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT
	m1 := 2.0 * oneMinusT * t
	m2 := t * t
	return Point{
		m0*p0.X + m1*p1.X + m2*p2.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y,
	}
}

func SegmentIntersect(p0 Point, p1 Point, p2 Point, p3 Point) (Point, error) {
	a1 := p1.Y - p0.Y
	b1 := p0.X - p1.X
	c1 := a1*p0.X + b1*p0.Y
	a2 := p3.Y - p2.Y
	b2 := p2.X - p3.X
	c2 := a2*p2.X + b2*p2.Y
	denominator := a1*b2 - a2*b1

	if denominator == 0.0 {
		return Point{}, errors.New("nope")
	} else {
		intersectX := (b2*c1 - b1*c2) / denominator
		intersectY := (a1*c2 - a2*c1) / denominator
		rx0 := (intersectX - p0.X) / (p1.X - p0.X)
		ry0 := (intersectY - p0.Y) / (p1.Y - p0.Y)
		rx1 := (intersectX - p2.X) / (p3.X - p2.X)
		ry1 := (intersectY - p2.Y) / (p3.Y - p2.Y)

		if ((rx0 >= 0.0 && rx0 <= 1.0) || (ry0 >= 0.0 && ry0 <= 1.0)) &&
			((rx1 >= 0.0 && rx1 <= 1.0) || (ry1 >= 0.0 && ry1 <= 1.0)) {
			return Point{
				intersectX,
				intersectY,
			}, nil
		} else {
			return Point{}, errors.New("nope")
		}
	}
}

func TangentPointToCircle(point *Point, circle *Circle, anticlockwise bool) Point {
	d := dist(point, circle.Center)
	dir := -1.0
	if anticlockwise {
		dir = 1.0
	}
	angle := math.Cos(-circle.Radius/d) * dir
	baseAngle := math.Atan2(circle.Center.Y-point.Y, circle.Center.X-point.X)
	totalAngle := baseAngle + angle

	return Point{
		circle.Center.X + math.Cos(totalAngle)*circle.Radius,
		circle.Center.Y + math.Sin(totalAngle)*circle.Radius,
	}
}
