package selfdividingnumbers

import (
	"reflect"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}
	if res := selfDividingNumbers(1, 22); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
