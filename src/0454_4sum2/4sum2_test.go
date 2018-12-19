package foursum2

import "testing"

func TestFourSumCount(t *testing.T) {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}

	expectedData := 2

	if res := fourSumCount(A, B, C, D); res != expectedData {
		t.Errorf("expected %d, got %d", expectedData, res)
	}
}
