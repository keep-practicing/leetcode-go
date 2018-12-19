package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	if res := twoSum(nums, 9); !reflect.DeepEqual(res, []int{0, 1}) {
		t.Error("Failed, two sum")
	}

	if res := twoSum(nums, 6); !reflect.DeepEqual(res, []int{}) {
		t.Error("Failed, two sum")
	}
}
