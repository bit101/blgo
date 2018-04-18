package random

import "testing"

func TestSeed(t *testing.T) {
	Seed(0)
	valueA := Float()

	Seed(1)
	valueB := Float()

	Seed(0)
	valueC := Float()

	if valueA == valueB {
		t.Errorf("%f == %f", valueA, valueB)
	}
	if valueA != valueC {
		t.Errorf("%f != %f", valueA, valueC)
	}
}

func TestRandSeed(t *testing.T) {
	RandSeed()
	valueA := Float()

	RandSeed()
	valueB := Float()

	if valueA == valueB {
		t.Errorf("%f == %f", valueA, valueB)
	}
}

func TestFloat(t *testing.T) {
	value := Float()

	if value < 0.0 || value >= 1.0 {
		t.Errorf("%f not within 0.0 - 1.0", value)
	}
}

func TestInt(t *testing.T) {
	value := Int()

	if value < 0 || value >= 9223372036854775807 {
		t.Errorf("%d not within 0 - 9223372036854775807", value)
	}
}

func TestFloatRange(t *testing.T) {
	value := FloatRange(100, 200)

	if value < 100 || value >= 200 {
		t.Errorf("%f not within 100 - 200", value)
	}
}

func TestIntRange(t *testing.T) {
	value := IntRange(100, 200)

	if value < 100 || value >= 200 {
		t.Errorf("%d not within 100 - 200", value)
	}
}

func TestBool(t *testing.T) {
	value := Boolean()
	// lame test
	if value != true && value != false {
		t.Errorf("%t not boolean", value)
	}
}

func TestWeightedBool(t *testing.T) {
	value := WeightedBool(0.5)
	// can do better here.
	if value != true && value != false {
		t.Errorf("%t not boolean", value)
	}
}
