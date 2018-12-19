/*
454. 4Sum II
https://leetcode.com/problems/4sum-ii/

Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.
*/

package foursum2

// Time complexity: O(n^2)
// Space complexity: O(n^2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	var (
		record = make(map[int]int)
		res    int
	)

	for _, i := range A {
		for _, j := range B {
			record[i+j]++
		}
	}

	for _, i := range C {
		for _, j := range D {
			if s, ok := record[-i-j]; ok && s > 0 {
				res += s
			}
		}
	}
	return res
}
