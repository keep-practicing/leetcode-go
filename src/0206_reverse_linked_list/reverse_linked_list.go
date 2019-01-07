/*
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/

Reverse a singly linked list.
*/
// time: 2019-01-07

package rll

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// recursive
// time complexity: O(n)
// space complexity: O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}

// iterative
// time complexity: O(n)
// space complexity: O(1)
func reverseList1(head *ListNode) *ListNode {
	var (
		pre *ListNode
		cur = head
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
