package dichotomy

// 1.在一个有序数组中，找到某个数是否存在
func IsExist(arr []int, target int, left int, middle int, right int) (bool, int) {
	if arr[middle] == target {
		return true, arr[middle]
	}
	if left > right {
		return false, -1
	}
	if arr[middle] > target {
		newLeft := left
		newRight := middle - 1                          // 关键，arr[middle]已经完成比较任务，下一次递归不需要再次进行比较。(如果不更新此边界，递归将无法退出)
		newMid := newLeft + ((newRight - newLeft) >> 1) // 计算中位数的公式
		return IsExist(arr, target, newLeft, newMid, newRight)
	} else {
		newLeft := middle + 1 // // 关键，arr[middle]已经完成比较任务，下一次递归不需要再次进行比较
		newRight := right
		newMid := newLeft + ((newRight - newLeft) >> 1)
		return IsExist(arr, target, newLeft, newMid, newRight)
	}
}

// 2.在一个有序数组中，找到大于某个数的最左侧位置
func AlmostLeft(arr []int, target int, left, middle, right int) (int, int) {
	if left > right { // 当left == right时区域缩减到一个数字,此时因为middle-1或是middle+1操作使得left > right
		return left, arr[left]
	}

	if arr[middle] > target {
		newLeft := left
		newRight := middle - 1
		newMid := newLeft + ((newRight - newLeft) >> 1)
		return AlmostLeft(arr, target, newLeft, newMid, newRight)
	} else {
		newLeft := middle + 1
		newRight := right
		newMid := newLeft + ((newRight - newLeft) >> 1)
		return AlmostLeft(arr, target, newLeft, newMid, newRight)
	}

}

// 3.局部最小值问题:在一个无序数组中，查询任意一个局部最小值
func LocalMin(arr []int) (int, int) {
	if arr[0] < arr[1] { // arr[0]为局部最小值
		return 0, arr[0]
	}
	if arr[len(arr)-2] > arr[len(arr)-1] { // arr[N-1]为局部最小值
		return len(arr) - 1, arr[len(arr)-1]
	}
	// 局部最小值在中间(两侧是递减的，最小值必然在中间)
	left := 1
	right := len(arr) - 2
	mid := (left + right) / 2
	return localMin(arr, left, mid, right)

}

func localMin(arr []int, left, middle, right int) (int, int) {
	if arr[middle] < arr[middle-1] && arr[middle] < arr[middle+1] { // 符合局部最小值的定义(比两侧都要小)
		return middle, arr[middle]
	}
	if arr[middle] > arr[middle-1] { // 左半区域是两侧递减的(仅在此区域内存在一个局部最小值)
		newLeft := left
		newRight := middle - 1
		newMid := (newLeft + newRight) / 2
		return localMin(arr, newLeft, newMid, newRight)
	} else { // 右半区域树两侧递减的(仅在此区域内存在一个局部最小值)
		newLeft := middle + 1
		newRight := right
		newMid := (newLeft + newRight) / 2
		return localMin(arr, newLeft, newMid, newRight)
	}
}
