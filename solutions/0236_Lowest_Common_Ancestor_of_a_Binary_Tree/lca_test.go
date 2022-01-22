package lca

import "testing"

func TestLCA(t *testing.T) {
	type arg struct {
		root, p, q *TreeNode
	}

	tree := createBinaryTree([]int{6, 2, 8, 0, 4, 7, 9, -1, 1, 3, 5})

	testCases := []arg{
		{root: tree, p: tree.Left.Right.Left, q: tree.Left.Right.Right},
		{p: &TreeNode{Val: 2}, q: &TreeNode{Val: 8}},
	}
	expected := []*TreeNode{{Val: 4}, nil}

	for index, data := range testCases {
		if res := lowestCommonAncestor(data.root, data.p, data.q); res != nil && res.Val != expected[index].Val {
			t.Errorf("expected %v, got %v", expected[index], res)
		} else if res == nil && res != expected[index] {
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
	root := &TreeNode{Val: nums[index]}
	root.Left = performCreate(nums, 2*index+1)
	root.Right = performCreate(nums, 2*index+2)
	return root
}
