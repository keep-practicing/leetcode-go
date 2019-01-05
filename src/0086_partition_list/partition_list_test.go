package partitionlist

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
	testCase := createSingleLinkedList([]int{1, 4, 3, 2, 5, 2})

	expected := createSingleLinkedList([]int{1, 2, 2, 4, 3, 5})
	x := 3
	if res := partition(testCase, x); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
