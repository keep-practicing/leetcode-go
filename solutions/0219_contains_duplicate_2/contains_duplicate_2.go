/*
219. Contains Duplicate II
https://leetcode.com/problems/contains-duplicate-ii/

Given an array of integers and an integer k,
find out whether there are two distinct indices i and j in the array
such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
*/
// time: 2019-01-07

package cond

// using hash table
// time complexity: O(n)
// space complexity: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	record := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := record[v]; ok {
			if i-j <= k {
				return true
			}
		}
		record[v] = i
	}
	return false
}
