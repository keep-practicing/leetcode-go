package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	digits := []int{9, 9, 9, 9}
	expected := []int{1, 0, 0, 0, 0}
	if res := plusOne(digits); !reflect.DeepEqual(expected, res) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
