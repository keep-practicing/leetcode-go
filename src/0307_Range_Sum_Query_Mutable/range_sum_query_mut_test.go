package rsqm

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{1, 2, 3, 6, 7, 8, 9, 3, 4, 2, 5}
	na := Constructor(nums)
	if res := na.SumRange(0, 8); res != 43 {
		t.Errorf("expected %d, got %d", 43, res)
	}
	na.Update(6, 10)
	if res := na.SumRange(0, 8); res != 44 {
		t.Errorf("expected %d, got %d", 44, res)
	}
}
