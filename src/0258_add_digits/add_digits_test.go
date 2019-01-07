package ad

import "testing"

func TestAddDigits(t *testing.T) {
	num := 38
	expected := 2

	testFuncs := []func(int) int{
		addDigits,
		addDigits1,
	}

	for _, function := range testFuncs {
		if res := function(num); res != expected {
			t.Errorf("expected %d, got %d", expected, res)
		}
	}
}
