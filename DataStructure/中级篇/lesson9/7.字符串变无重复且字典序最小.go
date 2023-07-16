package lesson9

func Operation(str string) string {
	if str == "" {
		return ""
	}
	freqMap := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		freqMap[str[i]]++
	}
	minACSIndex := 0
	resSubStr := ""
	curStr := str
	for {
		if len(curStr) == 1 {
			resSubStr += curStr
			return resSubStr
		}
		for i := 0; i < len(curStr); i++ {
			freqMap[curStr[i]]--
			if freqMap[curStr[i]] == 0 {
				// 1.计算在出现词频为0时，之前字符串中具有最小ascii码的下标
				minAscii := uint8(255)
				for index := 0; index <= i; index++ {
					if curStr[index] < minAscii {
						minAscii = curStr[index]
						minACSIndex = index
					}
				}
				// 2.字符追加
				resSubStr += string(curStr[minACSIndex])
				// 3.冗余字符删除
				newCurStr := curStr[minACSIndex+1:]   // 删除掉包含minACSIndex在内的之前的所有字符
				for j := 0; j < len(newCurStr); j++ { // 删除从 minACSIndex+1开始到末尾的所有 curStr[minACSIndex]字符
					if newCurStr[j] == curStr[minACSIndex] {
						newCurStr = newCurStr[:j] + newCurStr[j+1:]
					}
				}
				curStr = newCurStr
				break
			}
		}
		// 重新组件词频表
		freqMap = make(map[uint8]int)
		for i := 0; i < len(curStr); i++ {
			freqMap[curStr[i]]++
		}
	}

}
