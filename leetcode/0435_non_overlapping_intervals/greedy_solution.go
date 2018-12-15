package nonoverlappingintervals

import "sort"

func eraseOverlapIntervalsGreedy(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Sort(IntervalSliceCompareWithEnd(intervals))
	var (
		res = 1
		pre = 0
	)

	for i := 1; i < n; i++ {
		if intervals[i].Start >= intervals[pre].End {
			res++
			pre = i
		}
	}
	return n - res
}
