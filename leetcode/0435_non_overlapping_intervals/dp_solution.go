package nonoverlappingintervals

import (
	"algorithms/utils"
	"sort"
)

func eraseOverlapIntervalsDp(intervals []Interval) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Sort(IntervalSliceCompareWithStart(intervals))

	// memo[i]表示使用intervals[0...i]的区间能构成的最长不重叠区间序列
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}

	for i := 1; i < n; i++ {
		// memo[i]
		for j := 0; j < i; j++ {
			if intervals[i].Start >= intervals[j].End {
				memo[i] = utils.CalcMaxInt(memo[i], 1+memo[j])
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		res = utils.CalcMaxInt(res, memo[i])
	}
	return n - res
}
