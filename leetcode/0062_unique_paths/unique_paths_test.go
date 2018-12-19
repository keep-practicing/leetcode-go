package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	testData := [][2]int{
		{3, 2},
		{51, 9},
		{9, 9},
	}

	expectedData := []int{3, 1916797311, 12870}

	for index, data := range testData {
		if res := uniquePaths(data[0], data[1]); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
