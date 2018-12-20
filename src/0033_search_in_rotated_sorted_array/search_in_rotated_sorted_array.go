/*
33. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
*/
// time: 2018-12-20

package searchinrotatedsortedarray

// binary search
/*
问题分析：
二分情况分为三种：
1.mid在pivot左边，即mid左边部分是有序的。
	nums[mid] >= nums[l] && nums[mid] > nums[r]
	此时，只需要比较target是否在左边有序部分。
2.mid在pivot右边，并且包含pivot，即mid右边是有序的。
	nums[mid] < nums[l] && nums[mid] <= nums[r],
	此时，只需要比较target是否在右边有序部分。
3. 整个array是有序的，
	nums[mid] >= nums[l] &&nums[mid] <= nums[r]
	放在第一和第二中情况中处理都可以。
*/
// time complexity: O(logn)
// space complexity: O(1)
func search(nums []int, target int) int {
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[l] && nums[mid] > nums[r] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
