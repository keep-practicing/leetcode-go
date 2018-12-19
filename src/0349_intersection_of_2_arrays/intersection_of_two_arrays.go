/*
349. Intersection of Two Arrays
https://leetcode.com/problems/intersection-of-two-arrays

Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
	Each element in the result must be unique.
	The result can be in any order.
*/

package intersectionof2arrays

import "leetcode/utils"

func intersection(nums1 []int, nums2 []int) []int {
	set1 := utils.NewSet()
	for _, num := range nums1 {
		set1.Add(num)
	}
	set2 := utils.NewSet()
	for _, num := range nums2 {
		set2.Add(num)
	}

	var res []int
	for item := range set1 {
		if set2.Contains(item) {
			if value, ok := item.(int); ok {
				res = append(res, value)
			}
		}
	}
	return res
}
