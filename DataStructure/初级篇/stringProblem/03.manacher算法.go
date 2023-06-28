package stringProblem

import "math"

// 输入一个字符串，返回该字符串的最大回文子串的直径(总长度)
func Manacher(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := manacherFormat(s, "#") // 1221 -> #1#2#2#1#
	pArr := make([]int, len(str)) // 回文半径数组
	C := -1                       // 最大回文区域中心
	R := -1                       // 最大回文区域右边界的再往右一个位置(即出回文区域的下一个位置)
	max := math.MinInt            // 记录最大回文半径(增加'#'后的)

	for i := 0; i < len(str); i++ {
		// 1.第一步：确定i至少的回文区域，由pArr[i]进行记录
		if i >= R { // 在最大回文区域之外，那么i至少的回文区域是1，也就是单独该字符本身构成的字符串
			pArr[i] = 1
		} else { // 还在最大回文区域内
			pArr[i] = int(math.Min(float64(pArr[2*C-i]), float64(R-i))) // pArr[2*C-i]是i'的回文半径，因为 (i+i')/2 = C
		}

		// 2.第二步：再进一步试探回文区域能否再外扩，能的话不断扩大pArr[i]
		for {
			if i+pArr[i] >= len(str) || i-pArr[i] <= -1 { // i的回文区域右边界或左边界溢出
				break
			}
			if str[i+pArr[i]] == str[i-pArr[i]] { // 左右边界可以外扩
				pArr[i]++
			} else { // 不能外扩了，此即是str[i]的最大回文区域
				break
			}
		}

		if i+pArr[i] > R { // 更新最大回文区域右边界R及中点C
			R = i + pArr[i]
			C = i
		}
		max = int(math.Max(float64(max), float64(pArr[i]))) // 记录最大回文半径(增加'#'后的)，注意：最大回文半径不是指最大回文区域的半径
	}
	return max - 1 // 修改后的回文字符串半径 = 原回文字符串半径 * 2 +1 = 原回文字符串直径(总长度) +1
}

// 将普通字符串变为求解回文字符串的特殊形式(每两个字符之间插入一个特殊字符,只能是ascii码字符)
func manacherFormat(str string, special string) string {

	originStr := []byte(str)
	sep := []byte(special)
	format := make([]byte, 2*len(str))

	for i := 0; i < len(format); i++ {
		if i%2 == 0 { // 偶数位置按序等于原始字符串的个字符
			format[i] = originStr[i/2]
		} else { // 奇数位置等于特殊字符串
			format[i] = sep[0]
		}
	}
	return special + string(format)
}
