package compareversionnumbers

import "testing"

func TestCompareVersion(t *testing.T) {
	type arg struct {
		version1 string
		version2 string
	}

	testCases := []arg{
		{version1: "0.1", version2: "1.1"},
		{version1: "1.0.1", version2: "1"},
		{version1: "1.01", version2: "1.001"},
	}
	expected := []int{-1, 1, 0}

	for index, data := range testCases {
		if res := compareVersion(data.version1, data.version2); res != expected[index] {
			t.Errorf("expected %d, got %d", expected[index], res)
		}
	}
}
