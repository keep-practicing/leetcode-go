package nonoverlappingintervals

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	testData := [][]Interval{
		{
			Interval{Start: 1, End: 2},
			Interval{Start: 1, End: 2},
			Interval{Start: 1, End: 2},
		},
		{
			Interval{Start: 1, End: 2},
			Interval{Start: 2, End: 3},
		},
		{
			Interval{Start: 1, End: 2},
			Interval{Start: 2, End: 3},
			Interval{Start: 3, End: 4},
			Interval{Start: 1, End: 3},
		},
	}
	expectedData := []int{2, 0, 1}

	testFuncs := []func([]Interval) int{
		eraseOverlapIntervalsGreedy,
		eraseOverlapIntervalsDp,
	}

	for index, data := range testData {
		for _, testFunc := range testFuncs {
			if res := testFunc(data); res != expectedData[index] {
				t.Errorf("expected %d, got %d", expectedData[index], res)
			}
		}
	}
}
