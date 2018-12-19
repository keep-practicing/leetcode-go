/*
25. Reverse Nodes in k-Group

Source: https://leetcode.com/problems/reverse-nodes-in-k-group/

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:
	Only constant extra memory is allowed.
	You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

package reversenodesinkgroup

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	p := new(ListNode)
	s := new(ListNode)
	l := k
	p.Next = head
	head = p
	s = p

	for s != nil && k != 0 {
		s = s.Next
		k--
	}

	for s != nil {
		reverse(&p, &s)
		k = l
		for s != nil && k != 0 {
			s = s.Next
			k--
		}
	}
	return head.Next
}

func reverse(p **ListNode, s **ListNode) {
	var tmp *ListNode
	prev := (*p).Next
	tail := (*p).Next
	flag := (*s).Next

	(*p).Next = nil
	for prev != flag {
		tmp = (*p).Next
		(*p).Next = prev
		prev = prev.Next
		(*p).Next.Next = tmp
	}
	tail.Next = prev
	*p = tail
	*s = tail
}
