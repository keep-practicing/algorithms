package uniquepaths

// recursion
/*
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m, n-1) + uniquePaths(m-1, n)
}
*/

// memory search
/*
func uniquePaths(m int, n int) int {
	memo := make([][]int, m) // memo[i][j]存储(i+1)*(j+1)grid的路径数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	return calcPaths(m, n, memo)
}

func calcPaths(m int, n int, memo [][]int) int {
	if memo[m-1][n-1] != -1 {
		return memo[m-1][n-1]
	}

	if m == 1 || n == 1 {
		memo[m-1][n-1] = 1
		return 1
	}

	res := calcPaths(m, n-1, memo) + calcPaths(m-1, n, memo)
	memo[m-1][n-1] = res
	return res
}
*/

func uniquePaths(m int, n int) int {
	memo := make([][]int, m) // memo[i][j]存储(i+1)*(j+1)grid的路径数量
	for i := 0; i < m; i++ {
		memo[i] = append(memo[i], 1)
	}

	for i := 1; i < n; i++ {
		memo[0] = append(memo[0], 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i] = append(memo[i], memo[i-1][j]+memo[i][j-1])
		}
	}
	return memo[m-1][n-1]
}
