package rnthnfeol

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type arg struct {
		head *ListNode
		n    int
	}

	testCase := []arg{
		{head: createSingleLinkedList([]int{1, 2, 3, 4, 5}), n: 2},
		{head: createSingleLinkedList([]int{1, 2, 3, 4, 5}), n: 0},
	}

	expected := []*ListNode{
		createSingleLinkedList([]int{1, 2, 3, 5}),
		createSingleLinkedList([]int{1, 2, 3, 4, 5}),
	}

	for index, data := range testCase {
		if res := removeNthFromEnd(data.head, data.n); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected")
		}
	}
}

func createSingleLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, j := range arr {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}
	return head.Next
}
