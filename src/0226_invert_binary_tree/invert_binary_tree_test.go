package invertbinarytree

import (
	"reflect"
	"testing"
)

func createBinaryTree(nums []interface{}) *TreeNode {
	return performCreate(nums, 0)
}

func performCreate(nums []interface{}, index int) *TreeNode {
	if !(index < len(nums)) || nums[index] == nil {
		return nil
	}

	tree := TreeNode{}
	if num, ok := nums[index].(int); ok { // type assertion
		tree.Val = num
	}
	tree.Left = performCreate(nums, 2*index+1)
	tree.Right = performCreate(nums, 2*index+2)
	return &tree
}

func levelOrderTraversal(root *TreeNode) []interface{} {
	res := make([]interface{}, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func TestInvertTree(t *testing.T) {
	testData := [][]interface{}{
		{4, 2, 7, 1, 3, 6, 9},
	}

	expectedData := [][]interface{}{
		[]interface{}{4, 7, 2, 9, 6, 3, 1},
	}

	for index, data := range testData {
		tree := createBinaryTree(data)
		if res := levelOrderTraversal(invertTree(tree)); !reflect.DeepEqual(res, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], res)
		}
	}
}
