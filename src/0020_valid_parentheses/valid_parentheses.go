/*
20. Valid Parentheses

Source: https://leetcode.com/problems/valid-parentheses/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

	1. Open brackets must be closed by the same type of brackets.
	2. Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

package validparentheses

// Time complexity: O(n)
// Space complexity: O(n)
func isValid(s string) bool {
	var (
		stacks  []rune
		mapping = map[rune]rune{')': '(', ']': '[', '}': '{'}
	)
	for _, char := range s {
		if char == ')' || char == '}' || char == ']' {
			var topElement rune
			if len(stacks) > 0 {
				topElement = stacks[len(stacks)-1]
				stacks = stacks[:len(stacks)-1]
			} else {
				topElement = '#'
			}
			if mapping[char] != topElement {
				return false
			}
		} else {
			stacks = append(stacks, char)
		}
	}
	return len(stacks) == 0
}
