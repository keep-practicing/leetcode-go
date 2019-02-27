/*
304. Range Sum Query 2D - Immutable
https://leetcode.com/problems/range-sum-query-2d-immutable/
*/
// 2019-02-27

package rsq

// NumMatrix 累计区域和的数组
type NumMatrix struct {
	dp [][]int
}

// Constructor 初始化构造函数
func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	numMatrix := NumMatrix{dp: make([][]int, len(matrix)+1)}
	for i := 0; i < len(numMatrix.dp); i++ {
		numMatrix.dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			numMatrix.dp[i][j] = numMatrix.dp[i-1][j] + numMatrix.dp[i][j-1] - numMatrix.dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return numMatrix
}

// SumRegion 求区域和
func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.dp[row2+1][col2+1] - nm.dp[row1][col2+1] - nm.dp[row2+1][col1] + nm.dp[row1][col1]
}
