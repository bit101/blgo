package blmath

import (
	"math"
	"testing"
)

func TestNorm(t *testing.T) {
	var tests = []struct {
		val  float64
		min  float64
		max  float64
		want float64
	}{
		{0, 0, 100, 0},
		{100, 0, 100, 1},
		{50, 0, 100, 0.5},
		{-50, 0, 100, -0.5},
		{150, 0, 100, 1.5},
		{0, 100, 0, 1},
		{100, 100, 0, 0},
	}
	for _, test := range tests {
		result := Norm(test.val, test.min, test.max)
		if result != test.want {
			t.Errorf("Norm(%f, %f, %f) != %f", test.val, test.min, test.max, test.want)
		}
	}
}

func TestLerp(t *testing.T) {
	var tests = []struct {
		t    float64
		min  float64
		max  float64
		want float64
	}{
		{0, 0, 100, 0},
		{1, 0, 100, 100},
		{0.5, 0, 100, 50},
		{-0.5, 0, 100, -50},
		{1.5, 0, 100, 150},
		{0, 100, 0, 100},
		{1, 100, 0, 0},
	}
	for _, test := range tests {
		result := Lerp(test.t, test.min, test.max)
		if result != test.want {
			t.Errorf("Lerp(%f, %f, %f) != %f", test.t, test.min, test.max, test.want)
		}
	}
}

func TestMap(t *testing.T) {
	var tests = []struct {
		srcVal float64
		srcMin float64
		srcMax float64
		dstMin float64
		dstMax float64
		want   float64
	}{
		{0, 0, 100, 200, 400, 200},
		{100, 0, 100, 200, 400, 400},
		{50, 0, 100, 200, 400, 300},
		{-50, 0, 100, 200, 400, 100},
		{150, 0, 100, 200, 400, 500},
		{0, 100, 0, 200, 400, 400},
		{100, 100, 00, 200, 400, 200},
		{0, 0, 100, 400, 200, 400},
		{100, 0, 100, 400, 200, 200},
	}

	for _, test := range tests {
		result := Map(test.srcVal, test.srcMin, test.srcMax, test.dstMin, test.dstMax)
		if result != test.want {
			t.Errorf("MapTo(%f, %f, %f, %f, %f) != %f", test.srcVal, test.srcMin, test.srcMax, test.dstMin, test.dstMax, test.want)
		}
	}
}

func TestWrap(t *testing.T) {
	var tests = []struct {
		val  float64
		min  float64
		max  float64
		want float64
	}{
		{0, 0, 360, 0},
		{360, 0, 360, 0},
		{400, 0, 360, 40},
		{-50, 0, 360, 310},
		{720, 0, 360, 0},
		{719, 0, 360, 359},
		{-360, 0, 360, 0},
		{50, 0, 360, 50},
	}
	for _, test := range tests {
		result := Wrap(test.val, test.min, test.max)
		if result != test.want {
			t.Errorf("Wrap(%f, %f, %f) != %f", test.val, test.min, test.max, test.want)
		}
	}
}

func TestClamp(t *testing.T) {
	var tests = []struct {
		val  float64
		min  float64
		max  float64
		want float64
	}{
		{0, 0, 100, 0},
		{100, 0, 100, 100},
		{-100, 0, 100, 0},
		{150, 0, 100, 100},
		{0, 100, 0, 0},
		{100, 100, 0, 100},
		{-100, 100, 0, 0},
		{150, 100, 0, 100},
	}
	for _, test := range tests {
		result := Clamp(test.val, test.min, test.max)
		if result != test.want {
			t.Errorf("Clamp(%f, %f, %f) != %f", test.val, test.min, test.max, test.want)
		}
	}
}

func TestEqualish(t *testing.T) {
	var tests = []struct {
		a     float64
		b     float64
		delta float64
		want  bool
	}{
		{0, 0, 0.01, true},              // values equal
		{-10, -10, 0.01, true},          // values equal
		{0.0001, 0.0003, 0.0001, false}, // values 0.2 different
		{0.0003, 0.0001, 0.0001, false}, // values 0.2 different
		{0.1, -0.1, 0.2, true},          // values 0.2 different, cross 0
		{-0.1, 0.1, 0.2, true},          // values 0.2 different, cross 0
		{-0.1000001, 0.1, 0.2, false},   // values 0.2000001 different
		{0, 0, 0, true},                 // values equal, delta 0
		{0, 0, -1, false},               // delta negative, never true
	}
	for _, test := range tests {
		result := Equalish(test.a, test.b, test.delta)
		if result != test.want {
			t.Errorf("Equalish(%f, %f, %f) != %t", test.a, test.b, test.delta, test.want)
		}
	}
}

