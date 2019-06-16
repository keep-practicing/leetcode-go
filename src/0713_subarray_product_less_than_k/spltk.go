/*
713. Subarray Product Less Than K
https://leetcode.com/problems/subarray-product-less-than-k/

Your are given an array of positive integers nums.

Count and print the number of (contiguous) subarrays
where the product of all the elements in the subarray is less than k.
*/

// time: 2018-12-27

package spltk

// sliding window
// time complexity: O(n), where n = len(nums)
// space complexity: O(1)
func numSubArrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	var (
		n    = len(nums)
		l, r int
		res  int
		prod = 1
	)

	for l < n {
		if r < n && prod*nums[r] < k {
			prod *= nums[r]
			r++
		} else if l == r {
			l++
			r++
		} else {
			res += r - l
			prod /= nums[l]
			l++
		}
	}
	return res
}

// 2019-06-16
func numSubArrayProductLessThanK2(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	var (
		prod = 1
		res  = 0
		left = 0
	)

	for right, val := range nums {
		prod *= val
		for prod >= k {
			prod /= nums[left]
			left++
		}
		res += right - left + 1
	}
	return res
}
