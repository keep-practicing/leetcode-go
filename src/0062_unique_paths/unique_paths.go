/*
62. Unique Paths
Source: https://leetcode.com/problems/unique-paths/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?


Above is a 7 x 3 grid. How many possible unique paths are there?

Note: m and n will be at most 100.

Example 1:

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:

Input: m = 7, n = 3
Output: 28
*/

package uniquepaths

// recursion
/*
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m, n-1) + uniquePaths(m-1, n)
}
*/

// memory search
/*
func uniquePaths(m int, n int) int {
	memo := make([][]int, m) // memo[i][j]存储(i+1)*(j+1)grid的路径数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	return calcPaths(m, n, memo)
}

func calcPaths(m int, n int, memo [][]int) int {
	if memo[m-1][n-1] != -1 {
		return memo[m-1][n-1]
	}

	if m == 1 || n == 1 {
		memo[m-1][n-1] = 1
		return 1
	}

	res := calcPaths(m, n-1, memo) + calcPaths(m-1, n, memo)
	memo[m-1][n-1] = res
	return res
}
*/

func uniquePaths(m int, n int) int {
	memo := make([][]int, m) // memo[i][j]存储(i+1)*(j+1)grid的路径数量
	for i := 0; i < m; i++ {
		memo[i] = append(memo[i], 1)
	}

	for i := 1; i < n; i++ {
		memo[0] = append(memo[0], 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i] = append(memo[i], memo[i-1][j]+memo[i][j-1])
		}
	}
	return memo[m-1][n-1]
}
