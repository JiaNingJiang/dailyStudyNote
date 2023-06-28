package stringProblem

func PlainMatch(str, subStr string, pos int) int {
	ssindex := 0  // 遍历子串时的访问下标
	sindex := pos // 遍历原字符串时的访问下标

	for {
		if ssindex >= len(subStr) || sindex >= len(str) { // 子串越界或者原字符串越界
			break
		}

		if subStr[ssindex] == str[sindex] { // 字符匹配
			ssindex++
			sindex++
		} else { // 字符不匹配 （重要：注意这里的先后顺序，必须先让原串的指针回滚，再让子串的指针归零）
			sindex = sindex - ssindex + 1 // 原串回滚(sindex - ssindex是本次匹配开始的位置，+1意味着要从下一个开始重新匹配)
			ssindex = 0                   // 子串下标归零
		}
	}
	if ssindex == len(subStr) { // 子串完成了匹配
		return sindex - len(subStr)
	} else {
		return -1
	}
}
