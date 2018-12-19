package triangle

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	testData := [][][]int{
		[][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}},
		[][]int{},
		[][]int{[]int{-1}, []int{2, 3}, []int{1, -1, -3}},
	}

	expectedData := []int{11, 0, -1}

	for index, data := range testData {
		if mininum := minimumTotal(data); mininum != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], mininum)
		}
	}
}
