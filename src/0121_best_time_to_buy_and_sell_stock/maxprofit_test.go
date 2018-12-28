package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{5, 3, 2, 2, 5, 7, 9, 4},
		{},
		{3},
		{5, 3, 2, 2, 5, 7, 9, 4,5, 3, 2, 2, 5, 7, 9, 4},
		{2,4,1,11,7},
	}

	expected := []int{5, 7, 0, 0,7,10}

	for index, data := range testCases {
		if res := maxProfit(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}

func TestMaxProfit1(t *testing.T) {
	testCases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{5, 3, 2, 2, 5, 7, 9, 4},
		{},
		{3},
	}

	expected := []int{5, 7, 0, 0}

	for index, data := range testCases {
		if res := maxProfit1(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
