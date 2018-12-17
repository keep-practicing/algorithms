package foursum2

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var (
		record = make(map[int]int)
		res    int
	)

	for _, i := range A {
		for _, j := range B {
			record[i+j]++
		}
	}

	for _, i := range C {
		for _, j := range D {
			if s, ok := record[-i-j]; ok && s > 0 {
				res += s
			}
		}
	}
	return res
}
