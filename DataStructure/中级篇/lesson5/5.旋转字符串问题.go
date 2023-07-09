package lesson5

import "DataStructure2/utils"

// 判断两个字符串是否互为旋转词
func RotationStr(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	str1 += str1 // str1 = str1 + str1

	if loc := utils.KMP(str1, str2, 0); loc != -1 {
		return true
	} else {
		return false
	}
}
