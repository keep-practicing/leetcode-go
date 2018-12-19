/*
Source: https://leetcode.com/problems/find-all-anagrams-in-a-string/submissions/

Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

Time: 2018-12-17
*/

package allanagramsinastring

// Sliding Window
// Time complexity: O(len(p) + len(s)) = O(n)
// Space complexity: O(1)
func findAnagrams(s string, p string) []int {
	var (
		lenS  = len(s)
		lenP  = len(p)
		res   = []int{}
		freqP = make([]int, 26)
		freqS = make([]int, 26)
		left  int
		right = -1 // Sliding window: s[left, right]
	)

	if lenS < lenP {
		return res
	}

	for i := 0; i < lenP; i++ {
		freqP[p[i]-'a']++
	}

	for right+1 < lenS {
		right++
		freqS[s[right]-'a']++
		if right-left+1 > lenP {
			freqS[s[left]-'a']--
			left++
		}
		if right-left+1 == lenP && isSame(freqP, freqS) {
			res = append(res, left)
		}
	}
	return res
}

func isSame(p []int, q []int) bool {
	for i := 0; i < 26; i++ {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}
