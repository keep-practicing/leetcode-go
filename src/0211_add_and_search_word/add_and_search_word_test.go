package aasw

import "testing"

func TestAddAndSearchWord(t *testing.T) {
	obj := Constructor()

	for _, word := range []string{"bad", "dad", "mad"} {
		obj.AddWord(word)
	}

	for i, j := range map[string]bool{"pad": false, "bad": true, ".ad": true, "b..": true, "b..d": false} {
		if res := obj.Search(i); res != j {
			t.Errorf("expected %t, got %t", j, res)
		}
	}
}
