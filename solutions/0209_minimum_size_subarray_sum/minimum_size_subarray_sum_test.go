package minimumsizesubarraysum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type arg struct {
		s    int
		nums []int
	}
	testData := []arg{
		{
			s:    7,
			nums: []int{2, 3, 1, 2, 4, 3},
		},
		{
			s:    7,
			nums: []int{7, 4, 3, 2},
		},
		{
			s:    66,
			nums: []int{7, 4, 3, 2},
		},
	}

	expectedData := []int{2, 1, 0}

	for index, data := range testData {
		if res := minSubArrayLen(data.s, data.nums); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
