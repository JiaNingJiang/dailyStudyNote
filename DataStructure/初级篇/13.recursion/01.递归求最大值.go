package recursion

func GetMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, left, right int) int {
	if left == right { // left == right，说明此区间只有一个数，区间的最大值也就是这个
		return arr[left]
	}
	mid := left + ((right - left) >> 1)    // mid = left + (right - left)/2 (右移一位相当于除以二)
	leftMax := process(arr, left, mid)     // 求解左半区域和中点
	rightMax := process(arr, mid+1, right) // 求解右半区域

	return Max(leftMax, rightMax)
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
