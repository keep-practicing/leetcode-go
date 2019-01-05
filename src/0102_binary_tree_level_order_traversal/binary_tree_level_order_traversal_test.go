package btlot

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	testCases := []*TreeNode{
		createBinaryTree([]int{1, 2, 3, 4, 5}),
		nil,
	}
	expected := [][][]int{
		{{1}, {2, 3}, {4, 5}},
		{},
	}

	for index, root := range testCases {
		if res := levelOrder(root); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}

func createBinaryTree(nums []int) *TreeNode {
	return performCreate(nums, 0)
}

func performCreate(nums []int, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}
	tree := TreeNode{Val: nums[index]}
	tree.Left = performCreate(nums, 2*index+1)
	tree.Right = performCreate(nums, 2*index+2)
	return &tree
}
