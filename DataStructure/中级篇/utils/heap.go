package utils

// 将一个数组转化为大(小)根堆，(form == true (大根堆) form == false (小根堆) )
func NewHeap(src []interface{}, form bool, less func(interface{}, interface{}) bool) []interface{} {
	arr := make([]interface{}, len(src))
	for _, node := range src {
		if node == nil {
			continue
		}
		HeapInsert(arr, node, form, less)
	}
	return arr
}

func HeapInsert(arr []interface{}, node interface{}, form bool, less func(interface{}, interface{}) bool) ([]interface{}, int) {
	if len(arr) == 0 { // 设置大(小)根堆的根结点
		arr = append(arr, node)
		return arr, 0
	}
	arr = append(arr, node)
	// 构建大(小)根堆
	currentIndex := len(arr) - 1
	parentIndex := (currentIndex - 1) / 2

	curLoc := currentIndex // 记录节点被插入到的位置
	for {
		if currentIndex == 0 {
			break
		}
		if form { // 大根堆
			if less(arr[parentIndex], arr[currentIndex]) {
				swap(&arr[currentIndex], &arr[parentIndex])
				curLoc = parentIndex
			} else {
				break // 无法继续上移了
			}
		} else { // 小根堆
			if less(arr[currentIndex], arr[parentIndex]) {
				swap(&arr[currentIndex], &arr[parentIndex])
				curLoc = parentIndex
			} else {
				break // 无法继续上移了
			}
		}

		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2 // -1/2 = 0 , 因此不能用 parentIndex < 0 作为退出条件
	}
	return arr, curLoc
}

func Heapify(heap []interface{}, start, end int, form bool, less func(interface{}, interface{}) bool) {
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
				max = getMax(heap[leftChildIndex], heap[rightChildIndex], less)
				if max == heap[rightChildIndex] {
					maxIndex = rightChildIndex
				}
			}
			// 如果较大节点比根节点还要大，则交换两者
			if less(heap[currentIndex], max) {
				swap(&heap[maxIndex], &heap[currentIndex])
			}
			newRootIndex = maxIndex
		} else { // 小根堆
			// 获得当前节点 左右孩子 中较小节点的下标
			minIndex := leftChildIndex
			min := heap[leftChildIndex]
			if rightChildIndex <= len(heap)-1 { // 可能只有左孩子，没有右孩子
				min = getMin(heap[leftChildIndex], heap[rightChildIndex], less)
				if min == heap[rightChildIndex] {
					minIndex = rightChildIndex
				}
			}
			// 如果较小节点比根节点还要小，则交换两者
			if less(min, heap[currentIndex]) {
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

func PopAndheapify(heap []interface{}, heapIndex int, form bool, less func(interface{}, interface{}) bool) interface{} {

	num := heap[0]            // 每次总是返回根堆的根节点
	heap[0] = heap[heapIndex] // 让末尾的叶子结点替换掉根节点

	// 将新的根节点下沉到合适的位置
	Heapify(heap, 0, heapIndex-1, form, less) // 注意：end必须是heapIndex - 1，作用是相当于heap[heapIndex]被删除
	return num
}

func getMax(a, b interface{}, less func(interface{}, interface{}) bool) interface{} {
	if less(a, b) {
		return b
	} else {
		return a
	}
}

func getMin(a, b interface{}, less func(interface{}, interface{}) bool) interface{} {
	if less(a, b) {
		return a
	} else {
		return b
	}
}

func swap(a, b *interface{}) {
	temp := *a
	*a = *b
	*b = temp
}

func GetSortArrFromHeap(heap []interface{}, form bool, less func(interface{}, interface{}) bool) {
	arr := make([]interface{}, 0, len(heap))
	heapIndex := len(heap) - 1

	for i := heapIndex; i >= 0; i-- {
		arr = append(arr, PopAndheapify(heap, i, form, less))
	}

	for i, v := range arr {
		heap[i] = v
	}
}
