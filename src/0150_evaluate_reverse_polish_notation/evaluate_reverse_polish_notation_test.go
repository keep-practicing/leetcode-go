package evaluatereversepolishnotation

import "testing"

func TestEvalRPN(t *testing.T) {
	testCases := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		{"4", "3", "-"},
	}

	expected := []int{9, 6, 22, 1}

	for index, tokens := range testCases {
		if res := evalRPN(tokens); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
