package rsqm

import "testing"

func TestNumArray(t *testing.T) {
	nums := []int{1, 2, 3, 6, 7, 8, 9, 3, 4, 2, 5}
	na := Constructor(nums)
	if res := na.SumRange(0, 10); res != 50 {
		t.Errorf("expected %d, got %d", 50, res)
	}
	na.Update(6, 10)
	if res := na.SumRange(0, 10); res != 51 {
		t.Errorf("expected %d, got %d", 51, res)
	}
}
