package twosum

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for i, j := range nums {
		complement := target - j
		if res, ok := record[complement]; ok {
			return []int{res, i}
		}
		record[j] = i
	}
	return []int{}
}
