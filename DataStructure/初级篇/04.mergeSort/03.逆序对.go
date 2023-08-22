package mergeSort

import "fmt"

func ReversePair(arr []int) {
	reverseMap := make(map[string]int)
	process2(arr, 0, len(arr)-1, reverseMap)

	for pair, _ := range reverseMap {
		fmt.Printf("逆序对: %s \n", pair)
	}
}

func process2(arr []int, left, right int, reverseMap map[string]int) {
	if left == right { // 边界条件，分区只有一个元素，无序排序，也无法凑出逆序对
		return
	}
	mid := left + ((right - left) >> 1)
	process2(arr, left, mid, reverseMap)    // 左半区域进行排序，同时打印左半区域的逆序对
	process2(arr, mid+1, right, reverseMap) // 右半区域进行排序，同时打印右半区域的逆序对
	merge2(arr, left, right, mid, reverseMap)

	return
}

func merge2(arr []int, left, right, mid int, reverseMap map[string]int) {
	leftIndex := left
	rightIndex := mid + 1
	temp := make([]int, 0, len(arr))

	for {
		if leftIndex > mid || rightIndex > right {
			break
		}
		if arr[leftIndex] > arr[rightIndex] {
			for i := rightIndex; i > mid; i-- { // 注意，这里是递减，右边有序，因此左区域的该数在原数组中要大于右区域当前数左侧的所有数 (边界是mid+1)
				reverseMap[fmt.Sprintf("%d:%d", arr[leftIndex], arr[i])] = 1
			}
			temp = append(temp, arr[rightIndex])
			rightIndex++
		} else {
			temp = append(temp, arr[leftIndex])
			leftIndex++
		}
	}

	for {
		if leftIndex > mid {
			break
		}
		// 剩余的左半区域数字在原数组中肯定比右半区域所有数字都要大
		for i := mid + 1; i <= right; i++ { // 边界是 mid+1 ~ right
			reverseMap[fmt.Sprintf("%d:%d", arr[leftIndex], arr[i])] = 1
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
}
