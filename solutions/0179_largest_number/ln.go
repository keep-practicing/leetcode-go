/*
179. Largest Number
https://leetcode.com/problems/largest-number/

Given a list of non negative integers, arrange them such that they form the largest number.
Note: The result may be very large, so you need to return a string instead of an integer.
*/
// time: 2019-01-14

package ln

import (
	"sort"
	"strconv"
	"strings"
)

type sliceString []string

// Len is the number of elements in the collection.
func (s sliceString) Len() int { return len(s) }

// Swap swaps the elements with indexes i and j.
func (s sliceString) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less reports whether the element with
// index i should sort before the element with index j.
func (s sliceString) Less(i, j int) bool { return (s[i] + s[j]) > (s[j] + s[i]) }

// time complexity: O(n log n), dominated by the complexity of sort.
// space complexity: O(n)
func largestNumber(nums []int) string {
	numsString := make([]string, 0)
	for _, num := range nums {
		numsString = append(numsString, strconv.Itoa(num))
	}

	sort.Sort(sliceString(numsString))
	numStr := strings.Join(numsString, "")

	if strings.HasPrefix(numStr, "0") {
		return "0"
	}
	return numStr
}
