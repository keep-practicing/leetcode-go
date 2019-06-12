/*
11. Container With Most Water
https://leetcode.com/problems/container-with-most-water/

Given n non-negative integers a1, a2, ..., an ,
where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/
// time: 2018-12-26

package containerwithmostwater

import "github.com/zwfang/leetcode/utils"

// time complexity: O(n)
// space complexity: O(1)
func maxArea(height []int) int {
	var (
		water int
		l     int
		r     = len(height) - 1
	)

	for l < r {
		h := utils.CalcMinInt(height[l], height[r])
		water = utils.CalcMaxInt(water, (r-l)*h)
		for height[l] <= h && l < r {
			l++
		}
		for height[r] <= h && l < r {
			r--
		}
	}
	return water
}

// time complexity: O(n^2)
// space complexity: O(n)
func maxArea1(height []int) int {
	n := len(height)
	if n >= 0 && n < 2 {
		return 0
	}

	memo := make([]int, n)
	memo[1] = 1 * utils.CalcMinInt(height[0], height[1])
	for i := 2; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			memo[i] = utils.CalcMaxInt(memo[i], (i-j)*utils.CalcMinInt(height[i], height[j]))
		}
	}
	return utils.CalcMaxInt(memo...)
}

// time complexity: O(n^2)
// space complexity: O(1)
func maxArea2(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := (j - i) * utils.CalcMinInt(height[i], height[j])
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}
