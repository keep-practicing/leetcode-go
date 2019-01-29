package rsqm

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{1, 2, 3, 6, 7, 8, 9, 3, 4, 2, 5}
	na := Constructor(nums)
	if res := na.SumRange(4, 9); res != 33 {
		t.Errorf("expected %d, got %d", 43, res)
	}
	na.Update(6, 10)
	if res := na.SumRange(4, 9); res != 34 {
		t.Errorf("expected %d, got %d", 44, res)
	}
}
