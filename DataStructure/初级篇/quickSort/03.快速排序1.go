package quickSort

func QuickSort1(arr []int) {
	left := 0
	right := len(arr) - 1
	sort1(arr, left, right)
}

func sort1(arr []int, left, right int) {
	if left >= right { // 区域最多只有一个元素，无需进行排序
		return
	}

	leftIndex := left
	rightIndex := right - 1
	target := arr[right]
	i := left // i从区域的最左侧开始

	// 先将整个数组分割为两个区域：1.小于等于arr[right]的  2.大于arr[right]的
	for {
		if i > rightIndex {
			swap(&arr[right], &arr[rightIndex+1]) // 分割完成，将arr[right]与大于区域的第一个数进行交换(扩大小于等于区域)
			break
		}
		if arr[i] <= target {
			swap(&arr[i], &arr[leftIndex])
			i++
			leftIndex++
		} else {
			swap(&arr[i], &arr[rightIndex])
			//i++
			rightIndex--
		}
	}
	sort1(arr, left, leftIndex-1)   // 对小于等于区域进行递归分区
	sort1(arr, rightIndex+2, right) // 对大于区域进行递归分区

}
