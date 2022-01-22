package oddevenlinkedlist

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

func TestOddEvenList(t *testing.T) {
	testCases := []*ListNode{
		createSinglyLinkedList([]int{2, 1, 3, 5, 6, 4, 7}),
		nil,
	}
	expected := []*ListNode{
		createSinglyLinkedList([]int{2, 3, 6, 7, 1, 5, 4}),
		nil,
	}

	for index, head := range testCases {
		if res := oddEvenList(head); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
