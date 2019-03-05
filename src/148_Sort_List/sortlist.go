/*
148. Sort List
https://leetcode.com/problems/sort-list/
*/
// time: 2019-03-04

package sortlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// merge sort
// time complexity: O(n * log(n))
// using recursion, the system stack is used.
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(headA, headB *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for headA != nil && headB != nil {
		if headA.Val > headB.Val {
			tail.Next = headB
			headB = headB.Next
		} else {
			tail.Next = headA
			headA = headA.Next
		}
		tail = tail.Next
	}
	if headA != nil {
		tail.Next = headA
	} else {
		tail.Next = headB
	}
	return dummy.Next
}
