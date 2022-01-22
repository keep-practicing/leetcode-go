package cd

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
	}
	expected := []bool{true, false}

	for index, nums := range testCases {
		if res := containsDuplicate(nums); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
