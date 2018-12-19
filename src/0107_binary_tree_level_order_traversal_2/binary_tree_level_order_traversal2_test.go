package binarytreelevelordertraversal2

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

func TestBinaryTreeLevelOrderTraversal2(t *testing.T) {
	testData := [][]int{
		{3, 9, 20, 15, 17},
	}
	expectedData := [][][]int{
		{
			{15, 17},
			{9, 20},
			{3},
		},
	}

	for index, data := range testData {
		if res := levelOrderBottom(createBinaryTree(data)); !reflect.DeepEqual(res, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], res)
		}
	}
}
