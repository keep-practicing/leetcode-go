package rll

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := createSinglyLinkedList([]int{1, 2, 3, 4, 5})
	expected := createSinglyLinkedList([]int{5, 4, 3, 2, 1})
	if res := reverseList(head); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func TestReverseList1(t *testing.T) {
	head := createSinglyLinkedList([]int{1, 2, 3, 4, 5})
	expected := createSinglyLinkedList([]int{5, 4, 3, 2, 1})
	if res := reverseList1(head); !reflect.DeepEqual(res, expected) {
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
