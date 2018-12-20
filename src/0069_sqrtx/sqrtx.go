/*
69. Sqrt(x)
https://leetcode.com/problems/sqrtx/

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
*/

// time: 2018-12-20

package sqrtx

// binary search
// time complexity: O(logn)
// space complexity: O(1)
func mySqrt(x int) int {
	var (
		l int
		r = x // 在0～x范围内寻找平凡根
	)

	for l <= r {
		mid := l + (r-l)>>1 //位运算更高效
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			if (mid+1)*(mid+1) > x {
				return mid // 返回整数部分
			}
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
