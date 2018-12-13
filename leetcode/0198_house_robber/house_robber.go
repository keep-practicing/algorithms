package houserobber

import "op/utils"

/*
func rob(nums []int) int {
	return tryRob(nums, 0)
}

func tryRob(nums []int, index int) int {
	if index >= len(nums) {
		return 0
	}

	var res int
	for i := index; i < len(nums); i++ {
		res, _ = utils.CalcMaxInt(res, nums[i]+tryRob(nums, i+2))
	}
	return res
}
*/

// memory search
/*
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}
	return tryRob(nums, 0, memo)
}

func tryRob(nums []int, index int, memo []int) int {
	if index >= len(nums) {
		return 0
	}

	if memo[index] != -1 {
		return memo[index]
	}

	var res int
	for i := index; i < len(nums); i++ {
		res, _ = utils.CalcMaxInt(res, nums[i]+tryRob(nums, i+2, memo))
	}
	memo[index] = res
	return res
}*/

// dynamic programming
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	memo := make([]int, n)

	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	memo[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			if j+2 < n {
				memo[i] = utils.CalcMaxInt(memo[i], nums[j]+memo[j+2])
			} else {
				memo[i] = utils.CalcMaxInt(memo[i], nums[j])
			}
		}
	}
	return memo[0]
}
