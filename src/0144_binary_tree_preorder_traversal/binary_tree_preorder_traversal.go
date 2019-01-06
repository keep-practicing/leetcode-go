/*
144. Binary Tree Preorder Traversal
https://leetcode.com/problems/binary-tree-preorder-traversal/

Given a binary tree, return the preorder traversal of its nodes' values.
*/
// time: 2019-01-06

package btpot

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
// time complexity: O(n), where n is nodes numbers in the tree.
// space complexity: O(h), where h is the height of the tree.
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	res := []int{root.Val}
	res = append(res, left...)
	res = append(res, right...)
	return res
}

// iterative
// time complexity: O(n), where n is nodes numbers in the tree.
// space complexity: O(h), where h is the height of the tree.
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	var res []int
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
