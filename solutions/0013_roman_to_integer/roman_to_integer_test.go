package rti

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []string{
		"III",
		"IV",
		"IX",
		"LVIII",
		"MCMXCIV",
		"",
	}
	expected := []int{3, 4, 9, 58, 1994, 0}

	for index, s := range testCases {
		if res := romanToInt(s); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
