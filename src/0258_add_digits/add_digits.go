/*
258. Add Digits
https://leetcode.com/problems/add-digits/

Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
*/
// time: 2019-01-07

package ad

// time complexity: O(	log num	)
// space complexity: O(1)
func addDigits(num int) int {
	// num is a non-negative integer
	for num > 9 {
		num = performAdd(num)
	}
	return num
}

func performAdd(num int) (res int) {
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return
}

// time complexity: O(1)
// space complexity: O(1)
func addDigits1(num int) int {
	/*
		0	1	2	3	4	5	6	7	8		9
		0	1	2	3	4	5	6	7	8		9
		10	11	12	13	14	15	16	17	18		19
		1	2	3	4	5	6	7	8	9		10/1
		20	21	22	23	24	25	26	27	28		29
		2	3	4	5	6	7	8	9	10/1	11/2
	*/
	return (num-1)%9 + 1
}
