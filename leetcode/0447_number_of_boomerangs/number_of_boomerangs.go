package numberofboomerangs

import (
	"math"
)

func numberOfBoomerangs(points [][]int) int {
	var (
		res int
		n   = len(points)
	)

	for i := 0; i < n; i++ {
		record := make(map[int]int)
		for j := 0; j < n; j++ {
			if j != i {
				dis := dis(points[i], points[j])
				record[dis]++
			}
		}

		for _, j := range record {
			res += j * (j - 1)
		}
	}
	return res
}

func dis(point1 []int, point2 []int) int {
	return int(math.Pow(float64(point1[0]-point2[0]), float64(2)) + math.Pow(float64(point1[1]-point2[1]), float64(2)))
}
