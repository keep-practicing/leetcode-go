/*
209. Minimum Size Subarray Sum
https://leetcode.com/problems/minimum-size-subarray-sum/

Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.

Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/

package minimumsizesubarraysum

func minSubArrayLen(s int, nums []int) int {
	var (
		l   = 0
		r   = -1
		res = len(nums) + 1
		sum = 0
	)

	for l < len(nums) {
		if sum < s && r+1 < len(nums) {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s && res > r-l+1 {
			res = r - l + 1
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
