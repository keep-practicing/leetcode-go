package uniquepaths2

import "testing"

func TestUniquePaths2(t *testing.T) {
	testData := [][][]int{
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
			{0, 1, 0},
		},
		{
			{1},
		},
		{
			{1, 0},
		},
	}

	expectedData := []int{2, 2, 0, 0}

	for index, data := range testData {
		if res := uniquePathsWithObstacles(data); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
