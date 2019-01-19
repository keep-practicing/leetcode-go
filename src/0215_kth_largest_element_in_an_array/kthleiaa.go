/*
215. Kth Largest Element in an Array
https://leetcode.com/problems/kth-largest-element-in-an-array/

Find the kth largest element in an unsorted array.
 Note that it is the kth largest element in the sorted order, not the kth distinct element.

You may assume k is always valid, 1 ≤ k ≤ array's length.

*/

// time: 2018-12-22

package kthleiaa

// heap sort
// time complexity: O(k * logn)
// sapce complexity: O(1)
func findKthLargest(nums []int, k int) int {
	// heapify
	for i := (len(nums) - 1) / 2; i >= 0; i-- {
		siftDown(nums, len(nums), i)
	}

	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		siftDown(nums, i, 0)
	}
	return nums[len(nums)-k]
}

func siftDown(nums []int, n int, i int) {
	for 2*i+1 < n {
		j := 2*i + 1
		if j+1 < n && nums[j+1] > nums[j] {
			j++
		}

		if nums[i] >= nums[j] {
			break
		}
		nums[j], nums[i] = nums[i], nums[j]
		i = j
	}
}
