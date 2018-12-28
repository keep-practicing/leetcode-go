package rsqim

import "testing"

func TestSumRange(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	ranges := [][]int{
		{0, 2},
		{2, 5},
		{0, 5},
	}

	expected := []int{1, -1, -3}

	obj := Constructor(nums)

	for index, section := range ranges {
		if res := obj.SumRange(section[0], section[1]); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
