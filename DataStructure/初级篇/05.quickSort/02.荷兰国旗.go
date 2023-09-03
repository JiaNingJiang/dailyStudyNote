package quickSort

func DutchFlag(arr []int, target int) {
	process2(arr, target)
}

func process2(arr []int, target int) {
	leftIndex := -1
	rightIndex := len(arr)
	i := 0

	for {
		if i >= rightIndex { // 右区域必然都是比target大的数字，没有继续向下遍历的必要(重要：这里的rightIndex从len(arr)开始，因此是≥)
			break
		}
		//if i == len(arr) {
		//	break
		//}
		if arr[i] < target {
			leftIndex++
			swap(&arr[i], &arr[leftIndex])
			i++ // 可以 i++ , 因为左侧区域的数字都已经经过检测(因为整体是从左向右遍历的)，必定位于 <= target区域。
		} else if arr[i] == target {
			i++
		} else {
			rightIndex--
			swap(&arr[i], &arr[rightIndex])
			//i++      // 这里不能有 i++ , 因为从右侧区域并没有进行过检测，因此移动过来的数字可能不在 <= target 范围内
		}
	}

}
