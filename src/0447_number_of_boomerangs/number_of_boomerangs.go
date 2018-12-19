/*
447. Number of Boomerangs
https://leetcode.com/problems/number-of-boomerangs/

Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).
*/

package numberofboomerangs

import (
	"math"
)

// Time complexity: O(n^2)
// Space complexity: O(n)
func numberOfBoomerangs(points [][]int) int {
	var (
		res int
		n   = len(points)
	)

	for i := 0; i < n; i++ {
		record := make(map[int]int)
		for j := 0; j < n; j++ {
			if j != i {
				dis := dis(points[i], points[j])
				record[dis]++
			}
		}

		for _, j := range record {
			res += j * (j - 1)
		}
	}
	return res
}

func dis(point1 []int, point2 []int) int {
	return int(math.Pow(float64(point1[0]-point2[0]), float64(2)) + math.Pow(float64(point1[1]-point2[1]), float64(2)))
}
