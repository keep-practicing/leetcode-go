/*
350. Intersection of Two Arrays II
https://leetcode.com/problems/intersection-of-two-arrays-ii/

*/

package intersectionof2arrays2

func intersect(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := record[num]; !ok {
			record[num] = 1
		} else {
			record[num]++
		}
	}

	for _, num := range nums2 {
		if count, ok := record[num]; ok && count > 0 {
			res = append(res, num)
			record[num]--
		}
	}
	return res
}
