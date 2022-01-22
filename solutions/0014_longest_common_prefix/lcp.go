/*
14. Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
// time: 2018-12-31

package lcp

// time complexity: O(len(strs) * max len of string)
// space complexity: O(1)
func longestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}
	res := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return res
			}
		}
		res += string(c)
	}
	return res
}
