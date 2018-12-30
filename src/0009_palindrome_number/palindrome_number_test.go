package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []int{121, -121, 10, 0}
	expected := []bool{true, false, false, true}
	for index, data := range testCases {
		if res := isPalindrome(data); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
