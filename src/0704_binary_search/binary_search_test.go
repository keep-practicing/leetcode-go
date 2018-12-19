package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	expected := 4

	if res := search(nums, target); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
