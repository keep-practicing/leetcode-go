package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testData := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	expectedData := []int{3, 1, 3}

	for index, data := range testData {
		if res := lengthOfLongestSubstring(data); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
