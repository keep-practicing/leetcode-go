package onebitandtwobitcharacters

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	testCases := [][]int{
		{1, 0, 0},
		{1, 1, 1, 0},
		{0},
	}
	expected := []bool{true, false, true}

	for index, bits := range testCases {
		if res := isOneBitCharacter(bits); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
