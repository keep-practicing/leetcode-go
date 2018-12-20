/*
27. Remove Element
https://leetcode.com/problems/remove-element/

Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.
*/
// time: 2018-12-20

package removeelement

// two pointers
// time complexity: O(n)
// space complexity: O(1)
func removeElement(nums []int, val int) int {
	x := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			if x != j {
				nums[x] = nums[j]
			}
			x++
		}
	}
	return x
}

/*
func removeElement1(nums []int, val int) int {
	var (
		l int
		r = len(nums)
	)
	for l < r {
		if nums[l] == val {
			r--
			nums[l] = nums[r]
		} else {
			l++
		}
	}
	return r
}
*/
