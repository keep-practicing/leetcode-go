/*
226. Invert Binary Tree
https://leetcode.com/problems/invert-binary-tree/

Invert a binary tree.
*/

package invertbinarytree

// TreeNode binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
// Time complexity: O(n), where n is the node's number of the tree.
// Space complexity: O(h), where h is the height of the tree.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
