package partitionequalsubsetsum

import "testing"

func TestPartitionEqualSubsetSum(t *testing.T) {
	testData := [][]int{
		{1, 5, 11, 5},
		{1, 2, 3, 5},
	}

	expectedData := []bool{true, false}

	for index, data := range testData {
		if res := canPartition(data); res != expectedData[index] {
			t.Errorf("expected %t, got %t", expectedData[index], res)
		}
	}
}
