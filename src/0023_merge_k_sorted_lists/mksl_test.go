package mksl

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	testCases := [][]*ListNode{
		{
			createSingleList([]int{1, 4, 5}),
			createSingleList([]int{2, 3, 4}),
			createSingleList([]int{4, 6}),
		},
		{
			createSingleList([]int{}),
			createSingleList([]int{4, 5}),
			createSingleList([]int{}),
		},
		{
			createSingleList([]int{}),
			createSingleList([]int{}),
			createSingleList([]int{}),
		},
	}

	expected := []*ListNode{
		createSingleList([]int{1, 2, 3, 4, 4, 4, 5, 6}),
		createSingleList([]int{4, 5}),
		nil,
	}

	for index, lists := range testCases {
		if res := mergeKLists(lists); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}

func TestMergeKLists1(t *testing.T) {
	testCases := [][]*ListNode{
		{
			createSingleList([]int{1, 4, 5}),
			createSingleList([]int{1, 3, 4}),
			createSingleList([]int{2, 6}),
		},
		{
			createSingleList([]int{}),
			createSingleList([]int{4, 5}),
			createSingleList([]int{}),
		},
		{
			createSingleList([]int{}),
			createSingleList([]int{}),
			createSingleList([]int{}),
		},
	}

	expected := []*ListNode{
		createSingleList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		createSingleList([]int{4, 5}),
		nil,
	}

	for index, lists := range testCases {
		if res := mergeKLists1(lists); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
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
