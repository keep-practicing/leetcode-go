/*
704. Binary Search
https://leetcode.com/problems/binary-search/

Given a sorted (in ascending order) integer array nums of n elements and a target value,
write a function to search target in nums. If target exists, then return its index, otherwise return -1.

Note:
	1. You may assume that all elements in nums are unique.
	2. n will be in the range [1, 10000].
	3. The value of each element in nums will be in the range [-9999, 9999].
*/

// time: 2018-12-19

package binarysearch

// iterative
// Time complexity: O(log(n))
// Space complexity: O(1)
func search(nums []int, target int) int {
	var (
		l int
		r = len(nums) - 1
	) // 在[l...r]区间寻找target

	for l <= r { // 当l==r时，区间依然是有效的。
		mid := l + (r-l)/2 // (l+r)/2  可能会整型益处
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
