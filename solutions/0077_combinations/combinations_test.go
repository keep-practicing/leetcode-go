package combinations

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	type arg struct {
		n, k int
	}

	testCases := []arg{
		{n: 0, k: 2},
		{n: 2, k: 0},
		{n: 1, k: 5},
		{n: 4, k: 2},
	}
	expected := [][][]int{
		{},
		{},
		{},
		{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
	}

	for index, data := range testCases {
		if res := combine(data.n, data.k); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
