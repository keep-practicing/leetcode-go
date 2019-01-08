/*
557. Reverse Words in a String III
https://leetcode.com/problems/reverse-words-in-a-string-iii/

Given a string, you need to reverse the order of characters in each word within a sentence
while still preserving whitespace and initial word order.

Note: In the string, each word is separated by single space and there will not be any extra space in the string.
*/
// time: 2019-01-08

package rwias3

import "strings"

// split and reverse
// time complexity: O(len(s))
// space complexity: O(n)
func reverseWords(s string) string {
	words := strings.Split(s, " ")

	var temp []string
	for _, word := range words {
		temp = append(temp, reverseString(word))
	}
	return strings.Join(temp, " ")
}

func reverseString(s string) string {
	runes := []rune(s)

	for start, end := 0, len(runes)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}
	return string(runes)
}
