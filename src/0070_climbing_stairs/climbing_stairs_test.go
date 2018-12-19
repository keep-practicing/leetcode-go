package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	testDatas := []int{0, 1, 2, 3, 34}
	expected := []int{1, 1, 2, 3, 9227465}

	for i, data := range testDatas {
		if steps := climbStairs(data); steps != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], steps)
		}
	}
}
