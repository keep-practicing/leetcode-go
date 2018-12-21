/*
75. Sort Colors
https://leetcode.com/problems/sort-colors/

Given an array with n objects colored red, white or blue,
sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.
*/

// time: 2018-12-21

package sortcolors

// 计数排序
// time complexity: O(n)
// space complexity: O(1)
func sortColorsCountSort(nums []int) {
	var count = [3]int{}

	for _, num := range nums {
		count[num]++
	}

	index := 0

	for i, j := range count {
		for k := 0; k < j; k++ {
			nums[index] = i
			index++
		}
	}
}

// quick sort 3 ways
// time complexity: O(n)
// space complexity: O(1)
func sortColorsQuickSort3Ways(nums []int) {
	var (
		zero = -1
		two  = len(nums)
	)

	for i := 0; i < two; {
		if 1 == nums[i] {
			i++
		} else if 2 == nums[i] {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else { // nums[i] == 0
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}
