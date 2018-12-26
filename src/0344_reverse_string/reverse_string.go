/*
344. Reverse String
https://leetcode.com/problems/reverse-string/

Write a function that takes a string as input and returns the string reversed.
*/
// time: 2018-12-26

package reversestring

// double index
// time complexity: O(n)
// space complexity: O(1)
func reverseString(s string) string {
	var (
		l int
		r = len(s) - 1
	)

	for r >= l {
		charL := s[l]
		charR := s[r]
		s = s[:r] + string(charL) + s[r+1:]
		s = s[:l] + string(charR) + s[l+1:]
		r--
		l++
	}
	return s
}

// double index
// time complexity: O(n)
// space complexity: O(n)
func reverseString1(s string) string {
	sLen := len(s)
	runes := []rune(s)
	for i := 0; i < sLen/2; i++ {
		runes[i], runes[sLen-i-1] = runes[sLen-i-1], runes[i]
	}
	return string(runes)
}
