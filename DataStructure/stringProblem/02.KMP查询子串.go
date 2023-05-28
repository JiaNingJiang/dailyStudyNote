package stringProblem

// 生成子串的nextArray数组
func getNextArray(subStr string) []int {
	if len(subStr) == 0 {
		return nil
	}
	if len(subStr) == 1 { // 规定：nextArray[0] = -1
		return []int{-1}
	}
	nextArr := make([]int, len(subStr))

	nextArr[0] = -1 // 规定：nextArray[0] = -1
	nextArr[1] = 0  // 规定：nextArray[1] = 0

	i := 2  // 计算i之前的最大相等前后缀字符串的长度 ( i-1 则指向最大后缀字符串的下一个字符位置)
	cn := 0 // 总是指向最大前缀字符串的下一个字符位置(不属于前缀字符串的第一个字符)

	for {
		if i == len(subStr) { // 子串所有位的nextArr都已经完成计算
			return nextArr
		}
		if subStr[i-1] == subStr[cn] { // 判断前后缀字符串各自的下一个字符是否相等，相等则进行前后缀字符串的扩大
			cn++
			nextArr[i] = cn // 是一个累积量
			i++             // 继续计算下一位的nextArr
		}
		if cn > 0 { // 前后缀字符串各自的下一个字符不相等。缩小前后缀字符串(跳转到之前(上一个)的最大前缀字符串处进行比较)
			cn = nextArr[cn] // i不增加，而是在下一次循环时重新比较
		} else { // 无法再找到更小的前后缀字符串(意味着当前i之前没有任何前后缀字符串相等)
			nextArr[i] = 0
			i++ // 继续计算下一位的nextArr
		}
	}
}

func KMP(str, subStr string, pos int) int {
	if len(subStr) < 1 || len(str) < len(subStr) {
		return -1
	}
	strIndex := pos
	subIndex := 0
	next := getNextArray(subStr) // 获取子串的nextArray    O(M)

	for { // O(N)  因为 O(M)必定小于O(N),所以KMP算法时间复杂度为O(N)
		if strIndex >= len(str) || subIndex >= len(subStr) {
			break
		}
		if str[strIndex] == subStr[subIndex] {
			strIndex++
			subIndex++
		} else if next[subIndex] == -1 { // 当前子串索引处之前已经没有任何相等的前后缀子串,而且来到了 subIndex == 0 ，即子串的起始位置
			strIndex++ // 主串放弃，直接从下一个字符开始比较
		} else {
			subIndex = next[subIndex] // 子串直接跳到前缀字符串(因为子串的前缀字符串必定和主串的后缀字符串相等，而且是最大的前后缀字符串)之后开始进行比较
		}
	}
	if subIndex == len(subStr) { // 退出循环的原因是：子串完成了遍历(意味着在主串中完成了匹配)
		return strIndex - subIndex // 返回主串开始匹配该子串的起始位置
	} else {
		return -1
	}
}
