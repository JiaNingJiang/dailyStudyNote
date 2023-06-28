package heapSort

// 传统的方法是：遍历原始数组，将每个元素通过heapInsert插入到大(小)根堆数组中，时间复杂度为 O(N * logN)
// 优化后的方法：从原始数组最后一个结点开始，将该节点当做子树的根节点进行heapify。持续到遍历完整个原始数组，时间复杂度为O(N)
func FastNewHeap(src []int, form bool) {
	for i := len(src) - 1; i >= 0; i-- { // O(N)
		// 此处的heapify每次的时间复杂度不一样(是递增的)，并不是固定的O(logN) （只有src[0]的heapify时间复杂度是O(logN)），
		// 因此这里heapify的时间复杂度可以看做是常数项
		heapify(src, i, len(src), form) // O(1)
	}
}
