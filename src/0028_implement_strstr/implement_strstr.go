/*
28. Implement strStr()
https://leetcode.com/problems/implement-strstr/

Implement strStr().

Return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.

clarification:
For the purpose of this problem, we will return 0 when needle is an empty string.
*/
// time: 2019-01-02

package implstr

// time complexity: O(	(m-n+1)*n	), where m is len(haystack), n is len(needle)
// space complexity: O(1)
func strStr(haystack string, needle string) int {
	if 0 == len(needle) {
		return 0
	}

	for i, j := 0, 0; i <= len(haystack)-len(needle); i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}
