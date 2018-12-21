/*
80. Remove Duplicates from Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

Given a sorted array nums, remove the duplicates in-place such that
duplicates appeared at most twice and return the new length.

Do not allocate extra space for another array, you must do this
 by modifying the input array in-place with O(1) extra memory.
*/
// time: 2018-12-21

package rdfsa2

// double index
// time complexity: O(n)
// space complexity: O(1)
func removeDuplicates(nums []int) int {
	n := len(nums)
	if 0 == n || 1 == n {
		return n
	}

	var (
		res   = 2
		i     = 2
		index = nextDifferentCharacterIndex(nums, i, 2) // 寻找下一个满足条件的下标
	)

	for index < n {
		res++
		nums[i] = nums[index]
		i++
		index = nextDifferentCharacterIndex(nums, i, index+1)
	}
	return res
}

func nextDifferentCharacterIndex(nums []int, j int, p int) int {
	for ; p < len(nums); p++ {
		if nums[j-1] == nums[j-2] && nums[p] != nums[j-1] || nums[j-1] != nums[j-2] {
			break
		}
	}
	return p
}
