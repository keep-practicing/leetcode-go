/*
66. Plus One
https://leetcode.com/problems/plus-one/

Given a non-empty array of digits representing a non-negative integer,
plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
*/
// time: 2019-01-02

package plusone

// time complexity: O(n)
// space complexity: O(1)
func plusOne(digits []int) []int {
	var (
		carry = 1
		index = len(digits) - 1
	)
	for index >= 0 {
		sum := carry + digits[index]
		digits[index] = sum % 10
		carry = sum / 10
		index--
	}

	if carry > 0 {
		digits = append(digits, 0)
		copy(digits[1:], digits[0:])
		digits[0] = carry
	}
	return digits
}
