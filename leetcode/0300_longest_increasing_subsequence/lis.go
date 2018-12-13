package lis

import "op/utils"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				memo[i] = utils.CalcMaxInt(memo[i], 1+memo[j])
			}
		}
	}
	return utils.CalcMaxInt(memo...)
}
