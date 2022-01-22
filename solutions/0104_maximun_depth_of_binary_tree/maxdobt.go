/*
104. Maximum Depth of Binary Tree
https://leetcode.com/problems/maximum-depth-of-binary-tree/

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along
the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.
*/
// time: 2019-01-06

package maxdobt

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
// time complexity: O(n), where n is the nodes number in the tree.
// space complexity: O(h), where h is the height of the tree.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	var max int
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}
