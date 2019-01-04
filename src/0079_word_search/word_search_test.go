package wordsearch

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	testCases := []string{
		"ABCCED",
		"SEE",
		"ABCB",
	}
	expected := []bool{true, true, false}
	for index, word := range testCases {
		if res := exist(board, word); res != expected[index] {
			t.Errorf("expected %t, got %t", expected[index], res)
		}
	}
}
