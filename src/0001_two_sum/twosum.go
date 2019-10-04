/*
1. Two Sum

Source: https://leetcode.com/problems/two-sum/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

package twosum

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for i, j := range nums {
		complement := target - j
		if res, ok := record[complement]; ok {
			return []int{res, i}
		}
		record[j] = i
	}
	return []int{}
}

// brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func twoSum1(nums []int, target int) []int {
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
