/*
76. Minimum Window Substring
https://leetcode.com/problems/minimum-window-substring/

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:
	Input: S = "ADOBECODEBANC", T = "ABC"
	Output: "BANC"

Note:
	1. If there is no such window in S that covers all characters in T, return the empty string "".
	2. If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

package minimumwindowsubstring

/*
使用滑动窗口解决这一问题，使用map或者slice统计T字符串中的字母的个数，之后，开始遍历S字符串，对于S中遍历到
的每一个字母，都在map或者slice中的对应个数减一，如果减一后仍然大于等于0，说明当前遍历到的字母是T中的字母，
使用计数器count，使其加一，当count和T串字母个数相等时，说明此时的窗口已经包含了T串中的所有字母，此时记录
一下字串和字串长度，
然后开始收缩左边界，由于我们遍历的时候，对应值减了1，所以此时去除字母的时候，就要把减去的1加回来，
此时如果加1后的值大于0了，说明此时我们少了一个T中的字母，那么count值就要减1了，然后移动左边界left。
*/

import "algorithms/utils"

// sliding window
// Time complexity: O(n)
// Space complexity: O(128) = O(1)
func minWindow(s string, t string) string {
	var (
		res         string
		letterCount = make([]int, 128)
		left        int
		count       int
		minLen      = utils.MaxInt
	)

	for i := 0; i < len(t); i++ {
		letterCount[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if letterCount[s[i]]--; letterCount[s[i]] >= 0 {
			count++
		}

		for count == len(t) {
			if minLen > i-left+1 {
				minLen = i - left + 1
				res = s[left : minLen+left]
			}
			if letterCount[s[left]]++; letterCount[s[left]] > 0 {
				count--
			}
			left++
		}
	}
	return res
}
