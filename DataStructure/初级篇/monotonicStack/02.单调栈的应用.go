package monotonicStack

import "math"

func process1(arr []int) int {
	mStack := NewMonotonicStack(arr, false) // 递减栈，目的是找到当前元素左右两侧最近的较小元素的下标

	res := mStack.Order()

	maxFactor := math.MinInt
	for _, ele := range arr {
		borders := res[ele]
		leftBorder := borders[0]
		rightBorder := borders[1]
		start := 0
		end := 0

		if leftBorder == nil { // 左侧没有比当前值更小的数，那么意味着左边界可以扩展到arr[0]
			start = 0
		} else {
			start = leftBorder.(Element).index + 1 // 左边界到更小位置的下一个位置
		}

		if rightBorder == nil { // 右侧没有比当前值更小的数，那么意味着右边界可以扩展到arr[len(arr)-1]
			end = len(arr) - 1
		} else {
			end = rightBorder.(Element).index - 1 // 右边界到更小位置的上一个位置
		}

		littleSum := 0
		for i := start; i <= end; i++ {
			littleSum += arr[i]
		}
		factor := littleSum * ele

		if factor > maxFactor {
			maxFactor = factor
		}
	}
	return maxFactor
}
