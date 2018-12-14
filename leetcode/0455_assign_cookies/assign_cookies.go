package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(g)))

	var (
		si, gi, res int
	)

	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			res++
			gi++
			si++
		} else {
			gi++
		}
	}
	return res
}
