package lesson5

func AdjustArr(arr []int) bool {
	twoCount := 0      // 累积数组中2的个数
	oddCount := 0      // 累积数组中奇数的个数
	fourTimeCount := 0 // 累积数组中4倍数的个数

	for i := 0; i < len(arr); i++ {
		if arr[i] == 2 {
			twoCount++
		} else if arr[i]%2 != 0 {
			oddCount++
		} else if arr[i]%4 == 0 {
			fourTimeCount++
		}
	}

	if twoCount == 0 { // 数组中没有2，则需要将数组调整为： 奇4奇4奇4…… 或者  4奇4奇4奇…… 的形式
		if oddCount == 1 && fourTimeCount == 1 {
			return true
		}
		if oddCount > 1 && fourTimeCount >= oddCount-1 { // 第一种情况是fourTimeCount >= oddCount-1；第二种情况是fourTimeCount >= oddCount。综合就是fourTimeCount >= oddCount-1
			return true
		}
		return false
	} else { // 数组中有2，则需要将数组调整为: 222……4奇4奇4奇…… 的形式
		if oddCount == 1 && fourTimeCount == 1 {
			return true
		}
		if oddCount > 1 && fourTimeCount >= oddCount {
			return true
		}
		return false
	}
}
