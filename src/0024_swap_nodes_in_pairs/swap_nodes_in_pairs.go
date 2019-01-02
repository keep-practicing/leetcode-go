/*
24. Swap Nodes in Pairs
https://leetcode.com/problems/swap-nodes-in-pairs/

Given a linked list, swap every two adjacent nodes and return its head.
*/
// time: 2019-01-02

package swapnodesinpairs

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n/2), where n is len of linked list.
// space complexity: O(1)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := node1.Next
		node3 := node2.Next

		node2.Next = node1
		node1.Next = node3
		cur.Next = node2
		cur = node1
	}
	return dummy.Next
}
