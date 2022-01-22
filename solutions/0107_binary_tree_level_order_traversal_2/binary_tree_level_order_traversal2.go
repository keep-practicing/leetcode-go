/*
107. Binary Tree Level Order Traversal II
https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
*/

package binarytreelevelordertraversal2

// TreeNode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var record [][]int
	traversal(root, 0, &record)
	var reversed [][]int
	for i := len(record) - 1; i >= 0; i-- {
		reversed = append(reversed, record[i])
	}
	return reversed
}

func traversal(root *TreeNode, index int, record *[][]int) {
	if root == nil {
		return
	}
	if len(*record) == index {
		*record = append(*record, make([]int, 0))
	}
	(*record)[index] = append((*record)[index], root.Val)
	traversal(root.Left, index+1, record)
	traversal(root.Right, index+1, record)
	// (*record)[len(*record)-1-index] = append((*record)[len(*record)-1-index], root.Val)
	// 这个方法用java可以实现，go不行，猜测可能跟java的递归实现有关。
}
