/*
63. Unique Paths II

Source: https://leetcode.com/problems/unique-paths-ii/

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?



An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/

package uniquepaths2

// recursion
/*
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	cur := 1
	if m == 1 {
		for i := 0; i < n; i++ {
			if obstacleGrid[0][i] == 1 {
				cur = 0
			}
		}
		return cur
	}

	cur = 1
	if n == 1 {
		for i := 0; i < m; i++ {
			if obstacleGrid[i][0] == 1 {
				cur = 0
			}
		}
		return cur
	}
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	obstacleGridCut := make([][]int, m)

	for i := range obstacleGridCut {
		for j := 0; j < n-1; j++ {
			obstacleGridCut[i] = append(obstacleGridCut[i], obstacleGrid[i][j])
		}
	}

	return uniquePathsWithObstacles(obstacleGrid[:m-1]) + uniquePathsWithObstacles(obstacleGridCut)
}
*/

// memory search
/*
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i] = append(memo[i], -1)
		}
	}

	cur := 1
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			cur = 0
		}
		memo[i][0] = cur
	}

	cur = 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			cur = 0
		}
		memo[0][j] = cur
	}
	return calcUniquePaths(obstacleGrid, m-1, n-1, memo)
}

func calcUniquePaths(obstacleGrid [][]int, m int, n int, memo [][]int) int {
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	if obstacleGrid[m][n] == 1 {
		return 0
	}
	res := calcUniquePaths(obstacleGrid, m-1, n, memo) + calcUniquePaths(obstacleGrid, m, n-1, memo)
	memo[m][n] = res
	return res
}
*/

// dynamic programming
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	cur := 1
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			cur = 0
		}
		memo[i][0] = cur
	}

	cur = 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			cur = 0
		}
		memo[0][j] = cur
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				memo[i][j] = 0
			} else {
				memo[i][j] = memo[i-1][j] + memo[i][j-1]
			}
		}
	}
	return memo[m-1][n-1]
}
