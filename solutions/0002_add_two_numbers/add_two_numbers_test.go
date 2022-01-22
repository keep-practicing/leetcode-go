package addtwonumbers

import (
	"os"
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

func TestAddTwoNumbers1(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name:     "one",
			arg1:     createSingleLinkedList([]int{2, 4, 3}),
			arg2:     createSingleLinkedList([]int{5, 6, 4}),
			expected: createSingleLinkedList([]int{7, 0, 8}),
		},
		{
			name:     "two",
			arg1:     createSingleLinkedList([]int{5}),
			arg2:     createSingleLinkedList([]int{5}),
			expected: createSingleLinkedList([]int{0, 1}),
		},
		{
			name:     "three",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{9, 9, 9}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "four",
			arg1:     createSingleLinkedList([]int{9, 9, 9}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "five",
			arg1:     createSingleLinkedList([]int{4, 3, 1}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
		{
			name:     "six",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{4, 3, 1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := addTwoNumbers1(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name:     "one",
			arg1:     createSingleLinkedList([]int{2, 4, 3}),
			arg2:     createSingleLinkedList([]int{5, 6, 4}),
			expected: createSingleLinkedList([]int{7, 0, 8}),
		},
		{
			name:     "two",
			arg1:     createSingleLinkedList([]int{5}),
			arg2:     createSingleLinkedList([]int{5}),
			expected: createSingleLinkedList([]int{0, 1}),
		},
		{
			name:     "three",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{9, 9, 9}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "four",
			arg1:     createSingleLinkedList([]int{9, 9, 9}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "five",
			arg1:     createSingleLinkedList([]int{4, 3, 1}),
			arg2:     createSingleLinkedList([]int{1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
		{
			name:     "six",
			arg1:     createSingleLinkedList([]int{1}),
			arg2:     createSingleLinkedList([]int{4, 3, 1}),
			expected: createSingleLinkedList([]int{5, 3, 1}),
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := addTwoNumbers2(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
