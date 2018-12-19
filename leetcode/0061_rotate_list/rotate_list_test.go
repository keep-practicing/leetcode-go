package rotatelist

import (
	"reflect"
	"testing"
)

func createSingleLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, j := range arr {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}
	return head.Next
}

func TestRotateList(t *testing.T) {
	testDatas := []struct {
		name     string
		arg      *ListNode
		k        int
		expected *ListNode
	}{
		{
			name:     "one",
			arg:      createSingleLinkedList([]int{1, 2, 3, 4}),
			k:        2,
			expected: createSingleLinkedList([]int{3, 4, 1, 2}),
		},
		{
			name:     "two",
			arg:      createSingleLinkedList([]int{1, 2, 3}),
			k:        3,
			expected: createSingleLinkedList([]int{1, 2, 3}),
		},
		{
			name:     "three",
			arg:      createSingleLinkedList([]int{1, 2, 3}),
			k:        5,
			expected: createSingleLinkedList([]int{2, 3, 1}),
		},
		{
			name:     "four",
			arg:      createSingleLinkedList([]int{1, 2, 3}),
			k:        0,
			expected: createSingleLinkedList([]int{1, 2, 3}),
		},
		{
			name:     "five",
			arg:      nil,
			k:        0,
			expected: nil,
		},
		{
			name:     "four",
			arg:      createSingleLinkedList([]int{1}),
			k:        5,
			expected: createSingleLinkedList([]int{1}),
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := rotateRight(testData.arg, testData.k); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}
