package lenoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	s := "hello world "
	expected := 5
	if res := lengthOfLastWord(s); res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
