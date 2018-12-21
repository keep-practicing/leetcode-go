package msa

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type arg struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	testCases := []arg{
		{nums1: []int{1, 2, 3, 7, 0, 0, 0}, m: 4, nums2: []int{2, 5, 6}, n: 3},
		{nums1: []int{7, 8, 9, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3},
	}

	expected := [][]int{{1, 2, 2, 3, 5, 6, 7}, {2, 5, 6, 7, 8, 9}}

	for index, data := range testCases {
		if merge(data.nums1, data.m, data.nums2, data.n); !reflect.DeepEqual(expected[index], data.nums1) {
			t.Errorf("expected %v, got %v", expected[index], data.nums1)
		}
	}
}
