package heapSort

func BigRootHeapSort(src []int) {
	arr := NewHeap(src, true)     // 先根据原始数组构建大根堆
	GetSortArrFromHeap(arr, true) // 挨个弹出最大元素并随时堆化

	for i, v := range arr {
		src[i] = v
	}
}

func SmallRootHeapSort(src []int) {
	arr := NewHeap(src, false)     // 先根据原始数组构建小根堆
	GetSortArrFromHeap(arr, false) // 挨个弹出最小元素并随时堆化

	for i, v := range arr {
		src[i] = v
	}
}
