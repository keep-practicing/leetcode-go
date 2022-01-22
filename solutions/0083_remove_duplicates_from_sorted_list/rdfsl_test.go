package rdfsl

import (
	"reflect"
	"testing"
)

func createSingleLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head.Next
}

func TestDeleteDuplicates(t *testing.T) {
	testCases := []*ListNode{
		createSingleLinkedList([]int{1, 1, 2}),
		createSingleLinkedList([]int{1, 1, 2, 3, 3}),
		nil,
	}
	expected := []*ListNode{
		createSingleLinkedList([]int{1, 2}),
		createSingleLinkedList([]int{1, 2, 3}),
		nil,
	}
	for index, head := range testCases {
		if res := deleteDuplicates(head); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
