package ps3

import "testing"

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

func TestPathSum(t *testing.T) {
	root := createBinaryTree([]int{10, 5, -3, 3, 2, 9, 11, 3, -2, 3, 1})
	sum := 8
	expected := 3

	if res := pathSum(root, sum); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
