/*
437. Path Sum III
https://leetcode.com/problems/path-sum-iii/

You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf,
but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.
*/
// 2019-01-08

package ps3

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
// time complexity: O(n), where n is the nodes number of the tree.
// space complexity: O(h), where h is the height of the tree.
func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	return findPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func findPath(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}

	res := 0
	if val == node.Val {
		res += 1
	}

	res += findPath(node.Left, val-node.Val)
	res += findPath(node.Right, val-node.Val)
	return res
}
