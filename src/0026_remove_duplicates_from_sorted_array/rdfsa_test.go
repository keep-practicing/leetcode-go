package rdfsa

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	// removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{},
	}

	expected := []int{2, 5, 0}

	for index, data := range testCases {
		if res := removeDuplicates(data); expected[index] != res {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
