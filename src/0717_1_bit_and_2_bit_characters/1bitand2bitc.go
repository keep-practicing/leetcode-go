/*
717. 1-bit and 2-bit Characters
https://leetcode.com/problems/1-bit-and-2-bit-characters/

We have two special characters.
The first character can be represented by one bit 0.
The second character can be represented by two bits (10 or 11).

Now given a string represented by several bits.
Return whether the last character must be a one-bit character or not.
The given string will always end with a zero.

Note:
1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.
*/
// time: 2018-12-28

package onebitandtwobitcharacters

// time complexity: O(n)
// space complexity: O(1)
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if 1 == n {
		return true
	}

	cur := 0
	flag := false

	for cur < n {
		if 0 == bits[cur] {
			cur++
		} else {
			cur += 2
		}
		if cur == n-1 {
			flag = true
		}
	}
	return flag
}
