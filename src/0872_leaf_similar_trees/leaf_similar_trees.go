/*
872. Leaf-Similar Trees
https://leetcode.com/problems/leaf-similar-trees/
*/
// time: 2019-01-08

package lst

import (
	"reflect"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(n1 + n2), where n is nodes number in the tree.
// space complexity: O(h1 + h2), where h is height of the tree.
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := make([]int, 0)
	l2 := make([]int, 0)

	dfs(root1, &l1)
	dfs(root2, &l2)

	return reflect.DeepEqual(l1, l2)
}

func dfs(root *TreeNode, l *[]int) {
	if nil == root {
		return
	}

	if root.Left == nil && root.Right == nil {
		*l = append(*l, root.Val)
	}
	dfs(root.Left, l)
	dfs(root.Right, l)
}
