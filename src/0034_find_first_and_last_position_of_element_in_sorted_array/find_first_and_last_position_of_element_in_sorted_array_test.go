package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"runtime"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type arg struct {
		nums   []int
		target int
	}

	testCases := []arg{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
		},
		{
			nums:   []int{1},
			target: 1,
		},
		{
			nums:   []int{},
			target: 0,
		},
		{
			nums:   []int{2, 2},
			target: 2,
		},
		{
			nums:   []int{1},
			target: 0,
		},
	}

	expected := [][]int{
		{3, 4},
		{-1, -1},
		{0, 0},
		{-1, -1},
		{0, 1},
		{-1, -1},
	}

	testFuncs := []func([]int, int) []int{
		searchRange,
		searchRange1,
		searchRange2,
	}

	for _, testFunc := range testFuncs {
		for index, testData := range testCases {
			if res := testFunc(testData.nums, testData.target); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("function %s, expected %v, got %v", runtime.FuncForPC(reflect.ValueOf(testFunc).Pointer()).Name(), expected[index], res)
			}
		}
	}
}
