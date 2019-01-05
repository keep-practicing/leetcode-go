/*
102. Binary Tree Level Order Traversal
https://leetcode.com/problems/binary-tree-level-order-traversal/

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
*/
// time: 2019-01-04

package btlot

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueEle struct {
	Level int
	Ele   *TreeNode
}

// bfs
// time complexity: O(n), where n is number of tree nodes.
// space complexity: O(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		queue []queueEle
		res   = make([][]int, 0)
	)

	queue = append(queue, queueEle{Level: 0, Ele: root})

	for len(queue) > 0 {
		node := queue[0].Ele
		level := queue[0].Level
		queue = queue[1:]

		if len(res) > level {
			res[level] = append(res[level], node.Val)
		} else {
			res = append(res, []int{node.Val})
		}

		if node.Left != nil {
			queue = append(queue, queueEle{Level: level + 1, Ele: node.Left})
		}
		if node.Right != nil {
			queue = append(queue, queueEle{Level: level + 1, Ele: node.Right})
		}

	}
	return res
}
