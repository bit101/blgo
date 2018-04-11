package bitmath

import (
	"testing"
)

func TestNorm(t *testing.T) {
	if Norm(0, 0, 100) != 0 {
		t.Error("Norm(0, 0, 100) != 0")
	}
	if Norm(100, 0, 100) != 1 {
		t.Error("Norm(100, 0, 1) != 1")
	}
	if Norm(50, 0, 100) != 0.5 {
		t.Error("Norm(50, 0, 100) != 0.5")
	}
	if Norm(-50, 0, 100) != -0.5 {
		t.Error("Norm(-50, 0, 100) != -0.5")
	}
	if Norm(150, 0, 100) != 1.5 {
		t.Error("Norm(150, 0, 100) != 1.5")
	}
	if Norm(0, 100, 0) != 1 {
		t.Error("Norm(0, 100, 0) != 1")
	}
	if Norm(100, 100, 0) != 0 {
		t.Error("Norm(1, 100, 0) != 0")
	}

}
