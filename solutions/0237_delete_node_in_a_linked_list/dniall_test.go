package dniall

import (
	"reflect"
	"testing"
)

func createSinglyLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head.Next
}

func TestDeleteNode(t *testing.T) {
	root := createSinglyLinkedList([]int{4, 5, 1, 9})
	node := root.Next
	expected := createSinglyLinkedList([]int{4, 1, 9})
	if deleteNode(node); !reflect.DeepEqual(expected, root) {
		t.Errorf("expected %v, got %v", expected, root)
	}
}
