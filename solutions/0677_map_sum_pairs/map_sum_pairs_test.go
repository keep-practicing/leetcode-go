package mapsumpairs

import "testing"

func TestMapSumPairs(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple", 3)
	if res := obj.Sum("ap"); res != 3 {
		t.Errorf("expected %d, got %d", 3, res)
	}

	obj.Insert("app", 2)
	if res := obj.Sum("ap"); res != 5 {
		t.Errorf("expected %d, got %d", 5, res)
	}

	if res := obj.Sum("al"); res != 0 {
		t.Errorf("expected %d, got %d", 0, res)
	}
}
