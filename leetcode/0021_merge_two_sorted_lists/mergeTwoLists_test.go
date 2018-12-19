package mergetwolists

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

func TestMergeTwoLists(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name:     "one",
			arg1:     createSingleLinkedList([]int{1, 3, 5}),
			arg2:     createSingleLinkedList([]int{2, 4, 6}),
			expected: createSingleLinkedList([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			name:     "two",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{2, 4, 6}),
			expected: createSingleLinkedList([]int{1, 2, 4, 6}),
		},
		{
			name:     "three",
			arg1:     nil,
			arg2:     nil,
			expected: nil,
		},
		{
			name:     "four",
			arg1:     createSingleLinkedList([]int{2, 4, 6}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{1, 2, 4, 6}),
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := mergeTwoLists(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("mergeTwoLists() = %v, expected %v", result, testData.expected)
			}
		})
	}
}
