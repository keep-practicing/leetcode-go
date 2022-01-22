package lst

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

func TestLeafSimilar(t *testing.T) {
	type arg struct {
		root1, root2 *TreeNode
	}

	testCases := []arg{
		{root1: createBinaryTree([]int{3, 5, 1, 6, 2, 9, 8}), root2: createBinaryTree([]int{6, 4, 1, 6, 2, 9, 8})},
		{root1: createBinaryTree([]int{3, 5, 1, 6, 2, 9, 4}), root2: createBinaryTree([]int{6, 4, 1, 6, 2, 9, 8})},
	}

	expected := []bool{true, false}

	for index, data := range testCases {
		if res := leafSimilar(data.root1, data.root2); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
