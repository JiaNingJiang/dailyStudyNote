package quickSort

func QuickSort2(arr []int) {
	left := 0
	right := len(arr) - 1
	sort1(arr, left, right)
}

func sort2(arr []int, left, right int) {
	if left >= right { // 区域最多只有一个元素，无需进行排序
		return
	}

	leftIndex := left
	rightIndex := right - 1
	target := arr[right]
	i := left

	// 先将整个数组分割为三个区域：1.小于arr[right]的 2.等于arr[right]的  3.大于arr[right]的
	for {
		if i > rightIndex {
			swap(&arr[right], &arr[rightIndex+1]) // 分割完成，将arr[right]与大于区域的第一个数进行交换(扩大等于区域)
			break
		}
		if arr[i] < target {
			swap(&arr[i], &arr[leftIndex])
			i++
			leftIndex++
		} else if arr[i] == target {
			i++
		} else {
			swap(&arr[i], &arr[rightIndex])
			//i++
			rightIndex--
		}
	}
	sort2(arr, left, leftIndex-1)   // 对小于区域进行递归分区
	sort2(arr, rightIndex+2, right) // 对大于区域进行递归分区

}