func TestSinRange(t *testing.T) {
	var tests = []struct {
		val  float64
		min  float64
		max  float64
		want float64
	}{
		{math.Pi * 0.0, 0, 100, 50},
		{math.Pi * 0.5, 0, 100, 100},
		{math.Pi * 1.0, 0, 100, 50},
		{math.Pi * 1.5, 0, 100, 0},
		{math.Pi * 2.0, 0, 100, 50},
	}
	for _, test := range tests {
		result := SinRange(test.val, test.min, test.max)
		if !Equalish(result, test.want, 0.0001) {
			t.Errorf("SinRange(%f, %f, %f) != %f", test.val, test.min, test.max, test.want)
		}
	}
}

func TestCosRange(t *testing.T) {
	var tests = []struct {
		val  float64
		min  float64
		max  float64
		want float64
	}{
		{math.Pi * 0.0, 0, 100, 100},
		{math.Pi * 0.5, 0, 100, 50},
		{math.Pi * 1.0, 0, 100, 0},
		{math.Pi * 1.5, 0, 100, 50},
		{math.Pi * 2.0, 0, 100, 100},
	}
	for _, test := range tests {
		result := CosRange(test.val, test.min, test.max)
		if !Equalish(result, test.want, 0.0001) {
			t.Errorf("CosRange(%f, %f, %f) != %f", test.val, test.min, test.max, test.want)
		}
	}
}

func TestRoundTo(t *testing.T) {
	var tests = []struct {
		value   float64
		decimal int
		want    float64
	}{
		{1234.1234, 4, 1234.1234},
		{1234.1234, 3, 1234.123},
		{1234.1234, 2, 1234.12},
		{1234.1234, 1, 1234.1},
		{1234.1234, 0, 1234},
		{1234.1234, -1, 1230},
		{1234.1234, -2, 1200},
		{1234.1234, -3, 1000},
		{1234.1234, -4, 0},
		{1234.5555, 4, 1234.5555},
		{1234.5555, 3, 1234.556},
		{1234.5555, 2, 1234.56},
		{1234.5555, 1, 1234.6},
		{1234.5555, 0, 1235},
		{1234.5555, -1, 1230},
	}
	for _, test := range tests {
		result := RoundTo(test.value, test.decimal)
		if result != test.want {
			t.Errorf("RoundTo(%f, %d) != %f", test.value, test.decimal, test.want)
		}
	}
}

func TestComplexMag(t *testing.T) {
	var tests = []struct {
		value complex128
		want  float64
	}{
		{complex(3, 4), 5.0},
		{complex(4, 3), 5.0},
		{complex(5, 12), 13.0},
		{complex(12, 5), 13.0},
		{complex(21, 28), 35.0},
		{complex(28, 21), 35.0},
	}
	for _, test := range tests {
		result := ComplexMagnitude(test.value)
		if result != test.want {
			t.Errorf("ComplexMagnitude(%f) != %f", test.value, test.want)
		}
	}
}

func TestComplexImagAbs(t *testing.T) {
	var tests = []struct {
		value complex128
		want  complex128
	}{
		{complex(4, -5), complex(4, 5)},
		{complex(4, 5), complex(4, 5)},
		{complex(-4, -5), complex(-4, 5)},
		{complex(-4, 5), complex(-4, 5)},
		{complex(4.2, -5.3), complex(4.2, 5.3)},
		{complex(4.2, 5.3), complex(4.2, 5.3)},
		{complex(-4.2, -5.3), complex(-4.2, 5.3)},
		{complex(-4.2, 5.3), complex(-4.2, 5.3)},
	}
	for _, test := range tests {
		result := ComplexImagAbs(test.value)
		if result != test.want {
			t.Errorf("ComplexImagAbs(%f) != %f", test.value, test.want)
		}
	}
}
