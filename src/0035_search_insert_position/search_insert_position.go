/*
35. Search Insert Position
https://leetcode.com/problems/search-insert-position/

Given a sorted array and a target value,
return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.
*/
// time: 2019-01-02

package sip

// binary search
// time complexity: O(	log n		)
// space complexity: O(1)
func searchInsert(nums []int, target int) int {
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
