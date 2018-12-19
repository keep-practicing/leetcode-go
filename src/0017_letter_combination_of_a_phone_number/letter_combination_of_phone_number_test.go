package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testData := []string{
		"23",
		"",
	}

	expectedData := [][]string{
		[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		[]string{},
	}

	for index, digits := range testData {
		if res := letterCombinations(digits); !reflect.DeepEqual(res, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], res)
		}
	}
}
