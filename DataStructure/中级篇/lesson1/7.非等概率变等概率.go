package lesson1

import "fmt"

// 将非等概率返回0/1的函数修改为等概率返回0/1的函数
func EqualP(f func() int) int {
	for {
		str := ""
		for i := 0; i < 2; i++ { // 使用两次f函数
			str += fmt.Sprintf("%d", f())
		}
		if str == "01" { // p(1-p)的概率
			return 0
		} else if str == "10" { // p(1-p)的概率
			return 1
		}
		// 如果摇出的是 00(p*p) 或者 11((1-p)*(1-p), 则需要重新摇
	}
}
