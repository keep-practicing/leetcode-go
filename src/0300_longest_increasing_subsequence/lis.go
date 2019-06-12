/*
300. Longest Increasing Subsequence
https://leetcode.com/problems/longest-increasing-subsequence/

Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.

Follow up:
	Could you improve it to O(n log n) time complexity?
*/

package lis

import "github.com/zwfang/leetcode/utils"

// Dynamic Programming
// TIme complexity: O(n^2)
// Space Complexity: O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				memo[i] = utils.CalcMaxInt(memo[i], 1+memo[j])
			}
		}
	}
	return utils.CalcMaxInt(memo...)
}
