/*
67. Add Binary
https://leetcode.com/problems/add-binary/

Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.

Example:
	Input: a = "11", b = "1"
	Output: "100"
*/
// time: 2018-12-19

package addbinary

import (
	"strconv"
)

// Time complexity: O(	max(	len(a),    len(b)	)	)
// Space complexity: O(1)
func addBinary(a string, b string) string {
	var (
		lenA  = len(a)
		lenB  = len(b)
		carry int
		res   = ""
	)
	for lenA > 0 && lenB > 0 {
		tmp := int(a[lenA-1]-'0') + int(b[lenB-1]-'0') + carry
		res = strconv.Itoa(tmp%2) + res
		carry = tmp / 2
		lenA--
		lenB--
	}

	if lenA == 0 {
		for lenB > 0 {
			tmp := int(b[lenB-1]-'0') + carry
			res = strconv.Itoa(tmp%2) + res
			carry = tmp / 2
			lenB--
		}
	}

	if lenB == 0 {
		for lenA > 0 {
			tmp := int(a[lenA-1]-'0') + carry
			res = strconv.Itoa(tmp%2) + res
			carry = tmp / 2
			lenA--
		}
	}

	if carry == 1 {
		res = strconv.Itoa(carry) + res
	}
	return res
}
