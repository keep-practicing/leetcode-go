/*
328. Odd Even Linked List
https://leetcode.com/problems/odd-even-linked-list/

Given a singly linked list, group all odd nodes together followed by the even nodes.
Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place.
The program should run in O(1) space complexity and O(nodes) time complexity.

Note:
The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
*/
// time: 2019-01-07

package oddevenlinkedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		odd      = head
		evenHead = head.Next
		even     = evenHead
	)
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
