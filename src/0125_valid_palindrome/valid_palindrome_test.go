package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"0P",
		"",
	}

	expected := []bool{true, false, false, true}

	for index, data := range testCases {
		if res := isPalindrome(data); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
