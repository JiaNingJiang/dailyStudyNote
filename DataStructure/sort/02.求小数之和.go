package sort

func SmallSum(arr []int) int {
	left := 0
	right := len(arr) - 1
	return process1(arr, left, right)
}

func process1(arr []int, left, right int) int {
	if left == right { // 分区只有一个数，对于小数和没有贡献
		return 0
	}
	mid := left + ((right - left) >> 1)
	leftSmallSum := process1(arr, left, mid)       // 求出左分区的小数和，同时完成左分区的排序
	rightSmallSum := process1(arr, mid+1, right)   // 求出右分区的小数和,同时完成右分区的排序
	mergeSmallSum := merge1(arr, left, right, mid) // 完成左右分区的归并，并求出归并带来的小数和

	return leftSmallSum + rightSmallSum + mergeSmallSum
}

func merge1(arr []int, left, right, mid int) int {
	leftIndex := left
	rightIndex := mid + 1
	temp := make([]int, 0, len(arr))
	smallSum := 0

	for {
		if leftIndex > mid || rightIndex > right {
			break
		}
		if arr[leftIndex] < arr[rightIndex] {
			temp = append(temp, arr[leftIndex])
			smallSum += arr[leftIndex] * (right - rightIndex + 1) // 当前数将成为右边所有大数的小数(与原始归并排序不一样的地方)
			leftIndex++
		} else {
			temp = append(temp, arr[rightIndex])
			rightIndex++
		}
	}

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

	for i, v := range temp {
		arr[left+i] = v
	}

	return smallSum
}
