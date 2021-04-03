package geom

import (
	"math"
	"testing"
)

func TestDotProduct(t *testing.T) {
	p0 := NewPoint(0, -1)
	p1 := NewPoint(0, 1)
	p2 := NewPoint(-1, 0)
	p3 := NewPoint(1, 0)

	result := DotProduct(p0, p1, p2, p3)
	if result != 0 {
		t.Errorf("DotProduct of 90 degrees != 0")
	}

	p0.X = -0.1
	result = DotProduct(p0, p1, p2, p3)
	if result <= 0 {
		t.Errorf("DotProduct of > 90 degrees not > 0")
	}

	p0.X = 0.1
	result = DotProduct(p0, p1, p2, p3)
	if result >= 0 {
		t.Errorf("DotProduct of < 90 degrees not < 0")
	}
}

func TestAngleBetween(t *testing.T) {
	p0 := NewPoint(0, -1)
	p1 := NewPoint(0, 0)
	p2 := NewPoint(1, 0)

	result := AngleBetween(p1, p0, p1, p2)
	if result != math.Pi/2 {
		t.Errorf("angle between points not 90 degrees")
	}

	p0.X = -0.1
	result = AngleBetween(p1, p0, p1, p2)
	if result <= math.Pi/2 {
		t.Errorf("angle between points not > 90 degrees")
	}

	p0.X = 0.1
	result = AngleBetween(p1, p0, p1, p2)
	if result >= math.Pi/2 {
		t.Errorf("angle between points not < 90 degrees")
	}
}

func TestLerpPoint(t *testing.T) {
	p0 := NewPoint(0, 0)
	p1 := NewPoint(100, -200)
	var tests = []struct {
		t     float64
		wantX float64
		wantY float64
	}{
		{0, 0, 0},
		{1, 100, -200},
		{0.5, 50, -100},
	}

	for _, test := range tests {
		result := LerpPoint(test.t, p0, p1)
		if result.X != test.wantX || result.Y != test.wantY {
			t.Errorf("LerpPoint(%f, %v, %v) != %f, %f", test.t, p0, p1, test.wantX, test.wantY)
		}
	}

}

func TestMapRectangle(t *testing.T) {
	r0 := NewRectangle(0, 0, 100, 100)
	r1 := NewRectangle(200, 200, 400, 400)
	var tests = []struct {
		r0    *Rectangle
		r1    *Rectangle
		x     float64
		y     float64
		wantX float64
		wantY float64
	}{
		{r0, r1, 0, 0, 200, 200},
		{r0, r1, 100, 100, 600, 600},
		{r0, r1, 0, 100, 200, 600},
		{r0, r1, 100, 0, 600, 200},
		{r0, r1, 50, 50, 400, 400},
		{r0, r1, 25, 75, 300, 500},
	}

	for _, test := range tests {
		x, y := MapRectangle(test.x, test.y, test.r0, test.r1)
		if x != test.wantX || y != test.wantY {
			t.Errorf("MapRectangle(%f, %f, %v, %v) != %f, %f", test.x, test.y, test.r0, test.r1, x, y)
		}
	}
}
