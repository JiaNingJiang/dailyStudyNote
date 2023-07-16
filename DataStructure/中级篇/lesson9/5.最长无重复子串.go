package lesson9

import "math"

func MaxUnique(str string) int {
	if str == "" {
		return 0
	}

	recentCharMap := make(map[uint8]int) // 用于记录不同字符出现在字符串中的距离结尾最近的位置

	for i := 0; i <= 255; i++ {
		recentCharMap[uint8(i)] = -1
	}

	maxLen := math.MinInt
	lastLeft := -1 // i-1 位置的无重复子串的左边界位置
	index := 0     // 访问下标i

	for {
		if index >= len(str) {
			return maxLen
		}
		lastLeft = getMax(lastLeft, recentCharMap[str[index]]) // 当前以index结尾的无重复字符串的左边界
		recentCharMap[str[index]] = index
		curLen := index - lastLeft
		maxLen = getMax(maxLen, curLen)
		index++
	}

}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
