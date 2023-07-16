package lesson9

import "math"

func LookUpMissing(arr []int) []int {
	n := len(arr) // 在arr数组中出现数字的最大可能值（最小可能值是1）

	for i := 0; i < n; i++ { // 每轮交易都以从当前第i为开始，以重新回到第i位结束

		if arr[i] == i+1 { // 直接不需要交换，结束本轮
			continue
		}

		curVal := arr[i]
		curIndex := i

		// 需要交换，意味着需要从当前位置拿走一个数，导致该位置的空缺(只有每轮的第一次取是只取不填，因此导致空缺)
		temp := arr[curVal-1]
		arr[curVal-1] = curVal
		curVal = temp // 通过交换而获得的值

		arr[curIndex] = math.MinInt // 用来表示空缺
		curIndex = curVal - 1       // 下一个需要检查的位置

		for {

			if arr[curIndex] == curVal { // 如果arr[curVal-1] == curVal,说明不需要进行交换了，那么本轮结束
				break
			}
			// 此时 arr[curVal-1] != curVal ,那么令 arr[curVal-1] = curVal，并且把 arr[curVal-1] 原本的值拿到
			temp1 := arr[curVal-1]
			arr[curVal-1] = curVal
			curVal = temp1

			if curVal == math.MinInt { // 说明原本的位置是没有值的，刚好把一个值填了进去
				break // 因为本来是空的，因此通过交换无法获取新值，因此结束本轮交换
			}

			curIndex = curVal - 1 // 下一个需要被交换的数组元素的下标是 curVal - 1

		}
	}

	absenceSet := make([]int, 0)
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			absenceSet = append(absenceSet, i+1)
		}
	}
	return absenceSet
}
