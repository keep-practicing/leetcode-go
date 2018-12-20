/*
26. Remove Duplicates from Sorted Array
https://leetcode.com/problems/remove-duplicates-from-sorted-array/

Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
*/
// time: 2018-12-20

package rdfsa

// double index
// time complexity: O(n)
// space complexity: O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)

	if 0 == n {
		return n
	}

	var (
		res   = 1
		i     = 1
		index = nextDifferentCharacterIndex(nums, 1) // 下一个不重复的地址
	)

	for index < n {
		res++
		nums[i] = nums[index]
		i++
		index = nextDifferentCharacterIndex(nums, index+1)
	}
	return res
}

func nextDifferentCharacterIndex(nums []int, p int) int {
	for ; p < len(nums); p++ {
		if nums[p] != nums[p-1] {
			break
		}
	}
	return p
}
