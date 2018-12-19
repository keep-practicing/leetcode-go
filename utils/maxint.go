package utils

// CalcMaxInt calc max int from multi nums.
func CalcMaxInt(nums ...int) (res int) {
	// 此处也可使用堆排序，构建大顶堆，heapify之后直接取最大值.
	if len(nums) == 0 {
		return 0
	}

	res = MinInt

	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return
}
