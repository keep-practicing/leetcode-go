/*
257. Binary Tree Paths
https://leetcode.com/problems/binary-tree-paths/

Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.
*/
// time: 2019-01-07

package btp

import "strconv"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
// time complexity: O(n), where n is the nodes number in the tree.
// space complexity: O(h), where h is the height of the tree.
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var res []string

	if root.Left == nil && root.Right == nil {
		return append(res, strconv.Itoa(root.Val))
	}

	if root.Left != nil {
		for _, i := range binaryTreePaths(root.Left) {
			res = append(res, strconv.Itoa(root.Val)+"->"+i)
		}
	}
	if root.Right != nil {
		for _, i := range binaryTreePaths(root.Right) {
			res = append(res, strconv.Itoa(root.Val)+"->"+i)
		}
	}
	return res
}
