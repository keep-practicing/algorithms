package intersectionof2arrays

import "algorithms/utils"

func intersection(nums1 []int, nums2 []int) []int {
	set1 := utils.NewSet()
	for _, num := range nums1 {
		set1.Add(num)
	}
	set2 := utils.NewSet()
	for _, num := range nums2 {
		set2.Add(num)
	}

	var res []int
	for item := range set1 {
		if set2.Contains(item) {
			if value, ok := item.(int); ok {
				res = append(res, value)
			}
		}
	}
	return res
}
