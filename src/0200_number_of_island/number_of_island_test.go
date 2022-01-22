package noi

import "testing"

func TestNumIslands(t *testing.T) {
	grids := [][][]byte{
		{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		{},
	}
	expected := []int{3, 0}

	for index, grid := range grids {
		if res := numIslands(grid); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}

}
