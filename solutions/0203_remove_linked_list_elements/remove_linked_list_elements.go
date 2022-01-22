/*
203. Remove Linked List Elements
https://leetcode.com/problems/remove-linked-list-elements/

Remove all elements from a linked list of integers that have value val.
*/
// time: 2019-01-07

package rlle

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	cur := dummyHead

	for cur.Next != nil {
		if val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

// recursive
// time complexity: O(n)
// space complexity: O(n)
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements1(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
