package binarytreeinordertraversal

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

func TestInorderTraversal(t *testing.T) {
	testData := [][]int{
		{1, 2, 3},
	}
	expectedData := [][]int{
		{2, 1, 3},
	}

	for index, data := range testData {
		if res := inorderTraversal(createBinaryTree(data)); !reflect.DeepEqual(res, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], res)
		}
	}
}
