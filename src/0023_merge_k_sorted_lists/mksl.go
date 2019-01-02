/*
23. Merge k Sorted Lists
https://leetcode.com/problems/merge-k-sorted-lists/

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
*/
// time: 2019-01-02

package mksl

import (
	"sort"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// heap
// time complexity: O(n log n), where n is len of all lists, the main complexity is sort.
// space complexity: O(n), where n is len of all lists.
func mergeKLists(lists []*ListNode) *ListNode {
	for i := 0; i < len(lists); {
		if lists[i] == nil {
			lists = append(lists[0:i], lists[i+1:]...)
		} else {
			i++
		}
	}
	if 0 == len(lists) {
		return nil
	}

	head := &ListNode{}
	cur := head
	for len(lists) > 0 {
		for i := (len(lists) - 1) / 2; i >= 0; i-- {
			shiftUp(lists, len(lists), i)
		}
		cur.Next = &ListNode{Val: lists[0].Val}
		cur = cur.Next
		lists[0] = lists[0].Next
		if lists[0] == nil {
			lists = lists[1:]
		}
	}
	return head.Next
}

// build heap
func shiftUp(lists []*ListNode, n int, k int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && lists[j+1].Val < lists[j].Val {
			j = j + 1
		}
		if lists[k].Val < lists[j].Val {
			break
		}

		lists[j], lists[k] = lists[k], lists[j]
		k = j
	}
}

// brute force
// time complexity: O(n log n), where n is len of all lists, the main complexity is sort.
// space complexity: O(2n), where n is len of all lists.
func mergeKLists1(lists []*ListNode) *ListNode {
	nodes := make([]int, 0)
	head := &ListNode{}
	pointer := head

	for _, j := range lists {
		for {
			if j != nil {
				nodes = append(nodes, j.Val)
				j = j.Next
			} else {
				break
			}
		}
	}
	sort.Ints(nodes)
	for _, j := range nodes {
		pointer.Next = &ListNode{Val: j}
		pointer = pointer.Next
	}
	return head.Next
}
