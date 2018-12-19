package reversenodesinkgroup

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

func TestReverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "one",
			args: args{
				head: createSingleLinkedList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			expected: createSingleLinkedList([]int{2, 1, 4, 3, 5}),
		},
		{
			name: "two",
			args: args{
				head: createSingleLinkedList([]int{1, 2, 3, 4, 5}),
				k:    3,
			},
			expected: createSingleLinkedList([]int{3, 2, 1, 4, 5}),
		},

		{
			name: "three",
			args: args{
				head: createSingleLinkedList([]int{1}),
				k:    2,
			},
			expected: createSingleLinkedList([]int{1}),
		},
	}
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			if result := reverseKGroup(data.args.head, data.args.k); !reflect.DeepEqual(result, data.expected) {
				t.Errorf("reverseKGroup() = %v, expected %v", result, data.expected)
			}
		})
	}
}
