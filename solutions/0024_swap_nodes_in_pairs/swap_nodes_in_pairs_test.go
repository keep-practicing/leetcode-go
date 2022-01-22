package swapnodesinpairs

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := createSingleList([]int{1, 2, 3, 4})
	expected := createSingleList([]int{2, 1, 4, 3})
	if res := swapPairs(head); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func createSingleList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, j := range nums {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}
	return head.Next
}
