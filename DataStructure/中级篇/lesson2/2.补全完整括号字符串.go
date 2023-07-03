package lesson2

// 返回补全一个括号字符串到完整需要的最少括号数
func CompleteParentheses(str string) int {
	leftLack := 0
	rightLack := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			rightLack++
		} else if str[i] == ')' {
			rightLack--
		}
		if rightLack < 0 {
			leftLack++
			rightLack = 0
		}
	}
	return leftLack + rightLack
}
