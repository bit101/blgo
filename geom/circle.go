package geom

// Circle is a struct for holding a circle.
type Circle struct {
	Center *Point
	Radius float64
}

// NewCircle creates a new circle.
func NewCircle(x float64, y float64, r float64) *Circle {
	return &Circle{
		NewPoint(x, y),
		r,
	}
}
