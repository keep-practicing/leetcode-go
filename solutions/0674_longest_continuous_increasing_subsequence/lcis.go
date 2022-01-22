/*
674. Longest Continuous Increasing Subsequence
https://leetcode.com/problems/longest-continuous-increasing-subsequence/

Given an unsorted array of integers,
find the length of longest continuous increasing subsequence (subarray).
*/
// time: 2018-12-29

package lcis

// time complexity: O(n), where n is nums's length.
// space complexity: O(1)
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if 0 == n || 1 == n {
		return n
	}

	var res, l, r int

	for r < n {
		if res < r-l+1 {
			res = r - l + 1
		}
		if r+1 < n && nums[r+1] <= nums[r] {
			l = r + 1
		}
		r++
	}
	return res
}

// time complexity: O(n), where n is nums's length.
// space complexity: O(1)
func findLengthOfLCIS1(nums []int) int {
	n := len(nums)
	if 0 == n || 1 == n {
		return n
	}

	var (
		maxLen int
		count  = 1
	)

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count > maxLen {
			maxLen = count
		}
	}
	return maxLen
}
