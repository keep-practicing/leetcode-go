package nonoverlappingintervals

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

// IntervalSliceCompareWithStart slice of Interval sort with start
type IntervalSliceCompareWithStart []Interval

//  Len is the number of elements in the collection.
func (i IntervalSliceCompareWithStart) Len() int { return len(i) }

// Less compare Interval with start
func (i IntervalSliceCompareWithStart) Less(j, k int) bool {
	if i[j].Start != i[k].Start {
		return i[j].Start < i[k].Start
	}
	return i[j].End < i[k].End
}

func (i IntervalSliceCompareWithStart) Swap(j, k int) { i[j], i[k] = i[k], i[j] }

// IntervalSliceCompareWithEnd slice of Interval sort with end
type IntervalSliceCompareWithEnd []Interval

// Len is the number of elements in the collection.
func (i IntervalSliceCompareWithEnd) Len() int { return len(i) }

// Less compare Interval with end
func (i IntervalSliceCompareWithEnd) Less(j, k int) bool {
	if i[j].End != i[k].End {
		return i[j].End < i[k].End
	}
	return i[j].Start < i[k].Start
}

func (i IntervalSliceCompareWithEnd) Swap(j, k int) { i[j], i[k] = i[k], i[j] }
