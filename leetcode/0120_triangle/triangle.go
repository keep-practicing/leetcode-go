/*
120. Triangle
https://leetcode.com/problems/triangle/

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

*/

package triangle

import "leetcode/utils"

// dfs
/*
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	result := utils.MaxInt
	dfs(0, 0, 0, triangle, &result)
	return result
}

func dfs(x int, y int, sum int, triangle [][]int, result *int) {
	if len(triangle) == x {
		if sum < *result {
			*result = sum
		}
		return
	}
	sum += triangle[x][y]

	dfs(x+1, y, sum, triangle, result)
	dfs(x+1, y+1, sum, triangle, result)
}
*/

// dynamic programming
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}

	dp := make([][]int, m)

	dp[0] = append(dp[0], triangle[0][0])

	for i := 1; i < m; i++ {
		for j, num := range triangle[i] {
			if j >= len(dp[i-1]) {
				dp[i] = append(dp[i], num+dp[i-1][j-1])
			} else {
				if j-1 >= 0 && dp[i-1][j-1] < dp[i-1][j] {
					dp[i] = append(dp[i], num+dp[i-1][j-1])
				} else {
					dp[i] = append(dp[i], num+dp[i-1][j])
				}
			}
		}
	}

	var mininum = utils.MaxInt
	for _, num := range dp[m-1] {
		if num < mininum {
			mininum = num
		}
	}
	return mininum
}
