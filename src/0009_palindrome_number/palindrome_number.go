/*
9. Palindrome Number
https://leetcode.com/problems/palindrome-number/

Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.
*/
// time: 2018-12-30

package palindromenumber

// Time complexity: O(log 10 (n))
// Space complexity : O(1)
func isPalindrome(x int) bool {
	var (
		y int
		z = x
	)

	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return y == z
}
