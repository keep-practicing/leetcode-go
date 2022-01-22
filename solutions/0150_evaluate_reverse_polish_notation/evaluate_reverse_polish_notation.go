/*
150. Evaluate Reverse Polish Notation
https://leetcode.com/problems/evaluate-reverse-polish-notation/

Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid.
That means the expression would always evaluate to a result and there won't be any divide by zero operation.
*/
// time: 2019-01-07

package evaluatereversepolishnotation

import "strconv"

// stack
// time complexity: O(n)
// space complexity: O(n)
func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := -1
	for i := 0; i < len(tokens); i++ {
		switch ch := tokens[i]; ch {
		case "+":
			stack[top-1] += stack[top]
			top--
		case "-":
			stack[top-1] -= stack[top]
			top--
		case "*":
			stack[top-1] *= stack[top]
			top--
		case "/":
			stack[top-1] /= stack[top]
			top--
		default:
			top++
			stack[top], _ = strconv.Atoi(ch)
		}
	}
	return stack[0]
}
