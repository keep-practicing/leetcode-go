package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []string{
		"hello",
		"A man, a plan, a canal: Panama",
		"",
		"ab",
	}
	expected := []string{
		"olleh",
		"amanaP :lanac a ,nalp a ,nam A",
		"",
		"ba",
	}

	for index, data := range testCases {
		if res := reverseString(data); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}

func TestReverseString1(t *testing.T) {
	testCases := []string{
		"hello",
		"A man, a plan, a canal: Panama",
		"",
		"ab",
	}
	expected := []string{
		"olleh",
		"amanaP :lanac a ,nalp a ,nam A",
		"",
		"ba",
	}

	for index, data := range testCases {
		if res := reverseString1(data); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
