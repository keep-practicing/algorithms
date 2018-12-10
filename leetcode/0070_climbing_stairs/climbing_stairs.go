package climbingstairs

// recursion
/*
func climbStairs(n int) int {
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	//

	if n == 1 || n == 0 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
*/

// memory search
/*
func climbStairs(n int) int {
	var memo []int
	for i := 0; i <= n; i++ {
		memo = append(memo, -1)
	}
	return calcWays(n, memo)
}

func calcWays(n int, memo []int) int {
	if n == 0 || n == 1 {
		return 1
	}

	if memo[n] == -1 {
		memo[n] = calcWays(n-1, memo) + calcWays(n-2, memo)
	}

	return memo[n]
}
*/

// dynamic programming
func climbStairs(n int) int {
	var memo = []int{1, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n]
}
