/*
303. Range Sum Query - Immutable
https://leetcode.com/problems/range-sum-query-immutable/

Given an integer array nums,
find the sum of the elements between indices i and j (i ≤ j), inclusive.
*/
// time: 2018-12-28

package rsqim

// NumArray store sum of nums sub-list.
type NumArray struct {
	sum []int
}

// Constructor make NumArray instance.
func Constructor(nums []int) NumArray {
	nA := NumArray{sum: make([]int, 1)}
	for _, num := range nums {
		nA.sum = append(nA.sum, nA.sum[len(nA.sum)-1]+num)
	}
	return nA
}

// SumRange find the sum of the elements between indices i and j(i <= j), inclusive.
func (na *NumArray) SumRange(i int, j int) int {
	j++
	return na.sum[j] - na.sum[i]
}
