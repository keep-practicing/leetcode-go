package rdfsl

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []*ListNode{
		createSingleLinkedList([]int{1, 2, 3, 3, 4, 4, 5}),
		createSingleLinkedList([]int{1, 1, 1, 2, 3}),
	}

	expected := []*ListNode{
		createSingleLinkedList([]int{1, 2, 5}),
		createSingleLinkedList([]int{2, 3}),
	}

	for index, head := range testCases {
		if res := deleteDuplicates(head); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}

func createSingleLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head.Next
}
