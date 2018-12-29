/*
7. Reverse Integer
Given a 32-bit signed integer, reverse digits of an integer.

*/
// time: 2018-12-29

package reverseinteger

// time complexity: O(log 10 x ) = O(log x)
// space complexity: O(1)
func reverse(x int) int {
	if 0 == x {
		return x
	}

	isPositive := true
	if x < 0 {
		isPositive = false
		x *= -1
	}

	res := 0

	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}

	if !isPositive {
		res *= -1
	}

	if res < -1<<31 || res > (1<<31)-1 {
		return 0
	}
	return res
}
