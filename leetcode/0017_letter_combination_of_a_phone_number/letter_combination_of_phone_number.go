/*
17. Letter Combinations of a Phone Number

Source: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package lettercombinationsofaphonenumber

var (
	letterMap = []string{
		" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	res []string
)

// Time complexity: O(2^n)
// Space complexity: O(n)
func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}

	findCombinations(digits, 0, "")
	return res
}

func findCombinations(digits string, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}

	ch := digits[index]
	letters := letterMap[ch-'0']
	for _, i := range letters {
		findCombinations(digits, index+1, s+string(i))
	}
	return
}
