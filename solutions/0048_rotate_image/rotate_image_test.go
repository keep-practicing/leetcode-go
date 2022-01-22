package ri

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	if rotate(matrix); !reflect.DeepEqual(matrix, expected) {
		t.Errorf("expected %v, got %v", expected, matrix)
	}
}
