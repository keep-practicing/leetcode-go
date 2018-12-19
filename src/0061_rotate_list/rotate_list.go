/*
61. Rotate List

Source: https://leetcode.com/problems/rotate-list/

Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:
	Input: 1->2->3->4->5->NULL, k = 2
	Output: 4->5->1->2->3->NULL
	Explanation:
	rotate 1 steps to the right: 5->1->2->3->4->NULL
	rotate 2 steps to the right: 4->5->1->2->3->NULL

Example 2:
	Input: 0->1->2->NULL, k = 4
	Output: 2->0->1->NULL
	rotate 1 steps to the right: 2->0->1->NULL
	Explanation:
	rotate 2 steps to the right: 1->2->0->NULL
	otate 3 steps to the right: 0->1->2->NULL
	rotate 4 steps to the right: 2->0->1->NULL
*/

package rotatelist

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	var (
		listSize   int
		count      int
		p          = head
		rotateNode = head
	)

	for p != nil && count < k {
		p = p.Next
		listSize++
		count++
	}

	if p == nil {
		k = k % listSize
		if k == 0 {
			return head
		}
		p = head
		for count = 0; count < k; count++ {
			p = p.Next
		}
	}

	for p.Next != nil {
		rotateNode = rotateNode.Next
		p = p.Next
	}

	p.Next = head
	head = rotateNode.Next
	rotateNode.Next = nil
	return head
}
