package heapSort

// 修改一个大(小)根堆指定节点的值，最后再恢复堆结构
func RewriteAndRecovery(arr []int, index int, value int, form bool) {

	if index < 0 || index >= len(arr) { // 不允许越界修改
		return
	}
	oldValue := arr[index]
	arr[index] = value

	if form { // 大根堆
		if arr[index] > oldValue { // 修改后的值大于原节点的值，需要将当前节点上移
			currentIndex := index
			parentIndex := (currentIndex - 1) / 2
			form := true
			moveUP(arr, currentIndex, parentIndex, form)
		} else if arr[index] < oldValue { // 修改后的值小于原节点的值，需要将当前节点下移
			currentIndex := index
			leftChildIndex := currentIndex*2 + 1
			rightChildIndex := currentIndex*2 + 2
			form := true
			moveDown(arr, currentIndex, leftChildIndex, rightChildIndex, form)
		} else { // 值没变,无需恢复
			return
		}
	} else { // 小根堆
		if arr[index] > oldValue { // 变大了,下移
			currentIndex := index
			leftChildIndex := currentIndex*2 + 1
			rightChildIndex := currentIndex*2 + 2
			form := false
			moveDown(arr, currentIndex, leftChildIndex, rightChildIndex, form)
		} else if arr[index] < oldValue { // 变小了,上移
			currentIndex := index
			parentIndex := (currentIndex - 1) / 2
			form := false
			moveUP(arr, currentIndex, parentIndex, form)
		} else {
			return
		}
	}
}

func moveUP(arr []int, currentIndex, parentIndex int, form bool) {
	for {
		if currentIndex == 0 { // 已到边界
			break
		}
		if form { // 大根堆
			if arr[currentIndex] > arr[parentIndex] {
				swap(&arr[currentIndex], &arr[parentIndex])
			} else {
				break
			}
		} else { // 小根堆
			if arr[currentIndex] < arr[parentIndex] {
				swap(&arr[currentIndex], &arr[parentIndex])
			} else {
				break
			}
		}
		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2
	}
}

func moveDown(arr []int, currentIndex, leftChildIndex, rightChildIndex int, form bool) {
	for {
		if leftChildIndex > len(arr)-1 { // 没有任何孩子节点
			break
		}
		if form { // 大根堆, 因为比最大的孩子节点小才下移(只能跟最大的孩子节点交换)
			maxIndex := leftChildIndex
			max := arr[leftChildIndex]
			if rightChildIndex <= len(arr)-1 { // 可能只有左孩子，没有右孩子
				max = getMax(arr[leftChildIndex], arr[rightChildIndex])
				if max == arr[rightChildIndex] {
					maxIndex = rightChildIndex
				}
			}

			if arr[currentIndex] < max { // 与最大的孩子节点进行交换
				swap(&arr[currentIndex], &arr[maxIndex])
			} else { // 比两个孩子都大，则不必继续往下交换
				break
			}
			currentIndex = maxIndex
		} else { // 小根堆，因为比最小的孩子节点大才下移(只能跟最小的孩子节点交换)
			minIndex := leftChildIndex
			min := arr[leftChildIndex]
			if rightChildIndex <= len(arr)-1 { // 可能只有左孩子，没有右孩子
				min = getMin(arr[leftChildIndex], arr[rightChildIndex])
				if min == arr[rightChildIndex] {
					minIndex = rightChildIndex
				}
			}

			if arr[currentIndex] > min { // 与最小的孩子节点进行交换
				swap(&arr[currentIndex], &arr[minIndex])
			} else { // 比两个孩子都小，则不必继续往下交换
				break
			}
			currentIndex = minIndex
		}
		leftChildIndex = currentIndex*2 + 1
		rightChildIndex = currentIndex*2 + 2
	}
}
