/*
83. Remove Duplicates from Sorted List
https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given a sorted linked list, delete all duplicates such that each element appear only once.
*/

package rdfsl

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		pre = head
		cur = head.Next
	)

	for cur != nil {
		if cur.Val == pre.Val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head
}
