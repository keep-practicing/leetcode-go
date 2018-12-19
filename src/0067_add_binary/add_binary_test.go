package addbinary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	type arg struct {
		a string
		b string
	}

	testCases := []arg{
		arg{
			a: "11",
			b: "1",
		},
		arg{
			b: "11",
			a: "1",
		},
	}

	expected := []string{
		"100", "100",
	}

	for index, data := range testCases {
		if res := addBinary(data.a, data.b); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
