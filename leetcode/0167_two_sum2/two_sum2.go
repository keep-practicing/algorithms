package twosum2

func twoSum2(numbers []int, target int) []int {
	n := len(numbers)
	var (
		l = 0
		r = n - 1
	)

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	panic("The input has no solution.")
}
