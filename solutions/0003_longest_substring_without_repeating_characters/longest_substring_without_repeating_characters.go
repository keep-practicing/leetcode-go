/*
3. Longest Substring Without Repeating Characters

Source: https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package longestsubstringwithoutrepeatingcharacters

// Time complexity: O(n)
// Space complexity:  O(n)
func lengthOfLongestSubstring(s string) int {

	var (
		start              = 0
		res                = 0
		lastOccurredRecord = make(map[rune]int)
	)

	for i, ch := range []rune(s) {
		if lastIndex, ok := lastOccurredRecord[ch]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		if i-start+1 > res {
			res = i - start + 1
		}
		lastOccurredRecord[ch] = i
	}
	return res
}
