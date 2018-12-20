package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	testCases := []int{66, 99, 9}
	expected := []int{8, 9, 3}

	for index, data := range testCases {
		if res := mySqrt(data); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
