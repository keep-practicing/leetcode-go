/*
404. Sum of Left Leaves
https://leetcode.com/problems/sum-of-left-leaves/

Find the sum of all left leaves in a given binary tree.
*/
// time: 2019-01-07

package soll

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(n), where n is nodes number in the tree.
// space complexity: O(h), where h is height of the tree.
func sumOfLeftLeaves(root *TreeNode) int {
	return performSum(root, false)
}

func performSum(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return performSum(root.Left, true) + performSum(root.Right, false)
}
