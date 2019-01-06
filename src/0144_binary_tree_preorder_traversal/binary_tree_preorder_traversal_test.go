package btpot

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

func TestPreorderTraversal(t *testing.T) {
	testCases := []*TreeNode{
		nil,
		createBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8}),
	}
	expected := [][]int{
		{},
		{1, 2, 4, 8, 5, 3, 6, 7},
	}
	testFuncs := []func(node *TreeNode) []int{
		preorderTraversal,
		preorderTraversal1,
	}

	for _, testFunc := range testFuncs {
		for index, root := range testCases {
			if res := testFunc(root); !reflect.DeepEqual(res, expected[index]) {
				t.Errorf("expected %v, got %v", expected[index], res)
			}
		}
	}
}
