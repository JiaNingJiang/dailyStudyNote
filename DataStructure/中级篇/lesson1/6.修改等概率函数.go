package lesson1

// 将一个等概率返回1~5的函数修改为等概率返回0和1的函数
func EqualP01(f func() int) int {
	for {
		res := f()

		if res == 0 || res == 1 {
			return 0
		} else if res == 2 || res == 3 {
			return 1
		}
	}
}

// 将一个等概率返回1~5的函数修改为等概率1~7的函数
func EqualP1_7(f func() int) int {
	// 生成一个等概率返回 0~6 的函数(0~6至少需要3bit表示)
	res := 0
	for {
		res = EqualP01(f)<<2 + EqualP01(f)<<1 + EqualP01(f)<<0

		if res == 7 { // 3bit生成了7(111)，则需要重新生成
			continue
		} else {
			break
		}
	}
	return res + 1 // 0~6 + 1 == 1~7
}

// 将一个等概率返回1~5的函数修改为等概率30~59的函数
func EqualP30_59(f func() int) int {
	// 生成一个等概率返回 0~29 的函数(0~29至少需要5bit表示)
	res := 0
	for {
		res = EqualP01(f)<<5 + EqualP01(f)<<4 + EqualP01(f)<<3 +
			EqualP01(f)<<2 + EqualP01(f)<<1 + EqualP01(f)<<0

		if res == 30 || res == 31 { // 5bit生成了30(11110)或者31(11111)，则需要重新生成
			continue
		} else {
			break
		}
	}
	return res + 30 // 0~29 + 30 == 30~59
}
