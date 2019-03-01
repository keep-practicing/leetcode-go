package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testData := [][]int{
		{-1, 0, 1, 2, -6, -4},
		{-7, 3, 4, 5, 6, 6, 4},
	}
	expected := [][][]int{
		{{-1, 0, 1}},
		{{-7, 3, 4}},
	}

	for index, nums := range testData {
		if res := threeSum(nums); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
