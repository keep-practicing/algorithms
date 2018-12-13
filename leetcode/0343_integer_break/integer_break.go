package integerbreak

import "op/utils"

// recursion
/*
func integerBreak(n int) int {
	if n == 1 {
		return 1
	}

	res := utils.MinInt

	for i := 1; i < n; i++ {
		res, _ = utils.CalcMaxInt(res, i*(n-i), i*integerBreak(n-i))
	}
	return res
}
*/

// memory search
/*
func integerBreak(n int) int {
	memo := make([]int, n+1)

	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	return breakInt(n, memo)
}

func breakInt(n int, memo []int) int {
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	res := utils.MinInt
	for i := 1; i < n; i++ {
		res, _ = utils.CalcMaxInt(res, i*(n-i), i*breakInt(n-i, memo))
	}
	memo[n] = res
	return res
}
*/

// dynamic programming
func integerBreak(n int) int {
	// memo[i]表示将数字i分割（至少分割成两部分）后得到的最大乘积.
	memo := make([]int, n+1)

	memo[1] = 1

	for i := 2; i <= n; i++ {
		// 求解memo[i]
		for j := 1; j <= i-1; j++ {
			memo[i] = utils.CalcMaxInt(memo[i], j*(i-j), j*memo[i-j])
		}
	}
	return memo[n]
}
