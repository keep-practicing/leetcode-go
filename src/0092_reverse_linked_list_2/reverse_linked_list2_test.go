package reverselinkedlist2

import (
	"reflect"
	"testing"
)

func createSingleList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return head.Next
}

func TestReverseBetween(t *testing.T) {
	head := createSingleList([]int{1, 2, 3, 4, 5})
	m := 2
	n := 4
	expected := createSingleList([]int{1, 4, 3, 2, 5})

	if res := reverseBetween(head, m, n); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
