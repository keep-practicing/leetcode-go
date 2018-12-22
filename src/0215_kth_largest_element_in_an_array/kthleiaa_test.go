package kthleiaa

import "testing"

func TestFindKthLargest(t *testing.T) {
	type arg struct {
		nums []int
		k    int
	}

	testCases := []arg{
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4},
	}

	expected := []int{5, 4}

	for index, data := range testCases {
		if res := findKthLargest(data.nums, data.k); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
