/*
136. Single Number
https://leetcode.com/problems/single-number/

Given a non-empty array of integers, every element appears twice except for one. Find that single one.
*/
// time: 2019-02-01

package sn

// time complexity: O(n)
// space complexity: O(n)
func singleNumber(nums []int) int {
	record := make(map[int]int)
	for _, num := range nums {
		if _, ok := record[num]; ok {
			delete(record, num)
		} else {
			record[num] = 1
		}
	}
	var res int
	for key, _ := range record {
		res = key
	}
	return res
}

// time complexity: O(n)
// space complexity: O(1)
func singleNumber1(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
