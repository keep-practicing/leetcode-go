package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testData := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}
	expected := [][][]int{
		{{-1, 0, 1}, {-1, -1, 2}},
	}

	for index, nums := range testData {
		if res := threeSum(nums); reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
