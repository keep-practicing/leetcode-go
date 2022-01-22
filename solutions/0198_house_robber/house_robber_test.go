package houserobber

import "testing"

func TestHouseRobber(t *testing.T) {
	testData := [][]int{
		{1, 2, 3, 1},
		{9, 1, 293, 5986, 81, 384},
		{1, 2},
		{},
	}
	expectedData := []int{4, 6379, 2, 0}

	for index, data := range testData {
		if res := rob(data); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}
