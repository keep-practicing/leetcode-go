package integerbreak

import "testing"

func TestIntegerBreak(t *testing.T) {
	testData := []int{2, 10, 34}
	expectedData := []int{1, 36, 236196}

	for index, data := range testData {
		if res := integerBreak(data); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
