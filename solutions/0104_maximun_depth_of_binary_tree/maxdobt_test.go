package maxdobt

import "testing"

func TestMaxDepth(t *testing.T) {
	root := createBinaryTree([]int{1, 2, 3, 4, 5})
	expected := 3

	if res := maxDepth(root); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}

}

func createBinaryTree(nums []int) *TreeNode {
	return performCreate(nums, 0)
}

func performCreate(nums []int, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}
	tree := &TreeNode{Val: nums[index]}
	tree.Left = performCreate(nums, 2*index+1)
	tree.Right = performCreate(nums, 2*index+2)
	return tree
}
