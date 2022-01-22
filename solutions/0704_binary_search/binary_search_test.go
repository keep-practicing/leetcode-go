package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	type arg struct {
		nums   []int
		target int
	}

	testCases := []arg{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: -3,
		},
	}

	expected := []int{4, -1}

	for index, data := range testCases {
		if res := search(data.nums, data.target); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
