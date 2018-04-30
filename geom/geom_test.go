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
