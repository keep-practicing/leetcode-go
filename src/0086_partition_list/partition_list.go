/*
86. Partition List
https://leetcode.com/problems/partition-list/

Given a linked list and a value x,
partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.
*/
// time: 2019-01-04

package partitionlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// time complexity: O(n)
// space complexity: O(1)
func partition(head *ListNode, x int) *ListNode {
	var (
		dummyHead1 = &ListNode{}
		dummyHead2 = &ListNode{}
		cur1       = dummyHead1
		cur2       = dummyHead2
	)

	for cur := head; cur != nil; {
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur1.Next
			cur = cur.Next
			cur1.Next = nil
		} else {
			cur2.Next = cur
			cur2 = cur2.Next
			cur = cur.Next
			cur2.Next = nil
		}
	}
	cur1.Next = dummyHead2.Next
	return dummyHead1.Next
}
