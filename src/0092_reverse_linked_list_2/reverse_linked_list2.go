/*
92. Reverse Linked List II
https://leetcode.com/problems/reverse-linked-list-ii/

Reverse a linked list from position m to n. Do it in one-pass.
Note: 1 ≤ m ≤ n ≤ length of list.
*/
// time: 2019-01-04

package reverselinkedlist2

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n-m)
// space complexity: O(1)
func reverseBetween(head *ListNode, m, n int) *ListNode {
	var (
		dummy = &ListNode{}
		p     = dummy
		s     = p
	)
	dummy.Next = head

	// 1 ≤ m ≤ n ≤ length of list.
	for s != nil && n > 0 {
		s = s.Next
		n--
	}
	for s != nil && m > 1 {
		p = p.Next
		m--
	}
	reverse(&p, &s)
	return dummy.Next
}

func reverse(p, s **ListNode) {
	var (
		tmp  *ListNode
		prev = (*p).Next
		tail = (*p).Next
		flag = (*s).Next
	)
	(*p).Next = nil
	for prev != flag {
		tmp = (*p).Next
		(*p).Next = prev
		prev = prev.Next
		(*p).Next.Next = tmp
	}
	tail.Next = prev
	/*
		1->2->3->4->5
		1->3->2->4->5
		1->4->3->2->5
		1->5->4->3->2
	*/
}
