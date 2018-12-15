package minimumsizesubarraysum

func minSubArrayLen(s int, nums []int) int {
	var (
		l   = 0
		r   = -1
		res = len(nums) + 1
		sum = 0
	)

	for l < len(nums) {
		if sum < s && r+1 < len(nums) {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s && res > r-l+1 {
			res = r - l + 1
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}
