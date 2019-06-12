/*
4. Median of Two Sorted Arrays
https://leetcode.com/problems/median-of-two-sorted-arrays/

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.
*/
// time: 2018-12-29

package motsa

import "github.com/zwfang/leetcode/utils"

// binary search
// time complexity: O(log(m+n)), where m is nums1's length, n is nums2's length.
// space complexity: O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		m = len(nums1)
		n = len(nums2)
	)

	if 0 == m && 0 == n {
		panic("nums1 and nums2 cannot be both empty.")
	}

	if 0 == m {
		return float64(nums2[n>>1]+nums2[n>>1-(n+1)%2]) / 2
	}

	if 0 == n {
		return float64(nums1[m>>1]+nums1[m>>1-(m+1)%2]) / 2
	}

	var (
		left  = (m + n + 1) >> 1
		right = (m + n + 2) >> 1
	)
	return float64(findKth(nums1, 0, nums2, 0, left)+findKth(nums1, 0, nums2, 0, right)) / 2
}

/*
使用二分法, 在两个有序数组中找到第k个元素。
使用两个变量i和j分别来标记数组nums1和nums2的起始位置。
1. 	当某一个数组的起始位置大于等于其数组长度时，说明其所有数字均已经被淘汰了，
	相当于一个空数组了，那么实际上就变成了在另一个数组中找数字，直接就可以找出来了
2. 	如果K=1的话，那么我们只要比较nums1和nums2的起始位置i和j上的数字就可以了。
3.	对K二分，分别在nums1和nums2中查找第K/2个元素，注意这里由于两个数组的长度不定，所以有可能某个数组没有第K/2个数字，
	需要先检查一下，数组中到底存不存在第k/2个数字，如果存在就取出来，反之赋值为整型最大值，
	比较两个数组的第k/2个数字，如果第一个数组的第k/2个数字小的话，说明我们要找的数字肯定不再第一个数组的前k/2个数字中，
	所以舍弃掉，将第一个数组的起始位置移动k/2，并且此时k也减去k/2（舍弃了k/2个之后，只需要在剩下的里面寻找k-k/2就可以了）
*/
func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if 1 == k {
		if nums1[i] < nums2[j] {
			return nums1[i]
		}
		return nums2[j]
	}
	var (
		midVal1, midVal2 int
	)

	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = utils.MaxInt
	}

	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = utils.MaxInt
	}

	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	}
	return findKth(nums1, i, nums2, j+k/2, k-k/2)
}
