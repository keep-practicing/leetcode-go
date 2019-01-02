package implstr

import "testing"

func TestStrStr(t *testing.T) {
	type arg struct {
		haystack string
		needle   string
	}

	testCases := []arg{
		{haystack: "hello", needle: "ll"},
		{haystack: "aaaaa", needle: "bba"},
		{haystack: "hello"},
	}

	expected := []int{2, -1, 0}
	for index, data := range testCases {
		if res := strStr(data.haystack, data.needle); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
