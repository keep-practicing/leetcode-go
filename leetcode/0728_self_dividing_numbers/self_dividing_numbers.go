/*
728. Self Dividing Numbers

Source: https://leetcode.com/problems/self-dividing-numbers/

A self-dividing number is a number that is divisible by every digit it contains.

For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

Also, a self-dividing number is not allowed to contain the digit zero.

Given a lower and upper number bound, output a list of every possible self dividing number, including the bounds if possible.

Example 1:
Input:
left = 1, right = 22
Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
Note:

The boundaries of each input argument are 1 <= left <= right <= 10000.

Time: 2018-12-18
*/

package selfdividingnumbers

import "strconv"

//Time complexity: O(right-left+1)
// Space complexity: O(1)
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)

	for num := left; num <= right; num++ {
		flag := true
		strNum := strconv.Itoa(num)

		for j := 0; j < len(strNum); j++ {
			if divisor := int(strNum[j] - '0'); divisor == 0 || num%divisor != 0 {
				flag = false
			}
		}
		if flag {
			res = append(res, num)
		}
	}
	return res
}
