package sametree

import "testing"

type arg struct {
	p []int
	q []int
}

func createBinaryTree(nums []int) *TreeNode {
	return performCreate(nums, 0)
}

func performCreate(nums []int, index int) *TreeNode {
	if index >= len(nums) {
		return nil
	}

	tree := TreeNode{Val: nums[index]}
	tree.Left = performCreate(nums, 2*index+1)
	tree.Right = performCreate(nums, 2*index+2)
	return &tree
}

func TestSameTree(t *testing.T) {
	testData := []arg{
		{p: []int{1, 2, 3}, q: []int{1, 2, 3}},
		{p: []int{1, 2, 3}, q: []int{1, 2}},
		{p: []int{1, 2, 3}, q: []int{1, 2, 4}},
	}

	expectedData := []bool{true, false, false}

	for index, data := range testData {
		if res := isSameTree(createBinaryTree(data.p), createBinaryTree(data.q)); res != expectedData[index] {
			t.Errorf("expected %t, got %t", expectedData[index], res)
		}
	}
}
