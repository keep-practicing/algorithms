package utils

// CalcMaxInt calc max int from multi nums.
func CalcMaxInt(nums ...int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}

	res := MinInt

	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res, true
}
