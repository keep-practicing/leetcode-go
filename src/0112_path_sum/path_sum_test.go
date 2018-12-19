package pathsum

import (
	"os"
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

type arg struct {
	nums []interface{}
	sum  int
}

func TestHasPathSum(t *testing.T) {
	testData := []arg{
		arg{nums: []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, sum: 22},
		arg{nums: []interface{}{}, sum: 0},
	}
	expectedData := []bool{true, false}
	for index, data := range testData {
		if res := hasPathSum(createBinaryTree(data.nums), data.sum); res != expectedData[index] {
			t.Errorf("expected %t, got %t", expectedData[index], res)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
