package sortlist

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	head := createSingleLinkedList([]int{-1, 5, 3, 4, 0})
	expected := createSingleLinkedList([]int{-1, 0, 3, 4, 5})
	if res := sortList(head); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func createSingleLinkedList(nums []int) *ListNode {
	dummy := ListNode{}
	cur := &dummy
	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return dummy.Next
}
