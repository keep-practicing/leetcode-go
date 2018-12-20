package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expected := 5

	if res := removeElement(nums, val); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
