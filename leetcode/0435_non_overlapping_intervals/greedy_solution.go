/*
435. Non-overlapping Intervals
https://leetcode.com/problems/non-overlapping-intervals/

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note:
	1. You may assume the interval's end point is always bigger than its start point.
	2. Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.
*/

package nonoverlappingintervals

import "sort"

func eraseOverlapIntervalsGreedy(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Sort(IntervalSliceCompareWithEnd(intervals))
	var (
		res = 1
		pre = 0
	)

	for i := 1; i < n; i++ {
		if intervals[i].Start >= intervals[pre].End {
			res++
			pre = i
		}
	}
	return n - res
}
