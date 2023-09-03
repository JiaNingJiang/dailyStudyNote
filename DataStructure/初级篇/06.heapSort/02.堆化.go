package heapSort

// form == true表示heap为大根堆 ， form == false表示heap为小根堆
func GetSortArrFromHeap(heap []int, form bool) {
	arr := make([]int, 0, len(heap))
	heapIndex := len(heap) - 1

	for i := heapIndex; i >= 0; i-- {
		arr = append(arr, popAndheapify(heap, i, form))
	}

	for i, v := range arr {
		heap[i] = v
	}
}

func popAndheapify(heap []int, heapIndex int, form bool) int {

	num := heap[0]            // 每次总是返回根堆的根节点
	heap[0] = heap[heapIndex] // 让末尾的叶子结点替换掉根节点

	// 将新的根节点下沉到合适的位置
	heapify(heap, 0, heapIndex-1, form) // 注意：end必须是heapIndex - 1，作用是相当于heap[heapIndex]被删除
	return num
}

func heapify(heap []int, start, end int, form bool) {
	currentIndex := start // 对以 heap[start]为根节点的子树进行heapify
	leftChildIndex := currentIndex*2 + 1
	rightChildIndex := currentIndex*2 + 2

	for {
		if leftChildIndex > end { // 没有任何孩子节点(最多到heap[end-1])
			break
		}
		newRootIndex := currentIndex
		if form { // 大根堆
			// 获得当前节点 左右孩子 中较大节点的下标
			maxIndex := leftChildIndex
			max := heap[leftChildIndex]
			if rightChildIndex <= len(heap)-1 { // 可能只有左孩子，没有右孩子
				max = getMax(heap[leftChildIndex], heap[rightChildIndex])
				if max == heap[rightChildIndex] {
					maxIndex = rightChildIndex
				}
			}
			// 如果较大节点比根节点还要大，则交换两者
			if max > heap[currentIndex] {
				swap(&heap[maxIndex], &heap[currentIndex])
			}
			newRootIndex = maxIndex
		} else { // 小根堆
			// 获得当前节点 左右孩子 中较小节点的下标
			minIndex := leftChildIndex
			min := heap[leftChildIndex]
			if rightChildIndex <= len(heap)-1 { // 可能只有左孩子，没有右孩子
				min = getMin(heap[leftChildIndex], heap[rightChildIndex])
				if min == heap[rightChildIndex] {
					minIndex = rightChildIndex
				}
			}
			// 如果较小节点比根节点还要小，则交换两者
			if min < heap[currentIndex] {
				swap(&heap[minIndex], &heap[currentIndex])
			}
			newRootIndex = minIndex
		}
		// 更新循环变量
		currentIndex = newRootIndex
		leftChildIndex = currentIndex*2 + 1
		rightChildIndex = currentIndex*2 + 2
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
