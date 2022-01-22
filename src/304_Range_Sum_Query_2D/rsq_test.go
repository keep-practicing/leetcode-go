package rsq

import "testing"

func TestSunRegion(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)

	testData := [][]int{
		{2, 1, 4, 3},
		{1, 1, 2, 2},
		{1, 2, 2, 4},
	}
	expected := []int{8, 11, 12}

	for index, data := range testData {
		if res := obj.SumRegion(data[0], data[1], data[2], data[3]); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}

	if res := Constructor([][]int{}); res.dp != nil {
		t.Errorf("expected nil, got %v", res.dp)
	}
	if res := Constructor(make([][]int, 3)); res.dp != nil {
		t.Errorf("expected nil, got %v", res.dp)
	}
}
