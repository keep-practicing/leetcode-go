package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes2(t *testing.T) {
	testData := [][]int{
		{0, 1, 0, 3, 12},
	}

	expectedData := [][]int{
		{1, 3, 12, 0, 0},
	}

	for index, data := range testData {
		if moveZeroes2(data); !reflect.DeepEqual(data, expectedData[index]) {
			t.Errorf("expected %v, got %v", expectedData[index], data)
		}
	}
}
