package movezeroes

func moveZeroes(nums []int) {
	var k int
	for i := range nums {
		if nums[i] != 0 {
			if k != i {
				nums[i], nums[k] = nums[k], nums[i]
			}
			k++
		}
	}
}
