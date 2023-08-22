package heapSort

func Process(arr []int, k int) {
	sortArray := make([]int, 0, len(arr)) // 存储从小到大拍完序的数组
	preKelement := make([]int, 0, k+1)
	for i := 0; i <= k; i++ {
		preKelement = append(preKelement, arr[i]) // 存储基本有序数组的前 0 ~ k个元素
	}
	FastNewHeap(preKelement, false)     // 将preKelement变为小根堆
	for i := k + 1; i < len(arr); i++ { // 挨个将后 k+1 ~ len(arr) - 1 位置上的元素添加到小根堆,同时每次弹出小根堆的堆顶元素(当前小根堆的最小值)
		sortArray = append(sortArray, preKelement[0]) // 获取当前堆顶元素
		preKelement[0] = arr[i]                       // 插入新元素到堆顶
		heapify(preKelement, 0, k+1, false)           // 重新堆化
	}
	// 最后将小根堆的全部元素挨个弹出
	GetSortArrFromHeap(preKelement, false)
	sortArray = append(sortArray, preKelement...)

	for i, v := range sortArray {
		arr[i] = v
	}
}
