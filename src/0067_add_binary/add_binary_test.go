package addbinary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	expected := "100"

	if res := addBinary(a, b); res != expected {
		t.Errorf("expected %s, got %s", expected, res)
	}
}
