package sn

import "testing"

func TestSingleNumber(t *testing.T) {
	testCases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
	}
	expected := []int{1, 4}
	testFuncs := []func([]int) int{
		singleNumber, singleNumber1,
	}

	for _, testFunc := range testFuncs {
		for index, nums := range testCases {
			if res := testFunc(nums); res != expected[index] {
				t.Errorf("expected %d, got %d", expected[index], res)
			}
		}
	}
}
