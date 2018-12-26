package utils

// CalcMinInt calc min int from multi nums.
func CalcMinInt(nums ...int) (res int) {
	// 此处也可使用堆排序，构建小顶堆，heapify之后直接取最小值.
	if len(nums) == 0 {
		return 0
	}

	res = nums[0]

	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return
}
