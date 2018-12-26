package reversevowels

import "testing"

func TestReverseVowels(t *testing.T) {
	testCases := []string{
		"hello",
		"leetcode",
		"aA",
		"a.b,.",
	}

	expected := []string{
		"holle",
		"leotcede",
		"Aa",
		"a.b,.",
	}

	for index, data := range testCases {
		if res := reverseVowels(data); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}

func TestReverseVowels1(t *testing.T) {
	testCases := []string{
		"hello",
		"leetcode",
		"aA",
		"a.b,.",
	}

	expected := []string{
		"holle",
		"leotcede",
		"Aa",
		"a.b,.",
	}

	for index, data := range testCases {
		if res := reverseVowels1(data); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
