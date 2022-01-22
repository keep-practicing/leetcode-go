package lis

import (
	"os"
	"testing"
)

func TestLIS(t *testing.T) {
	testData := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{},
	}
	expectedData := []int{4, 0}

	for index, data := range testData {
		if res := lengthOfLIS(data); res != expectedData[index] {
			t.Errorf("expected %d, got %d", expectedData[index], res)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
