package utils

import "testing"

func TestCalcMinInt(t *testing.T) {
	testData := [][]int{
		{3, 4, 67, 8},
		{},
	}
	expectedData := []int{3, 0}

	for index, data := range testData {
		if res := CalcMinInt(data...); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
