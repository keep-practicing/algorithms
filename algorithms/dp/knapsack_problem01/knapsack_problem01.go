package knapsackproblem01

func knapsack01(w []int, v []int, c int) int {
	if len(w) != len(v) {
		panic("重量列表和价值列表数量不符！")
	}

	n := len(w)

	if n == 0 {
		return 0
	}

	memo := make([]int, c+1)
	for i := 0; i < c+1; i++ {
		memo[i] = -1
	}

	// 只把第0个物品放入背包，动态规划二维数组值的变化。循环变量i表示背包的容量。
	for i := 0; i <= c; i++ {
		if i >= w[0] {
			memo[i] = v[0]
		} else {
			memo[i] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := c; j >= w[i]; j-- {
			temp := v[i] + memo[j-w[i]]
			if temp > memo[j] {
				memo[j] = temp
			}
		}
	}
	return memo[c]
}
