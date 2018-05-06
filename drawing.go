package blg

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/geom"
)

////////////////////////////////////////
// Line
////////////////////////////////////////

// Line draws a line between two x, y points.
func (s *Surface) Line(x0, y0, x1, y1 float64) {
	s.MoveTo(x0, y0)
	s.LineTo(x1, y1)
	s.Stroke()
}

// LineThrough draws a line through two x, y points.
func (s *Surface) LineThrough(x0, y0, x1, y1, overlap float64) {
	s.Save()
	s.Translate(x0, y0)
	s.Rotate(math.Atan2(y1-y0, x1-x0))
	p2 := math.Hypot(x0-x1, y0-y1)

	s.MoveTo(-overlap, 0)
	s.LineTo(p2+overlap, 0)
	s.Stroke()
	s.Restore()
}

////////////////////////////////////////
// Ray
////////////////////////////////////////

// Ray draws a line at an angle.
func (s *Surface) Ray(x, y, angle, offset, length float64) {
	s.Save()
	s.Translate(x, y)
	s.Rotate(angle)
	s.MoveTo(offset, 0)
	s.LineTo(offset+length, 0)
	s.Stroke()
	s.Restore()
}

////////////////////////////////////////
// Rectangle
////////////////////////////////////////

// FillRectangle draws a filled rectancle.
func (s *Surface) FillRectangle(x, y, w, h float64) {
	s.Rectangle(x, y, w, h)
	s.Fill()
}

// StrokeRectangle draws a stroked rectangle.
func (s *Surface) StrokeRectangle(x, y, w, h float64) {
	s.Rectangle(x, y, w, h)
	s.Stroke()
}

////////////////////////////////////////
// RoundRectangle
////////////////////////////////////////

// RoundRectangle draws a rounded rectangle.
func (s *Surface) RoundRectangle(x, y, w, h, r float64) {
	s.MoveTo(x+r, y)
	s.LineTo(x+w-r, y)
	s.Arc(x+w-r, y+r, r, -blmath.HalfPi, 0.0)
	s.LineTo(x+w, y+h-r)
	s.Arc(x+w-r, y+h-r, r, 0.0, blmath.HalfPi)
	s.LineTo(x+r, y+h)
	s.Arc(x+r, y+h-r, r, blmath.HalfPi, math.Pi)
	s.LineTo(x, y+r)
	s.Arc(x+r, y+r, r, math.Pi, -blmath.HalfPi)
}

// StrokeRoundRectangle draws a stroked, rounded rectangle.
func (s *Surface) StrokeRoundRectangle(x, y, w, h, r float64) {
	s.RoundRectangle(x, y, w, h, r)
	s.Stroke()
}

// FillRoundRectangle draws a filled, rounded rectangle.
func (s *Surface) FillRoundRectangle(x, y, w, h, r float64) {
	s.RoundRectangle(x, y, w, h, r)
	s.Fill()
}

////////////////////////////////////////
// Circle
////////////////////////////////////////

// Circle draws a circle
func (s *Surface) Circle(x, y, r float64) {
	s.Arc(x, y, r, 0.0, blmath.TwoPi)
}

// FillCircle draws a filled circle.
func (s *Surface) FillCircle(x, y, r float64) {
	s.Circle(x, y, r)
	s.Fill()
}

// StrokeCircle draws a stroked circle.
func (s *Surface) StrokeCircle(x, y, r float64) {
	s.Circle(x, y, r)
	s.Stroke()
}

////////////////////////////////////////
// Ellipse
////////////////////////////////////////

// Ellipse draws an ellipse.
func (s *Surface) Ellipse(x, y, xr, yr float64) {
	s.Save()
	s.Translate(x, y)
	s.Scale(xr, yr)
	s.Circle(0.0, 0.0, 1.0)
	s.Restore()
}

// FillEllipse draws a filled ellipse.
func (s *Surface) FillEllipse(x, y, xr, yr float64) {
	s.Ellipse(x, y, xr, yr)
	s.Fill()
}

// StrokeEllipse draws a stroked ellipse.
func (s *Surface) StrokeEllipse(x, y, xr, yr float64) {
	s.Ellipse(x, y, xr, yr)
	s.Stroke()
}

////////////////////////////////////////
// Path
////////////////////////////////////////

// Path draws a path of points.
func (s *Surface) Path(points []*geom.Point) {
	for _, point := range points {
		s.LineTo(point.X, point.Y)
	}
}

// FillPath draws a filled path of points.
func (s *Surface) FillPath(points []*geom.Point) {
	s.Path(points)
	s.Fill()
}

// StrokePath draws a stroked path of points.
func (s *Surface) StrokePath(points []*geom.Point, close bool) {
	s.Path(points)
	if close {
		s.ClosePath()
	}
	s.Stroke()
}

////////////////////////////////////////
// Polygon
////////////////////////////////////////

// Polygon draws a polygon.
func (s *Surface) Polygon(x, y, r float64, sides int, rotation float64) {
	s.Save()
	s.Translate(x, y)
	s.Rotate(rotation)
	s.MoveTo(r, 0.0)
	for i := 0; i < sides; i++ {
		angle := blmath.TwoPi / float64(sides) * float64(i)
		s.LineTo(math.Cos(angle)*r, math.Sin(angle)*r)
	}
	s.LineTo(r, 0.0)
	s.Restore()
}

