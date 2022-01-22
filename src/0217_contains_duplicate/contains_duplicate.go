/*
217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array,
and it should return false if every element is distinct.
*/
// time: 2019-01-07

package cd

// using HashTable
// time complexity: O(n)
// space complexity: O(n)
func containsDuplicate(nums []int) bool {
	record := make(map[int]int)

	for _, num := range nums {
		if _, ok := record[num]; !ok {
			record[num] = 1
		} else {
			return true
		}
	}
	return false
}
