/*
455. Assign Cookies
https://leetcode.com/problems/assign-cookies/

Assume you are an awesome parent and want to give your children some cookies.
But, you should give each child at most one cookie.
Each child i has a greed factor gi, which is the minimum size of a cookie that the child will be content with;
and each cookie j has a size sj. If sj >= gi, we can assign the cookie j to the child i, and the child i will be content.
Your goal is to maximize the number of your content children and output the maximum number.

Note:
	1. You may assume the greed factor is always positive.
	2. You cannot assign more than one cookie to one child.
*/

package assigncookies

import "sort"

// greedy
// Time complexity: O(n), where n is min (  len(g), len(s)       )
// Space complexity: O(1)
func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(g)))

	var (
		si, gi, res int
	)

	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			gi++
			si++
		} else {
			gi++
		}
	}

	return res
}
