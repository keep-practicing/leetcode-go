/*
15. 3Sum
https://leetcode.com/problems/3sum/
*/
// time: 2019-03-01

package threesum

import "sort"

// time complexity: O(n^2)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l, r := i+1, len(nums)-1; l < r; {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			switch sum := nums[i] + nums[l] + nums[r]; {
			case sum < 0:
				l++
			case sum > 0:
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
		}
	}
	return res
}
