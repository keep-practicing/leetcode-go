/*
58. Length of Last Word
https://leetcode.com/problems/length-of-last-word/

Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word in the string.
If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.
*/
// time: 2019-01-02

package lenoflastword

// time complexity: O(n)
// space complexity: O(1)
func lengthOfLastWord(s string) int {
	var (
		n   = len(s)
		cur = n - 1
	)

	for cur >= 0 {
		if n-1 == cur && s[cur] == 32 {
			cur--
			n--
			continue
		}
		if s[cur] != 32 {
			cur--
		} else {
			break
		}
	}
	return n - cur - 1
}
