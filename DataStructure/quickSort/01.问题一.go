package quickSort

// 将数组arr中小于等于target统一移动到左半区域，右半区域都是大于target的数字
func DiscreteSort(arr []int, target int) {
	process(arr, target)
}

func process(arr []int, target int) {
	leftIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= target {
			swap(&arr[leftIndex], &arr[i])
			leftIndex++
		}
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
