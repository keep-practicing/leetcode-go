package rdfsa2

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := [][]int{
		{1, 1, 1, 2, 2, 3},
		{1, 2, 2},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
		{},
		{1},
		{2, 2},
		{2, 2, 2},
	}
	expected := []int{5, 3, 7, 0, 1, 2, 2}

	for index, data := range testCases {
		if res := removeDuplicates(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
