package p3d

import (
	"math"

	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
)

type Point3D struct {
	X, Y, Z float64
}

func NewPoint3D(x, y, z float64) *Point3D {
	return &Point3D{x, y, z}
}

func LerpPoints(t float64, p0 *Point3D, p1 *Point3D) *Point3D {
	return NewPoint3D(
		blmath.Lerp(t, p0.X, p1.X),
		blmath.Lerp(t, p0.Y, p1.Y),
		blmath.Lerp(t, p0.Z, p1.Z),
	)
}

func LerpSinPoints(t float64, p0 *Point3D, p1 *Point3D) *Point3D {
	return NewPoint3D(
		blmath.LerpSin(t, p0.X, p1.X),
		blmath.LerpSin(t, p0.Y, p1.Y),
		blmath.LerpSin(t, p0.Z, p1.Z),
	)
}

func RandomPoint3D(minX, maxX, minY, maxY, minZ, maxZ float64) *Point3D {
	return &Point3D{
		random.FloatRange(minX, maxX),
		random.FloatRange(minY, maxY),
		random.FloatRange(minZ, maxZ),
	}
}

func RandomPointOnSphere(x, y, z, radius float64) *Point3D {
	p := RandomPoint3D(-1, 1, -1, 1, -1, 1)
	return p.Norm().Scale(radius)
}

func (p *Point3D) Mag() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p *Point3D) Project(z, fl float64) (float64, float64, float64) {
	scale := fl / (fl + p.Z + z)
	return p.X * scale, p.Y * scale, scale
}

func (p *Point3D) Norm() *Point3D {
	mag := p.Mag()
	return NewPoint3D(p.X/mag, p.Y/mag, p.Z/mag)
}

func (p *Point3D) RotateX(angle float64) *Point3D {
	s := math.Sin(angle)
	c := math.Cos(angle)
	y := c*p.Y - s*p.Z
	z := c*p.Z + s*p.Y
	return NewPoint3D(p.X, y, z)
}

func (p *Point3D) RotateY(angle float64) *Point3D {
	s := math.Sin(angle)
	c := math.Cos(angle)
	x := c*p.X - s*p.Z
	z := c*p.Z + s*p.X
	return NewPoint3D(x, p.Y, z)
}

func (p *Point3D) RotateZ(angle float64) *Point3D {
	s := math.Sin(angle)
	c := math.Cos(angle)
	x := c*p.X - s*p.Y
	y := c*p.Y + s*p.X
	return NewPoint3D(x, y, p.Z)
}

func (p *Point3D) Rotate(x, y, z float64) *Point3D {
	p1 := p.RotateX(x)
	p1 = p1.RotateY(y)
	p1 = p1.RotateZ(z)
	return p1
}

func (p *Point3D) Translate(x, y, z float64) *Point3D {
	return NewPoint3D(p.X+x, p.Y+y, p.Z+z)
}

func (p *Point3D) TranslateX(x float64) *Point3D {
	return NewPoint3D(p.X+x, p.Y, p.Z)
}

func (p *Point3D) TranslateY(y float64) *Point3D {
	return NewPoint3D(p.X, p.Y+y, p.Z)
}

func (p *Point3D) TranslateZ(z float64) *Point3D {
	return NewPoint3D(p.X, p.Y, p.Z+z)
}

func (p *Point3D) Scale(s float64) *Point3D {
	return NewPoint3D(p.X*s, p.Y*s, p.Z*s)
}

func (p *Point3D) ScaleX(s float64) *Point3D {
	return NewPoint3D(p.X*s, p.Y, p.Z)
}

func (p *Point3D) ScaleY(s float64) *Point3D {
	return NewPoint3D(p.X, p.Y*s, p.Z)
}

func (p *Point3D) ScaleZ(s float64) *Point3D {
	return NewPoint3D(p.X, p.Y, p.Z*s)
}
