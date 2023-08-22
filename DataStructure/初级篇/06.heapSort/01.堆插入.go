package heapSort

// 将一个数组转化为大(小)根堆，(form == true (大根堆) form == false (小根堆) )
func NewHeap(src []int, form bool) []int {
	arr := make([]int, len(src))
	for heapSize, num := range src {
		heapInsert(arr, num, heapSize, form)
	}
	return arr
}

func heapInsert(arr []int, num int, heapSize int, form bool) {
	if heapSize == 0 { // 设置大(小)根堆的根结点
		arr[0] = num
		return
	}
	arr[heapSize] = num
	// 构建大(小)根堆
	currentIndex := heapSize
	parentIndex := (heapSize - 1) / 2
	for {
		if currentIndex == 0 {
			break
		}
		if form { // 大根堆
			if arr[currentIndex] > arr[parentIndex] {
				swap(&arr[currentIndex], &arr[parentIndex])
			} else {
				break // 无法继续上移了
			}
		} else { // 小根堆
			if arr[currentIndex] < arr[parentIndex] {
				swap(&arr[currentIndex], &arr[parentIndex])
			} else {
				break // 无法继续上移了
			}
		}

		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2 // -1/2 = 0 , 因此不能用 parentIndex < 0 作为退出条件
	}

}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
