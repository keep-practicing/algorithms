package wigglesubsequence

import "op/utils"

/*
 * 用up[i]和down[i]分别记录到第i个元素为止以上升沿和下降沿结束的最长“摆动”
 * 序列长度，遍历数组，如果nums[i]>nums[i-1]，表明第i-1到第i个元素是上升的，
 * 因此up[i]只需在down[i-1]的基础上加1即可，而down[i]保持down[i-1]不变；
 * 如果nums[i]<nums[i-1]，表明第i-1到第i个元素是下降的，因此down[i]
 * 只需在up[i-1]的基础上加1即可，而up[i]保持up[i-1]不变；如果nums[i]==nums[i-1]，
 * 则up[i]保持up[i-1]，down[i]保持down[i-1]。比较最终以上升沿和下降沿结束的
 * 最长“摆动”序列长度即可获取最终结果
 * */
func wiggleMaxLength(nums []int) int {
	n := len(nums)

	if n == 0 || n == 1 {
		return n
	}

	up := make([]int, n)
	for i := 0; i < n; i++ {
		up[i] = 1
	}
	down := make([]int, n)
	copy(down, up)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				up[i] = utils.CalcMaxInt(up[i], down[j]+1)
			} else if nums[i] < nums[j] {
				down[i] = utils.CalcMaxInt(down[i], up[j]+1)
			}
		}
	}
	return utils.CalcMaxInt(up[n-1], down[n-1])
}
