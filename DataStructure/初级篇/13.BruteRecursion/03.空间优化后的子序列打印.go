package BruteRecursion

import "fmt"

func PrintAllSubStrImproved(str string) int {
	charSet := []byte(str)
	return printAllSubStringImproved(charSet, 0)
}

func printAllSubStringImproved(str []byte, index int) int {
	if index == len(str) {
		fmt.Println(string(str))
		return 1
	}

	subStrCount := 0
	subStrCount += printAllSubStringImproved(str, index+1) // 要当前字符

	oldData := str[index]
	str[index] = 32                                        // 空格字符的ASCII码
	subStrCount += printAllSubStringImproved(str, index+1) // 不要当前字符(将原本字符串对应位变为空格)

	str[index] = oldData // 再恢复到原来的状态

	return subStrCount
}
