package mergeSort

func MergeSort(arr []int) {
	left := 0
	right := len(arr) - 1

	process(arr, left, right)
}

func process(arr []int, left, right int) {
	if left == right { // 当前分区只有一个元素
		return
	}
	mid := left + ((right - left) >> 1)
	process(arr, left, mid)      // 左半区域进行排序
	process(arr, mid+1, right)   // 对右半区域进行排序
	merge(arr, left, right, mid) // 进行归并
}

func merge(arr []int, left, right, mid int) {

	leftIndex := left     //左半区域扫描指针
	rightIndex := mid + 1 // 右半区域扫描指针

	temp := make([]int, 0)

	// 第一个循环，直到将左右分区中的一个读完
	for {
		if leftIndex > mid || rightIndex > right {
			break
		}
		if arr[leftIndex] < arr[rightIndex] {
			temp = append(temp, arr[leftIndex])
			leftIndex++
		} else {
			temp = append(temp, arr[rightIndex])
			rightIndex++
		}
	}
	//  第二个循环,负责读完剩下的分区
	for {
		if leftIndex > mid {
			break
		}
		temp = append(temp, arr[leftIndex])
		leftIndex++
	}

	for {
		if rightIndex > right {
			break
		}
		temp = append(temp, arr[rightIndex])
		rightIndex++
	}
	// 第三个循环,将完成排序的本区域全部拷贝到原数组对应区域
	for i, v := range temp {
		arr[left+i] = v
	}
}
