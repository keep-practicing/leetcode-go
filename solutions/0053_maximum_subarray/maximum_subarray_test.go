package maximumsubarray

import "testing"

func TestMaxSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6

	if res := maxSubArray(nums); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
