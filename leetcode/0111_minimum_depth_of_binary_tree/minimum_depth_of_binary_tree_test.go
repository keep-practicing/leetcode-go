package minimumdepthofbinarytree

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
func TestMinDepth(t *testing.T) {
	testData := [][]int{
		{3, 9, 20, 15, 7},
	}

	expectedData := []int{2}

	for index, data := range testData {
		if res := minDepth(createBinaryTree(data)); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
