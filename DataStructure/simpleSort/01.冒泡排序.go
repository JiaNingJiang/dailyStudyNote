package simpleSort

func BubbleSort(arr []int) {
	flag := true
	for i := 0; i < len(arr); i++ {
		if !flag { // 意味着上一轮没有进行任何交换，数组本身已经是有序的了
			break
		}
		flag = false
		for j := len(arr) - 2; j >= i; j-- { // 从数组尾部开始。将小数上浮，大数下沉
			if arr[j] > arr[j+1] {
				flag = true
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
