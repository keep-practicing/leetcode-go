package lcis

import "testing"

func TestFindLengthOfLCIS(t *testing.T) {
	testCases := [][]int{
		{1, 3, 5, 4, 7},
		{2, 2, 2, 2, 2},
		{},
		{2},
	}
	testFuncs := []func([]int) int{
		findLengthOfLCIS,
		findLengthOfLCIS1,
	}

	expected := []int{3, 1, 0, 1}
	for _, tfn := range testFuncs {
		for index, data := range testCases {
			if res := tfn(data); res != expected[index] {
				t.Errorf("expected %d, got %d", expected[index], res)
			}
		}
	}
}
