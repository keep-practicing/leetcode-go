/*
125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.
*/

// time: 2018-12-26

package validpalindrome

// double index
// time complexity: O(n)
// space complexity: O(n)
func isPalindrome(s string) bool {
	chars := make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			chars = append(chars, s[i])
		}
		if s[i] >= 97 && s[i] <= 122 {
			chars = append(chars, s[i]-32)
		}
		if s[i] >= 48 && s[i] <= 57 {
			chars = append(chars, s[i])
		}
	}

	var (
		l int
		r = len(chars) - 1
	)

	for r >= l {
		if chars[r] != chars[l] {
			return false
		}
		r--
		l++
	}
	return true
}
