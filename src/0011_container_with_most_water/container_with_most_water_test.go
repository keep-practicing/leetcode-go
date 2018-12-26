package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{},
		{1},
		{1, 2},
		{2, 3, 5, 7, 8, 9, 5, 3, 2},
	}
	expected := []int{49, 0, 0, 1, 20}
	for index, data := range testCases {
		if res := maxArea(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}

func TestMaxArea1(t *testing.T) {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{},
		{1},
		{1, 2},
		{2, 3, 5, 7, 8, 9, 5, 3, 2},
	}
	expected := []int{49, 0, 0, 1, 20}
	for index, data := range testCases {
		if res := maxArea1(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}

func TestMaxArea2(t *testing.T) {
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{},
		{1},
		{1, 2},
		{2, 3, 5, 7, 8, 9, 5, 3, 2},
	}
	expected := []int{49, 0, 0, 1, 20}
	for index, data := range testCases {
		if res := maxArea2(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
