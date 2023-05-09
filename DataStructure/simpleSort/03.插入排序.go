package simpleSort

func InsertSort(arr []int) {

	temp := append([]int{0}, arr[:]...) // 在待排序数组的最前方添加一个辅助空间

	for i := 1; i < len(temp)-1; i++ { // 辅助空间不需要进行排序
		if temp[i] > temp[i+1] { // 前一个比后一个要大，那么需要进行交换
			temp[0] = temp[i+1] // 辅助空间存储较小的那一个数，作为哨兵
			var j int = 0
			for j = i; temp[j] > temp[0]; j-- { // 从arr[i]往前,只要比哨兵大,就往后移动一位；直到比哨兵小才停止
				temp[j+1] = temp[j]
			}
			temp[j+1] = temp[0] //将哨兵再次插入到数组中，此时对于哨兵来说：左边必然都是比自己小的，右边必然都是比自己大的
		}

	}
	for i := 1; i < len(temp); i++ {
		arr[i-1] = temp[i]
	}
	// TODO:问题，直接对arr进行扩展，再缩减。
}
