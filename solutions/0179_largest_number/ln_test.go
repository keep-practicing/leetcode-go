package ln

import "testing"

func TestLargestNumber(t *testing.T) {
	testCases := [][]int{
		{10, 2},
		{3, 30, 34, 5, 9},
		{0},
	}

	expected := []string{
		"210",
		"9534330",
		"0",
	}

	for index, nums := range testCases {
		if res := largestNumber(nums); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
