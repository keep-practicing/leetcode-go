package impltree

import "testing"

func TestImplTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	for i, j := range map[string]bool{"apple": true, "app": false, "hello": false} {
		if res := trie.Search(i); res != j {
			t.Errorf("expected %t, got %t", j, res)
		}
	}

	for i, j := range map[string]bool{"app": true, "as": false} {
		if res := trie.StartsWith(i); res != j {
			t.Errorf("expected %t, got %t", j, res)
		}
	}

	trie.Insert("app")
	if res := trie.Search("app"); res != true {
		t.Errorf("expected %t, got %t", true, res)
	}
}
