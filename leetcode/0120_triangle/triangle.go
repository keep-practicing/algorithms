package triangle

import "op/utils"

func mininumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	result := utils.MaxInt
	dfs(0, 0, 0, triangle, &result)
	return result
}

func dfs(x int, y int, sum int, triangle [][]int, result *int) {
	if len(triangle) == x {
		if sum < *result {
			*result = sum
		}
		return
	}
	sum += triangle[x][y]

	dfs(x+1, y, sum, triangle, result)
	dfs(x+1, y+1, sum, triangle, result)
}
