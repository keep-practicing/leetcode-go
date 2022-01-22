package topkfe

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}
	if res := topKFrequent(nums, k); !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
