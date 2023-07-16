package lesson9

// 求出以begin开头(begin: 0~25 表示26个不同字母)，长度为length的字符串个数
func spBeginFixedLen(begin int, length int) int {
	if length == 1 { // 长度为1，开头又固定，那必然只有一种可能
		return 1
	}

	sum := 0
	for i := begin + 1; i <= 25; i++ { // 子字符串的开头必须从 begin+1 开始，且长度-1
		sum += spBeginFixedLen(i, length-1)
	}

	return sum
}

// 求出指定长度的所有可能字符串的数量
func fixedLen(length int) int {
	sum := 0

	for i := 0; i <= 25; i++ {
		sum += spBeginFixedLen(i, length)
	}
	return sum
}

func DesignatedStrIndex(str string) int {
	index := 0
	strLen := len(str)

	// 长度为1的字符串求解是特殊的
	if strLen == 1 {
		end := int(str[0] - 'a')
		for i := 0; i < end; i++ {
			index++
		}
		return index + 1
	}

	for i := 1; i <= strLen-1; i++ {
		index += fixedLen(i) // 长度比当前字符串小的必然在前面
	}

	curStr := str
	for {
		if len(curStr) <= 1 {
			return index + 1 // index只是当前字符串之前的字符串个数，因此还需要+1
		}
		left := int(curStr[0]-'a') + 1 // 长度至少为2
		right := int(curStr[1]-'a') - 1

		for j := left; j <= right; j++ {
			index += spBeginFixedLen(j, len(curStr)-1)
		}

		curStr = curStr[1:] // 每次循环去掉一个最左侧的字符

	}

}
