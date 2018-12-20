/*
34. Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].
*/
// time: 2018-12-20

package findfirstandlastpositionofelementinsortedarray

// binary search
// time complexity: O(logn)
// space complexity: O(1)
func searchRange(nums []int, target int) []int {
	lowerIndex := firstOccurance(nums, target)
	first := -1
	if lowerIndex != len(nums) && target == nums[lowerIndex] {
		first = lowerIndex
	}

	upperIndex := lastOccurance(nums, target)
	last := -1
	if upperIndex == len(nums) && len(nums) > 0 && target == nums[len(nums)-1] {
		last = len(nums) - 1
	} else if upperIndex != len(nums) && upperIndex > 0 && target == nums[upperIndex-1] {
		last = upperIndex - 1
	}
	return []int{first, last}
}

func firstOccurance(nums []int, target int) int {
	var (
		l int
		r = len(nums)
	)

	for l != r { // 夹逼思想
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func lastOccurance(nums []int, target int) int {
	var (
		l int
		r = len(nums)
	)
	for l != r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// double index scan
// Time complexity: O(n)
// Space complexity: O(1)
func searchRange1(nums []int, target int) []int {
	var (
		l int
		r = len(nums) - 1
	)

	for l <= r {
		flag := false
		if nums[l] != target {
			l++
			flag = true
		}
		if nums[r] != target {
			r--
			flag = true
		}
		if !flag {
			break
		}
	}
	if r < l {
		return []int{-1, -1}
	}
	return []int{l, r}
}

// binary search + linear scan
// Time complexity: O(logn) ~ O(n)
// Space complexity: O(1)
func searchRange2(nums []int, target int) []int {
	var (
		l   int
		r   = len(nums) - 1
		tmp = -1
	)

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			tmp = mid
			break
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if -1 == tmp {
		return []int{-1, -1}
	}
	l = tmp
	r = tmp
	for true {
		if l > 0 && nums[l-1] == target {
			l--
		} else {
			break
		}
	}
	for true {
		if r < len(nums)-1 && target == nums[r+1] {
			r++
		} else {
			break
		}
	}
	return []int{l, r}
}
