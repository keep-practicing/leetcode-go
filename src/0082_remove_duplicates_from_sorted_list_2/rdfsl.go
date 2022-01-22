/*
82. Remove Duplicates from Sorted List II
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

Given a sorted linked list, delete all nodes that have duplicate numbers,
leaving only distinct numbers from the original list.
*/
// time: 2019-01-04

package rdfsl

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	pre := dummyHead
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			num := cur.Val
			for cur != nil && cur.Val == num {
				cur = cur.Next
			}
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
