package motsa

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	type arg struct {
		nums1 []int
		nums2 []int
	}

	testCases := []arg{
		{nums1: []int{1, 3}, nums2: []int{2}},
		{nums1: []int{1, 2}, nums2: []int{3, 4}},
		{nums1: []int{}, nums2: []int{3, 4, 5}},
		{nums1: []int{4, 5, 6}, nums2: []int{}},
		{nums1: []int{2, 3, 4, 5}, nums2: []int{1}},
		{nums1: []int{1}, nums2: []int{2, 3, 4, 5, 6, 7}},
		{nums1: []int{1, 2, 3, 4, 5, 6}, nums2: []int{7}},
	}

	expected := []float64{2, 2.5, 4, 5, 3, 4}

	for index, data := range testCases {
		if res := findMedianSortedArrays(data.nums1, data.nums2); res != expected[index] {
			t.Errorf("expected %f, got %f", expected[index], res)
		}
	}

	defer func() {
		if err := recover(); err == nil || err != "nums1 and nums2 cannot be both empty." {
			t.Errorf("expected err: nums1 and nums2 cannot be both empty.")
		}
	}()
	findMedianSortedArrays([]int{}, []int{})
}
