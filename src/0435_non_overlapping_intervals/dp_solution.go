/*
435. Non-overlapping Intervals
https://leetcode.com/problems/non-overlapping-intervals/

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note:
	1. You may assume the interval's end point is always bigger than its start point.
	2. Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.
*/

package nonoverlappingintervals

import (
	"leetcode/utils"
	"sort"
)

func eraseOverlapIntervalsDp(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Sort(IntervalSliceCompareWithStart(intervals))

	// memo[i]表示使用intervals[0...i]的区间能构成的最长不重叠区间序列
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}

	for i := 1; i < n; i++ {
		// memo[i]
		for j := 0; j < i; j++ {
			if intervals[i].Start >= intervals[j].End {
				memo[i] = utils.CalcMaxInt(memo[i], 1+memo[j])
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		res = utils.CalcMaxInt(res, memo[i])
	}
	return n - res
}
