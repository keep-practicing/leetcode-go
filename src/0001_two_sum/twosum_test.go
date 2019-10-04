package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	funcs := []func ([]int, int) []int {twoSum, twoSum1}

	for _, testFunc := range funcs {
		if res := testFunc(nums, 9); !reflect.DeepEqual(res, []int{0, 1}) {
			t.Error("Failed, two sum")
		}
	
		if res := testFunc(nums, 6); !reflect.DeepEqual(res, []int{}) {
			t.Error("Failed, two sum")
		}
	}
}
