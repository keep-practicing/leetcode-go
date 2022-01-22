package triangle

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	testData := [][][]int{
		{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		{},
		{{-1}, {2, 3}, {1, -1, -3}},
	}

	expectedData := []int{11, 0, -1}

	for index, data := range testData {
		if mininum := minimumTotal(data); mininum != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], mininum)
		}
	}
}
