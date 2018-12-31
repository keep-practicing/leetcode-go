package lcp

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{},
		{""},
	}

	expected := []string{
		"fl",
		"",
		"",
		"",
	}

	for index, strs := range testCases {
		if res := longestCommonPrefix(strs); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
