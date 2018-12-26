/*
345. Reverse Vowels of a String
https://leetcode.com/problems/reverse-vowels-of-a-string/

Write a function that takes a string as input and reverse only the vowels of a string.

*/
// time: 2018-12-26

package reverse_vowels

// time complexity: O(n)
// space complexity: O(1)
func reverseVowels(s string) string {
	var (
		l int
		r = len(s) - 1
	)

	for r > l {
		for r >= 0 && !isVowel(s[r]) {
			r--
		}

		for l < len(s) && !isVowel(s[l]) {
			l++
		}

		if l >= r {
			break
		}
		charL := s[l]
		charR := s[r]
		s = s[:r] + string(charL) + s[r+1:]
		s = s[:l] + string(charR) + s[l+1:]
		r--
		l++
	}
	return s
}

// time complexity: O(n)
// space complexity: O(n)
func reverseVowels1(s string) string {
	bytes := []byte(s)
	var (
		l int
		r = len(bytes) - 1
	)
	for r > l {
		for r >= 0 && !isVowel(s[r]) {
			r--
		}
		for l < len(bytes) && !isVowel(s[l]) {
			l++
		}
		if l >= r {
			break
		}

		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}

func isVowel(char byte) bool {
	vowels := [...]byte{'a', 'o', 'e', 'i', 'u', 'A', 'O', 'E', 'I', 'U'}

	for _, k := range vowels {
		if char == k {
			return true
		}
	}
	return false
}
