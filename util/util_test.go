package util

import (
	"testing"
)

func TestNorm(t *testing.T) {
	result := ParentDir()
	expected := "util"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
