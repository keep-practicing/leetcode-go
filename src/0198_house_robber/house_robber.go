/*
198. House Robber
https://leetcode.com/problems/house-robber/

You are a professional robber planning to rob houses along a street. Each house has a certain amount
 of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses
  have security system connected and it will automatically contact the police if two adjacent houses were
  broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the
maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

package houserobber

import "github.com/zwfang/leetcode/utils"

/*
func rob(nums []int) int {
	return tryRob(nums, 0)
}

func tryRob(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}

	var res int
	for i := index; i < len(nums); i++ {
		res, _ = utils.CalcMaxInt(res, nums[i]+tryRob(nums, i+2))
	}
	return res
}
*/

// memory search
/*
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	return tryRob(nums, 0, memo)
}

func tryRob(nums []int, index int, memo []int) int {
	if index >= len(nums) {
		return 0
	}

	if memo[index] != -1 {
		return memo[index]
	}

	var res int
	for i := index; i < len(nums); i++ {
		res, _ = utils.CalcMaxInt(res, nums[i]+tryRob(nums, i+2, memo))
	}
	memo[index] = res
	return res
}*/

// dynamic programming
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)

	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	memo[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			if j+2 < n {
				memo[i] = utils.CalcMaxInt(memo[i], nums[j]+memo[j+2])
			} else {
				memo[i] = utils.CalcMaxInt(memo[i], nums[j])
			}
		}
	}
	return memo[0]
}
