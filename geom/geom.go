package geom

import (
	"errors"
	"math"
)

// Distance returns the distance between two x,y positions
func Distance(x0, y0, x1, y1 float64) float64 {
	return math.Hypot(x1-x0, y1-y0)
}

// DotProduct returns the dot product between two lines.
func DotProduct(p0 *Point, p1 *Point, p2 *Point, p3 *Point) float64 {
	dx0 := p1.X - p0.X
	dy0 := p1.Y - p0.Y
	dx1 := p3.X - p2.X
	dy1 := p3.Y - p2.Y
	return dx0*dx1 + dy0*dy1
}

// AngleBetween returns the angle between two lines.
func AngleBetween(p0 *Point, p1 *Point, p2 *Point, p3 *Point) float64 {
	dp := DotProduct(p0, p1, p2, p3)
	mag0 := p0.Distance(p1)
	mag1 := p2.Distance(p3)
	return math.Acos(dp / mag0 / mag1)
}

// InRect returns whether or not an x, y point is within a rectangle.
func InRect(xp, yp, x, y, w, h float64) bool {
	return xp >= x && xp <= x+w && yp >= y && yp <= y+h
}

// PointInRect returns whether or not an x, y point is within a rectangle.
func PointInRect(p *Point, x, y, w, h float64) bool {
	return InRect(p.X, p.Y, x, y, w, h)
}

// InCircle returns whether or not an x, y point is within a circle.
func InCircle(xp, yp, x, y, r float64) bool {
	dx := xp - x
	dy := yp - y
	return math.Sqrt(dx*dx+dy*dy) <= r
}

// PointInCircle returns whether or not an x, y point is within a circle.
func PointInCircle(p *Point, x, y, r float64) bool {
	return InCircle(p.X, p.Y, x, y, r)
}

// XYToPolar returns an angle and distance to origin from a given x, y location.
func XYToPolar(x, y float64) (float64, float64) {
	return math.Atan2(y, x), math.Hypot(x, y)
}

// BezierPoint calculates a point along a Bezier curve.
func BezierPoint(p0 *Point, p1 *Point, p2 *Point, p3 *Point, t float64) *Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT * oneMinusT
	m1 := 3.0 * oneMinusT * oneMinusT * t
	m2 := 3.0 * oneMinusT * t * t
	m3 := t * t * t
	return &Point{
		m0*p0.X + m1*p1.X + m2*p2.X + m3*p3.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y + m3*p3.Y,
	}
}

// QuadraticPoint calculated a point along a quadratic Bezier curve.
func QuadraticPoint(p0 *Point, p1 *Point, p2 *Point, t float64) *Point {
	oneMinusT := 1.0 - t
	m0 := oneMinusT * oneMinusT
	m1 := 2.0 * oneMinusT * t
	m2 := t * t
	return &Point{
		m0*p0.X + m1*p1.X + m2*p2.X,
		m0*p0.Y + m1*p1.Y + m2*p2.Y,
	}
}

// SegmentIntersect returns whether or not two line segments cross.
func SegmentIntersect(p0 *Point, p1 *Point, p2 *Point, p3 *Point) (*Point, error) {
	a1 := p1.Y - p0.Y
	b1 := p0.X - p1.X
	c1 := a1*p0.X + b1*p0.Y
	a2 := p3.Y - p2.Y
	b2 := p2.X - p3.X
	c2 := a2*p2.X + b2*p2.Y
	denominator := a1*b2 - a2*b1

	if denominator == 0.0 {
		return &Point{}, errors.New("nope")
	}
	intersectX := (b2*c1 - b1*c2) / denominator
	intersectY := (a1*c2 - a2*c1) / denominator
	rx0 := (intersectX - p0.X) / (p1.X - p0.X)
	ry0 := (intersectY - p0.Y) / (p1.Y - p0.Y)
	rx1 := (intersectX - p2.X) / (p3.X - p2.X)
	ry1 := (intersectY - p2.Y) / (p3.Y - p2.Y)

	if ((rx0 >= 0.0 && rx0 <= 1.0) || (ry0 >= 0.0 && ry0 <= 1.0)) &&
		((rx1 >= 0.0 && rx1 <= 1.0) || (ry1 >= 0.0 && ry1 <= 1.0)) {
		return &Point{
			intersectX,
			intersectY,
		}, nil
	}
	return &Point{}, errors.New("nope")
}

// TangentPointToCircle calculates the point at which a line extending from a point will contact a circle.
func TangentPointToCircle(point *Point, circle *Circle, anticlockwise bool) *Point {
	d := point.Distance(circle.Center)
	dir := -1.0
	if anticlockwise {
		dir = 1.0
	}
	angle := math.Cos(-circle.Radius/d) * dir
	baseAngle := math.Atan2(circle.Center.Y-point.Y, circle.Center.X-point.X)
	totalAngle := baseAngle + angle

	return &Point{
		circle.Center.X + math.Cos(totalAngle)*circle.Radius,
		circle.Center.Y + math.Sin(totalAngle)*circle.Radius,
	}
}
