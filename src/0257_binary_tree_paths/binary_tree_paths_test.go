package btp

import (
	"reflect"
	"testing"
)

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

func TestBinaryTreePaths(t *testing.T) {
	testCases := []*TreeNode{
		createBinaryTree([]int{1, 2, 3, 4, 5}),
		nil,
	}
	expected := [][]string{
		{"1->2->4", "1->2->5", "1->3"},
		{},
	}
	for index, root := range testCases {
		if res := binaryTreePaths(root); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