// StrokePolygon draws a stroked polygon.
func (s *Surface) StrokePolygon(x, y, r float64, sides int, rotation float64) {
	s.Polygon(x, y, r, sides, rotation)
	s.Stroke()
}

// FillPolygon draws a filled polygon.
func (s *Surface) FillPolygon(x, y, r float64, sides int, rotation float64) {
	s.Polygon(x, y, r, sides, rotation)
	s.Fill()
}

////////////////////////////////////////
// Star
////////////////////////////////////////

// Star draws a star.
func (s *Surface) Star(x, y, r0, r1 float64, points int, rotation float64) {
	s.Save()
	s.Translate(x, y)
	s.Rotate(rotation)
	for i := 0; i < points*2; i++ {
		r := r1
		if i%2 == 1 {
			r = r0
		}
		angle := math.Pi / float64(points) * float64(i)
		s.LineTo(math.Cos(angle)*r, math.Sin(angle)*r)
	}
	s.ClosePath()
	s.Restore()
}

// StrokeStar draws a stroked star.
func (s *Surface) StrokeStar(x, y, r0, r1 float64, points int, rotation float64) {
	s.Star(x, y, r0, r1, points, rotation)
	s.Stroke()
}

// FillStar draws a filled star.
func (s *Surface) FillStar(x, y, r0, r1 float64, points int, rotation float64) {
	s.Star(x, y, r0, r1, points, rotation)
	s.Fill()
}

////////////////////////////////////////
// Splat
////////////////////////////////////////

// Splat draws a splat.
func (s *Surface) Splat(
	x, y float64,
	numNodes int,
	radius, innerRadius, variation float64,
) {
	var points []*geom.Point
	slice := blmath.TwoPi / float64(numNodes*2)
	angle := 0.0
	curve := 0.3
	radiusRange := radius - innerRadius
	variation = blmath.Clamp(variation, 0.0, 1.0)
	for i := 0; i < numNodes; i++ {
		radius := radius + variation*(rand.Float64()*radiusRange*2.0-radiusRange)
		radiusRange := radius - innerRadius
		points = append(points, makePoint(angle-slice*(1.0+curve), innerRadius))
		points = append(points, makePoint(angle+slice*curve, innerRadius))
		points = append(points, makePoint(angle-slice*curve, innerRadius+radiusRange*0.8))
		points = append(points, makePoint(angle+slice/2.0, radius))
		points = append(points, makePoint(angle+slice*(1.0+curve), innerRadius+radiusRange*0.8))
		angle += slice * 2.0
	}

	s.Save()
	s.Translate(x, y)
	s.MultiLoop(points)
	s.Restore()

}

func makePoint(angle, radius float64) *geom.Point {
	return geom.NewPoint(
		math.Cos(angle)*radius,
		math.Sin(angle)*radius,
	)
}

// StrokeSplat draws a stroked splat
func (s *Surface) StrokeSplat(
	x, y float64,
	numNodes int,
	radius, innerRadius, variation float64,
) {
	s.Splat(x, y, numNodes, radius, innerRadius, variation)
	s.Stroke()
}

// FillSplat draws a filled splat.
func (s *Surface) FillSplat(
	x, y float64,
	numNodes int,
	radius, innerRadius, variation float64,
) {
	s.Splat(x, y, numNodes, radius, innerRadius, variation)
	s.Fill()
}

////////////////////////////////////////
// FractalLine
////////////////////////////////////////

// FractalLine draws a fractal line.
func (s *Surface) FractalLine(x1, y1, x2, y2, roughness float64, iterations int) {
	dx := x2 - x1
	dy := y2 - y1
	offset := math.Sqrt(dx*dx+dy*dy) * 0.15

	var path []*geom.Point
	path = append(path, geom.NewPoint(x1, y1))
	path = append(path, geom.NewPoint(x2, y2))

	for i := 0; i < iterations; i++ {
		var newPath []*geom.Point
		for j, point := range path {
			newPath = append(newPath, geom.NewPoint(point.X, point.Y))
			if j < len(path)-1 {
				x := (point.X+path[j+1].X)/2.0 + rand.Float64()*offset*2.0 - offset
				y := (point.Y+path[j+1].Y)/2.0 + rand.Float64()*offset*2.0 - offset
				newPath = append(newPath, geom.NewPoint(x, y))
			}
		}
		offset *= roughness
		path = newPath
	}
	s.Path(path)
}

func (s *Surface) strokeFractalLine(x1, y1, x2, y2, roughness float64, iterations int) {
	s.FractalLine(x1, y1, x2, y2, roughness, iterations)
	s.Stroke()
}

////////////////////////////////////////
// Heart
////////////////////////////////////////

