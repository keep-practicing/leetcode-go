package spltk

import "testing"

func TestNumSubArrayProductLessThanK(t *testing.T) {
	testCases := [][]int{
		{10, 5, 2, 6},
		{10, 5, 2, 6},
		{10, 5, 100, 6},
	}

	ks := []int{100, 0, 100}
	expected := []int{8, 0, 4}
	for index, data := range testCases {
		if res := numSubArrayProductLessThanK(data, ks[index]); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
