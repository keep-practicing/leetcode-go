package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	testData := []string{"()", "(((((())))))", "()()()()", "(((((((()", "((()(())))", "", ")(", "}{", "][", "(][)"}
	expectedData := []bool{true, true, true, false, true, true, false, false, false, false}

	for i, s := range testData {
		result := isValid(s)
		if result != expectedData[i] {
			t.Error("测试不通过")
		}
	}
}
