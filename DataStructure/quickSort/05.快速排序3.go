package quickSort

import (
	"math/rand"
	"time"
)

func QuickSort3(arr []int) {
	left := 0
	right := len(arr) - 1
	sort3(arr, left, right)
}

func sort3(arr []int, left, right int) {
	if left >= right {
		return
	}
	leftIndex := left
	rightIndex := right - 1
	i := left

	targetIndex := getRandomTarget(arr, left, right) // 从当前区域中获取随机获取一个数作为target
	swap(&arr[targetIndex], &arr[right])             // 将该数与区域的最后一个数进行交换
	target := arr[right]

	for {
		if i > rightIndex {
			swap(&arr[right], &arr[rightIndex+1])
			break
		}
		if arr[i] < target {
			swap(&arr[i], &arr[leftIndex])
			leftIndex++
			i++
		} else if arr[i] == target {
			i++
		} else {
			swap(&arr[i], &arr[rightIndex])
			rightIndex--
		}
	}

	sort3(arr, left, leftIndex-1)
	sort3(arr, rightIndex+2, right)

}

func getRandomTarget(arr []int, left, right int) int {
	rand.Seed(time.Now().UnixNano())
	if left >= right {
		return right
	}
	index := rand.Intn(right-left+1) + left

	return index
}
