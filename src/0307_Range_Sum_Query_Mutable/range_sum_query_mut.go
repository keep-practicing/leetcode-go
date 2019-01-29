/*
307. Range Sum Query - Mutable
https://leetcode.com/problems/range-sum-query-mutable/

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Note:
The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly.
*/
// time: 2019-01-29
// 线段树时间复杂度为：O(log n)

package rsqm

// NumArray 线段树
type NumArray struct {
	segmentTree []int
	data        []int
}

// Constructor 线段树生成函数
func Constructor(nums []int) NumArray {
	na := NumArray{segmentTree: make([]int, len(nums)*4), data: nums[:]}
	na.buildSegmentTree(0, 0, len(na.data)-1)
	return na
}

func (na *NumArray) buildSegmentTree(treeIndex, left, right int) {
	// 线段树处理的数据不能为空。
	//if right < left {
	//	return
	//}
	if left == right {
		na.segmentTree[treeIndex] = na.data[left]
		return
	}
	mid := left + (right-left)>>1
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2
	na.buildSegmentTree(leftTreeIndex, left, mid)
	na.buildSegmentTree(rightTreeIndex, mid+1, right)
	na.segmentTree[treeIndex] = na.segmentTree[leftTreeIndex] + na.segmentTree[rightTreeIndex]
}

// Update 线段树更新操作。
func (na *NumArray) Update(i int, val int) {
	na.setter(0, 0, len(na.data)-1, i, val)
}

func (na *NumArray) setter(treeIndex, left, right, index, val int) {
	if left == right {
		na.segmentTree[treeIndex] = val
		return
	}
	mid := left + (right-left)>>1
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2
	if index >= mid+1 {
		na.setter(rightTreeIndex, mid+1, right, index, val)
	} else {
		na.setter(leftTreeIndex, left, mid, index, val)
	}
	na.segmentTree[treeIndex] = na.segmentTree[leftTreeIndex] + na.segmentTree[rightTreeIndex]
}

// SumRange 线段树查询
func (na *NumArray) SumRange(i int, j int) int {
	return na.query(0, 0, len(na.data)-1, i, j)
}

func (na *NumArray) query(treeIndex, left, right, queryL, queryR int) int {
	if left == queryL && right == queryR {
		return na.segmentTree[treeIndex]
	}
	mid := left + (right-left)>>1
	leftTreeIndex := 2*treeIndex + 1
	rightTreeIndex := 2*treeIndex + 2
	if queryL >= mid+1 {
		return na.query(rightTreeIndex, mid+1, right, queryL, queryR)
	} else if queryR <= mid {
		return na.query(leftTreeIndex, left, mid, queryL, queryR)
	} else {
		return na.query(leftTreeIndex, left, mid, queryL, mid) + na.query(rightTreeIndex, mid+1, right, mid+1, queryR)
	}
}
