/*
376. Wiggle Subsequence
https://leetcode.com/problems/wiggle-subsequence/

A sequence of numbers is called a wiggle sequence if the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with fewer than two elements is trivially a wiggle sequence.

For example, [1,7,4,9,2,5] is a wiggle sequence because the differences (6,-3,5,-7,3) are alternately positive and negative. In contrast, [1,4,7,2,5] and [1,7,4,5,5] are not wiggle sequences, the first because its first two differences are positive and the second because its last difference is zero.

Given a sequence of integers, return the length of the longest subsequence that is a wiggle sequence. A subsequence is obtained by deleting some number of elements (eventually, also zero) from the original sequence, leaving the remaining elements in their original order.

Example 1:

Input: [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence.
Example 2:

Input: [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length. One is [1,17,10,13,10,16,8].
Example 3:

Input: [1,2,3,4,5,6,7,8,9]
Output: 2

Follow up:
Can you do it in O(n) time?
*/

package wigglesubsequence

import "algorithms/utils"

/*
 * 用up[i]和down[i]分别记录到第i个元素为止以上升沿和下降沿结束的最长“摆动”
 * 序列长度，遍历数组，如果nums[i]>nums[i-1]，表明第i-1到第i个元素是上升的，
 * 因此up[i]只需在down[i-1]的基础上加1即可，而down[i]保持down[i-1]不变；
 * 如果nums[i]<nums[i-1]，表明第i-1到第i个元素是下降的，因此down[i]
 * 只需在up[i-1]的基础上加1即可，而up[i]保持up[i-1]不变；如果nums[i]==nums[i-1]，
 * 则up[i]保持up[i-1]，down[i]保持down[i-1]。比较最终以上升沿和下降沿结束的
 * 最长“摆动”序列长度即可获取最终结果
 * */
// dynamic programming
// Time complexity: O(n^2)
// Space complexity: O(2n) ==> O(n)
func wiggleMaxLength(nums []int) int {
	n := len(nums)

	if n == 1 || n == 0 {
		return n
	}

	up := make([]int, n)
	for i := 0; i < n; i++ {
		up[i] = 1
	}
	down := make([]int, n)
	copy(down, up)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				up[i] = utils.CalcMaxInt(up[i], down[j]+1)
			} else if nums[i] < nums[j] {
				down[i] = utils.CalcMaxInt(down[i], up[j]+1)
			}
		}
	}
	return utils.CalcMaxInt(up[n-1], down[n-1])
}
