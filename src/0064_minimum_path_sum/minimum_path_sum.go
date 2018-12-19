/*
64. Minimum Path Sum
source: https://leetcode.com/problems/minimum-path-sum/

Description
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
 Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

package minimumpathsum

// recursion dfs
/*
func minPathSum(grid [][]int) int {
    if len(grid) == 0{
        return 0
    }

    MaxUint:= ^uint(0)
    MaxInt := int(MaxUint >> 1)

    result := MaxInt

    dfs(0,0,0, grid, &result)
    return result
}

func dfs(x int, y int, sum int, grid [][]int, result *int) {
    m := len(grid)
    n := len(grid[0])

    sum += grid[x][y]

    if x+1 < m && y+1 < n {
        dfs(x+1, y, sum, grid, result)
        dfs(x, y+1, sum, grid, result)
    }

    if x +1 == m && y+1 < n {
        dfs(x, y+1, sum, grid, result)
    }

    if y+1 == n && x+1 < m {
        dfs(x+1, y, sum, grid, result)
    }


    if x+1 == m && y+1 == n {
        if sum < *result {
            *result = sum
        }
        return
    }
}*/

// dynamic programming
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	dp[0] = append(dp[0], grid[0][0])

	for i := 1; i < m; i++ {
		dp[i] = append(dp[i], grid[i][0]+dp[i-1][0])
	}

	for i := 1; i < n; i++ {
		dp[0] = append(dp[0], grid[0][i]+dp[0][i-1])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i] = append(dp[i], grid[i][j]+dp[i-1][j])
			} else {
				dp[i] = append(dp[i], grid[i][j]+dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}
