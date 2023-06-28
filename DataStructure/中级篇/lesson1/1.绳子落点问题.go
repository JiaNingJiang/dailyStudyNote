package lesson1

import "math"

// arr[]必须是有序数组
func ScopeIncluedSpot(arr []int, L int) int {
	if len(arr) == 0 {
		return 0
	}

	maxCover := math.MinInt
	for i := 0; i < len(arr); i++ {
		end := arr[i]    // arr[i]作为终点
		start := end - L // 根据绳子长度算出起点

		_, index := FindRightNearest(arr, start)

		coverPoint := i - index + 1
		maxCover = int(math.Max(float64(maxCover), float64(coverPoint)))
	}
	return maxCover
}

// 在有序数组中找出大于等于num的最小的数字以及其下标
func FindRightNearest(arr []int, num int) (int, int) {

	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	res := math.MaxInt
	index := -1

	for {
		if left > right {
			break
		}
		if num == arr[mid] { // arr[mid]刚好等于num，那么就返回num
			res = num
			index = mid
			break
		} else if num > arr[mid] { // num > arr[mid],则找大于num的要到右区域找
			left = mid + 1
			mid = (left + right) / 2
		} else { // num < arr[mid],则还需要检测左侧是否有比arr[mid]更小的，大于等于num的数
			res = int(math.Min(float64(res), float64(arr[mid])))
			index = mid
			right = mid - 1
			mid = (left + right) / 2
		}
	}
	return res, index
}
