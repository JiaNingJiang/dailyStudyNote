package simpleSort

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i                       // 用来记录当前轮最小值的下标，默认设置为i
		for j := i + 1; j < len(arr); j++ { // 找出从 i~len(arr)-1 (包括i在内)的最小值的下标
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i { // 交换本轮最小值到i位置
			swap(&arr[i], &arr[minIndex])
		}
	}
}
