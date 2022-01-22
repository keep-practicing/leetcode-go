package removeoutmostparentheses

import (
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	testCases := []map[string]string{
		{"case": "(()())(())", "expected": "()()()"},
		{"case": "(()())(())(()(()))", "expected": "()()()()(())"},
		{"case": "", "expected": ""},
		{"case": "()()", "expected": ""},
	}

	for _, testCase := range testCases {
		if testCase["expected"] != removeOuterParentheses(testCase["case"]) {
			t.Errorf("hello")
		}
	}
}