// Heart draws a heart shape.
func (s *Surface) Heart(x, y, w, h, r float64) {
	s.Save()
	s.Translate(x, y)
	s.Rotate(r)
	var path []*geom.Point
	res := math.Sqrt(w * h)
	for i := 0; i < int(res); i++ {
		a := blmath.TwoPi * float64(i) / res
		x := w * math.Pow(math.Sin(a), 3.0)
		y := h*(0.8125*math.Cos(a)) - 0.3125*math.Cos(2.0*a) - 0.125*math.Cos(3.0*a) - 0.0625*math.Cos(4.0*a)
		path = append(path, geom.NewPoint(x, -y))
	}
	s.Path(path)
	s.Restore()
}

// FillHeart draws a filled heart shape.
func (s *Surface) FillHeart(x, y, w, h, r float64) {
	s.Heart(x, y, w, h, r)
	s.Fill()
}

// StrokeHeart draws a stroked heart shape.
func (s *Surface) StrokeHeart(x, y, w, h, r float64) {
	s.Heart(x, y, w, h, r)
	s.Stroke()
}

////////////////////////////////////////
// Points
////////////////////////////////////////

// Points draws a number of points.
func (s *Surface) Points(points []*geom.Point, radius float64) {
	for _, point := range points {
		s.FillCircle(point.X, point.Y, radius)
	}
}

////////////////////////////////////////
// CurveTo
////////////////////////////////////////

// StrokeCurveTo draws a stroked curve.
func (s *Surface) StrokeCurveTo(x0, y0, x1, y1, x2, y2 float64) {
	s.CurveTo(x0, y0, x1, y1, x2, y2)
	s.Stroke()
}

////////////////////////////////////////
// QuadraticCurveTo
////////////////////////////////////////

// QuadraticCurveTo draws a quadratic curve to two points.
func (s *Surface) QuadraticCurveTo(x0, y0, x1, y1 float64) {
	px, py := s.GetCurrentPoint()
	s.CurveTo(
		2.0/3.0*x0+1.0/3.0*px,
		2.0/3.0*y0+1.0/3.0*py,
		2.0/3.0*x0+1.0/3.0*x1,
		2.0/3.0*y0+1.0/3.0*y1,
		x1, y1,
	)
}

// StrokeQuadraticCurveTo draws a stroked quadratic curve.
func (s *Surface) StrokeQuadraticCurveTo(x0, y0, x1, y1 float64) {
	s.QuadraticCurveTo(x0, y0, x1, y1)
	s.Stroke()
}

////////////////////////////////////////
// MultiCurve
////////////////////////////////////////

// MultiCurve draws a smooth curve between a set of points.
func (s *Surface) MultiCurve(points []*geom.Point) {
	s.MoveTo(points[0].X, points[0].Y)
	s.LineTo(
		(points[0].X+points[1].X)/2.0,
		(points[0].Y+points[1].Y)/2.0,
	)
	i := 1
	for i < len(points)-1 {
		fmt.Println(i, points[i], points[i+1])
		p0 := points[i]
		p1 := points[i+1]
		midx := (p0.X + p1.X) / 2.0
		midy := (p0.Y + p1.Y) / 2.0
		s.QuadraticCurveTo(p0.X, p0.Y, midx, midy)
		i = i + 1

	}
	p := points[len(points)-1]
	s.LineTo(p.X, p.Y)
}

// StrokeMultiCurve draws a stroked curve between a set of points.
func (s *Surface) StrokeMultiCurve(points []*geom.Point) {
	s.MultiCurve(points)
	s.Stroke()
}

////////////////////////////////////////
// MultiLoop
////////////////////////////////////////

// MultiLoop draws a smooth, closed curve between a set of points.
func (s *Surface) MultiLoop(points []*geom.Point) {
	pA := points[0]
	pZ := points[len(points)-1]
	mid1x := (pZ.X + pA.X) / 2.0
	mid1y := (pZ.Y + pA.Y) / 2.0
	s.MoveTo(mid1x, mid1y)
	for i := 0; i < len(points)-1; i++ {
		p0 := points[i]
		p1 := points[i+1]
		midx := (p0.X + p1.X) / 2.0
		midy := (p0.Y + p1.Y) / 2.0
		s.QuadraticCurveTo(p0.X, p0.Y, midx, midy)
	}
	s.QuadraticCurveTo(pZ.X, pZ.Y, mid1x, mid1y)
}

// FillMultiLoop draws a filled, smooth, closed curve between a set of points.
func (s *Surface) FillMultiLoop(points []*geom.Point) {
	s.MultiLoop(points)
	s.Fill()
}

// StrokeMultiLoop draws a stroked, smooth, closed curve between a set of points.
func (s *Surface) StrokeMultiLoop(points []*geom.Point) {
	s.MultiLoop(points)
	s.Stroke()
}

////////////////////////////////////////
// Grid
////////////////////////////////////////

// Grid draws a grid.
func (s *Surface) Grid(x, y, w, h, xres, yres float64) {
	xx := x
	yy := y
	for xx <= x+w {
		s.MoveTo(xx, y)
		s.LineTo(xx, y+h)
		xx += xres
	}
	for yy <= y+h {
		s.MoveTo(x, yy)
		s.LineTo(x+w, yy)
		yy += yres
	}
	s.Stroke()
}
