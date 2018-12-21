package sortcolors

import (
	"reflect"
	"testing"
)

func TestSortColorsCountSort(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	expected := []int{0, 0, 1, 1, 2, 2}

	sortColorsCountSort(nums)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}
}
func TestSortColorsQuickSort3Ways(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	expected := []int{0, 0, 1, 1, 2, 2}

	sortColorsQuickSort3Ways(nums)

	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}
}
