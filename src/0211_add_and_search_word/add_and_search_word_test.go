package aasw

import "testing"

func TestAddAndSearchWord(t *testing.T) {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	for i, j := range map[string]bool{"pad": false, "bad": true, ".ad": true, "b..": true, "b..d": false} {
		if res := obj.Search(i); res != j {
			t.Errorf("expected %t, got %t", j, res)
		}
	}
}
