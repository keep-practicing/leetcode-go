/*
153. Find Minimum in Rotated Sorted Array
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.
*/
// time: 2019-01-07

package fmirsa

// binary search
// time complexity: O( log n)
// space complexity: O(1)
func findMin(nums []int) int {
	var (
		low  int
		high = len(nums) - 1
		mid  int
	)
	for low < high {
		mid = low + (high-low)>>1
		if nums[high] < nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
