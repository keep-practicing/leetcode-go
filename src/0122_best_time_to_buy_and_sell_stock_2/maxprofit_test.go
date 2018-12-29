package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
		{}, {1},
	}
	expected := []int{7, 4, 0, 0, 0}

	for index, data := range testCases {
		if res := maxProfit(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
