/*
237. Delete Node in a Linked List
https://leetcode.com/problems/delete-node-in-a-linked-list/

Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Note:
The linked list will have at least two elements.
All of the nodes' values will be unique.
The given node will not be the tail and it will always be a valid node of the linked list.
Do not return anything from your function.
*/
// time: 2019-01-07

package dniall

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(1)
// space complexity: O(1)
func deleteNode(node *ListNode) {
	// The linked list will have at least two elements.
	// All the nodes' values will be unique.
	// The given node will not be the tail and it will always be a valid node of the linked list.
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
