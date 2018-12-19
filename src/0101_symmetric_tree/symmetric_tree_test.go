package symmetrictree

import (
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

func TestIsSymmetric(t *testing.T) {
	testData := [][]int{
		{1, 2, 2, 3, 4, 4, 3},
		{1, 2, 2, 3, 3},
		{},
	}
	expectedData := []bool{true, false, true}

	for index, data := range testData {
		if res := isSymmetric(createBinaryTree(data)); res != expectedData[index] {
			t.Errorf("expected %t, got %t", expectedData[index], res)
		}
	}
}
