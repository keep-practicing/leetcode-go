/*
343. Integer Break
https://leetcode.com/problems/integer-break

Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

Example 1:

Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
Example 2:

Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
Note: You may assume that n is not less than 2 and not larger than 58.
*/

package integerbreak

import "github.com/zwfang/leetcode/utils"

// recursion
/*
func integerBreak(n int) int {
	if n == 1 {
		return 1
	}

	res := utils.MinInt

	for i := 1; i < n; i++ {
		res, _ = utils.CalcMaxInt(res, i*(n-i), i*integerBreak(n-i))
	}
	return res
}
*/

// memory search
/*
func integerBreak(n int) int {
	memo := make([]int, n+1)

	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	return breakInt(n, memo)
}

func breakInt(n int, memo []int) int {
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	res := utils.MinInt
	for i := 1; i < n; i++ {
		res, _ = utils.CalcMaxInt(res, i*(n-i), i*breakInt(n-i, memo))
	}
	memo[n] = res
	return res
}
*/

// dynamic programming
// Time complexity: O(n^2)
// Space complexity: O(n+1)
func integerBreak(n int) int {
	// memo[i]表示将数字i分割（至少分割成两部分）后得到的最大乘积.
	memo := make([]int, n+1)

	memo[1] = 1

	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			memo[i] = utils.CalcMaxInt(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
}
