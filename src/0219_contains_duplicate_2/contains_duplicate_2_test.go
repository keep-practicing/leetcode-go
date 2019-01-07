package cond

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 0, 1, 1},
		{1, 2, 3, 1, 2, 3},
	}
	ks := []int{3, 1, 2}
	expected := []bool{true, true, false}

	for index, nums := range testCases {
		if res := containsNearbyDuplicate(nums, ks[index]); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
