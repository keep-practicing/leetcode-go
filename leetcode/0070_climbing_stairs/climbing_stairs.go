/*
70. Climbing Stairs

source:https://leetcode.com/problems/climbing-stairs/

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

package climbingstairs

// recursion
/*
func climbStairs(n int) int {
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	//

	if n == 1 || n == 0 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
*/

// memory search
/*
func climbStairs(n int) int {
	var memo []int
	for i := 0; i <= n; i++ {
		memo = append(memo, -1)
	}
	return calcWays(n, memo)
}

func calcWays(n int, memo []int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if memo[n] == -1 {
		memo[n] = calcWays(n-1, memo) + calcWays(n-2, memo)
	}

	return memo[n]
}
*/

// dynamic programming
func climbStairs(n int) int {
	var memo = []int{1, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n]
}
