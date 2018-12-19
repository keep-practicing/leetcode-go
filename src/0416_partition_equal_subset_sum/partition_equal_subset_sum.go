/*
416. Partition Equal Subset Sum
https://leetcode.com/problems/partition-equal-subset-sum/

Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:
	1. Each of the array element will not exceed 100.
	2. The array size will not exceed 200.
*/

package partitionequalsubsetsum

/*
此问题可以使用0-1背包问题的思路求解

c是数组和的一半。

状态定义：F(n c)

状态转移方程：F(n c) = max(    F(n-1, c) ,     n + F(n-1, c-n)          )
*/
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	c := sum / 2
	n := len(nums)
	memo := make([]bool, c+1)

	for i := 0; i <= c; i++ {
		memo[i] = nums[0] == i
	}

	for i := 0; i < n; i++ {
		for j := c; j >= nums[i]; j-- {
			memo[j] = memo[j] || memo[j-nums[i]]
		}
	}
	return memo[c]
}
