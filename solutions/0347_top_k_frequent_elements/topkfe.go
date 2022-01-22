/*
347. Top K Frequent Elements
https://leetcode.com/problems/top-k-frequent-elements/

Given a non-empty array of integers, return the k most frequent elements.

Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
// time: 2019-01-08

package topkfe

import "container/heap"

// using priority
// time complexity: O(n log n)
// space complexity: O(n)
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		if _, ok := count[num]; !ok {
			count[num] = 1
		} else {
			count[num]++
		}
	}

	nums_ := make(Nums, 0)

	for num, cnt := range count {
		nums_ = append(nums_, Num{Val: num, Count: cnt})
	}
	heap.Init(&nums_)
	var res []int
	for i := 0; i < k; i++ {
		num := heap.Pop(&nums_).(Num)
		res = append(res, num.Val)
	}
	return res
}

// Num stores its value and frequency as Count.
type Num struct {
	Val   int
	Count int
}

// Nums struct for impl Interface
type Nums []Num

// Len sort Interface
func (n Nums) Len() int { return len(n) }

// Swap sort Interface
func (n Nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// Less sort Interface
func (n Nums) Less(i, j int) bool { return n[i].Count >= n[j].Count }

// Push heap Interface
func (n *Nums) Push(num interface{}) {
	m := num.(Num)
	*n = append(*n, m)
}

// Pop heap Interface
func (n *Nums) Pop() interface{} {
	res := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return res
}
