/*
19. Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given a linked list, remove the n-th node from the end of list and return its head.
*/
// time: 2018-12-31

package rnthnfeol

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head

	var (
		p = dummy
		q = dummy
	)

	for i := 0; i < n; i++ {
		q = q.Next
	}

	for q.Next != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next
	return dummy.Next
}
