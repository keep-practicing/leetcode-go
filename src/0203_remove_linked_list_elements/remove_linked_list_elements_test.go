package rlle

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	head := createSinglyLinkedList([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	expected := createSinglyLinkedList([]int{1, 2, 3, 4, 5})

	if res := removeElements(head, val); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func createSinglyLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head.Next
}
