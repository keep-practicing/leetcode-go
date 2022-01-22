/*
77. Combinations
https://leetcode.com/problems/combinations/

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
*/
// time: 2019-01-04

package combinations

// backtracking
// time complexity: O(n^k)
// space complexity: O(k)
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	res := make([][]int, 0)
	c := make([]int, 0, k)
	generateCombinations(n, k, 1, c, &res)
	return res
}

// 求解C(n,k), 当前已经找到的组合存储在c中，需要从start开始搜索新元素。
func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		cpy := make([]int, k)
		copy(cpy, c)
		*res = append(*res, cpy)
		return
	}
	// 回朔法剪枝。
	// 还有k - len(c)个空位，所以， [i...n]中至少要有k - len(c)个元素
	// i最多为n-(k-len(c))+1
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}
