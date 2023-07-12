package lesson6

import "math"

func Convert(numStr string) int {
	var neg bool
	if numStr[0] == '-' {
		neg = true
		numStr = numStr[1:]
	} else {
		neg = false
	}

	var res int // 最后转换得到的整数值（总是负数值）
	// 为了防止转化后数值溢出，准备了下面两个变量
	spillCheck := math.MinInt / 10
	spillCheckRe := math.MinInt % 10

	for i := 0; i < len(numStr); i++ {
		eleInt := -int(numStr[i] - '0')
		if res < spillCheck || res == spillCheck && eleInt < spillCheckRe {
			panic("数值转换溢出")
		}

		res = res*10 + eleInt
	}

	if !neg && res == math.MinInt {
		panic("数值转换溢出")
	}
	if neg {
		return res
	} else {
		return -res
	}
}
