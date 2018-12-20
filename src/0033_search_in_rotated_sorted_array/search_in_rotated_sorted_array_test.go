package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	type arg struct {
		nums   []int
		target int
	}

	testCases := []arg{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
		{nums: []int{}, target: 2},
		{nums: []int{4}, target: 4},
		{nums: []int{4, 5, 6, 7, 0}, target: 0},
		{nums: []int{9, 7}, target: 7},
		{nums: []int{1, 3}, target: 3},
	}

	expected := []int{4, -1, -1, 0, 4, 1, 1}

	for index, data := range testCases {
		if res := search(data.nums, data.target); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
