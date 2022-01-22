package fmirsa

import "testing"

func TestFindMin(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	expected := 1
	if res := findMin(nums); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
