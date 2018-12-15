package movezeroes

func moveZeroes2(nums []int) {
	var k int
	for i := range nums {
		if nums[i] != 0 {
			if k != i {
				nums[k] = nums[i]
			}
			k++
		}
	}

	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
