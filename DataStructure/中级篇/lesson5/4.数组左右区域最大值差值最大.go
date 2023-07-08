package lesson5

import "math"

func LeftAndRightMaxDiff(arr []int) int {
	// 1.获取整个数组的最大值
	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	leftDiff := max - arr[len(arr)-1] // 将max划分到左侧区域
	rightDiff := max - arr[0]         // 将max划分到右侧区域

	return getMax(leftDiff, rightDiff)
}
