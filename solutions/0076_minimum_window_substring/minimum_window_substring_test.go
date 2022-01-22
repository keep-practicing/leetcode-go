package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	l := "ABC"
	expected := "BANC"

	if res := minWindow(s, l); res != expected {
		t.Errorf("expected %s, got %s", expected, res)
	}

}
