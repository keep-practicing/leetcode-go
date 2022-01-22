/*
94. Binary Tree Inorder Traversal
source: https://leetcode.com/problems/binary-tree-inorder-traversal/

Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package binarytreeinordertraversal

// TreeNode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 1)
	inorderTraversalHelp(root, &res)
	return res
}

func inorderTraversalHelp(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelp(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalHelp(root.Right, res)
}
