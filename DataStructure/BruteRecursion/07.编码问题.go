package BruteRecursion

// 除了用递归之外，还可以用字典树尝试，字典树还能打印出所有的编码可能性。递归只能打印结果数
func CodeProblem(str string) int {
	return codeProblem(str, 0)
}

func codeProblem(str string, index int) int {
	if index == len(str) { // 所有的位都完成了编码，可能性+1
		return 1
	}
	if str[index] == 48 { // 当前字符为字符‘0’,'0'不能单独进行编码，因此此路线失败
		return 0
	}
	if str[index] == 49 { // 当前字符为字符‘1’
		res := codeProblem(str, index+1) // 仅对当前位进行编码
		if index+1 < len(str) {
			res += codeProblem(str, index+2) // 如果可以对当前位+下一位进行共同编码
		}
		return res
	} else if str[index] == 50 { // 当前字符为字符‘2’
		res := codeProblem(str, index+1)                                    // 仅对当前位进行编码
		if index+1 < len(str) && str[index+1] >= 48 && str[index+1] <= 54 { // 第二位还存在，同时必须是字符‘0’~'6'间的一个
			res += codeProblem(str, index+2) // 如果可以对当前位+下一位进行共同编码
		}
		return res
	} else { // 当前字符为3~9
		res := codeProblem(str, index+1) // 仅对当前位进行编码
		return res
	}
}
