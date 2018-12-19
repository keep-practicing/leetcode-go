package largestnumberatleasttwiceofothers

import "testing"

func TestDominantIndex(t *testing.T) {
	testCases := [][]int{
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{1, 2, 3, 4},
		{3, 6, 1, 0},
	}
	expected := []int{3, 2, -1, 1}
	for index, data := range testCases {
		if res := dominantIndex(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
