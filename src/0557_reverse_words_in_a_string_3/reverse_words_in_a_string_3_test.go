package rwias3

import "testing"

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"

	if res := reverseWords(s); res != expected {
		t.Errorf("expected %s, got %s", expected, res)
	}
}
