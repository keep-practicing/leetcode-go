package utils

import (
	"testing"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	if unsafe.Sizeof(Exists) != 0 {
		t.Error("Exists size must be zero.")
	}
}

func TestContains(t *testing.T) {
	set := NewSet(3, 4)
	set.Add(5)
	if set.Contains(3) != true {
		t.Error("should contains 4.")
	}

	if set.Contains(6) != false {
		t.Error("should not contains 6.")
	}
}

func TestSize(t *testing.T) {
	set := NewSet(3, 4)
	set.Add(5)
	if set.Size() != 3 {
		t.Error("size should be 3.")
	}
}

func TestEqual(t *testing.T) {
	set := NewSet(3, 4)
	set.Add(5)

	set1 := NewSet(3, 4, 5)
	if set.Equal(set1) != true {
		t.Error("set should equal with set1.")
	}
	set1.Add(6)
	if set.Equal(set1) == true {
		t.Error("set shouldn't equal with set1.")
	}
}

func TestIsSubset(t *testing.T) {
	set := NewSet(3, 4)
	set.Add(5)

	set1 := NewSet(3, 4, 5, 6)

	if set1.IsSubset(set) == true {
		t.Error("set1 shouldn't be set's subset.")
	}

	if set.IsSubset(set1) == false {
		t.Error("set should be set1's subset.")
	}
}

func TestClear(t *testing.T) {
	set := NewSet(3, 4)
	set.Clear()
	if set.Size() != 0 {
		t.Error("set should be clear.")
	}
}
