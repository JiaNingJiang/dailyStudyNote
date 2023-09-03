package BruteRecursion

import "fmt"

func PrintAllSubString(str string) int {
	charSet := []byte(str)

	subStr := make([]byte, 0)
	return printAllSubString(charSet, 0, subStr)
}

func printAllSubString(str []byte, index int, subStr []byte) int {
	if index == len(str) { // 已经遍历完原字符串的所有字符
		fmt.Println(string(subStr))
		return 1
	}
	subStrCount := 0

	// 要当前字符
	withChar := make([]byte, 0)
	withChar = append(withChar, subStr...)
	withChar = append(withChar, str[index])
	subStrCount += printAllSubString(str, index+1, withChar)

	// 不要当前字符
	noWithChar := make([]byte, 0)
	noWithChar = append(noWithChar, subStr...)
	subStrCount += printAllSubString(str, index+1, noWithChar)

	return subStrCount
}
