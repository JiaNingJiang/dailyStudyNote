package lesson9

func LogicalOpt(logicStr string, desire bool) int {
	if logicStr == "" {
		return 0
	}
	if !isValid(logicStr) {
		return 0
	}

	return logicalOpt(logicStr, desire, 0, len(logicStr)-1)

}

func logicalOpt(logicStr string, desire bool, left, right int) int {
	// 边界条件
	if left > right {
		return 0
	}

	if left == right { // 此时只剩一个字符
		if desire { // 目标渴望是true
			if logicStr[left] == '1' {
				return 1
			} else if logicStr[left] == '0' {
				return 0
			}
		} else { // 目标渴望是false
			if logicStr[left] == '1' {
				return 0
			} else if logicStr[left] == '0' {
				return 1
			}
		}
	}

	// 正常递归
	res := 0 // 可能的结果数

	if desire { // 目标渴望是true
		for i := left + 1; i < right; i += 2 { // 遍历每一个逻辑运算符(注意：这里每次只需要遍历当前区域内(left~right)的奇数位置)
			op := logicStr[i]
			switch op {
			case '&': // 左右必须都是true
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '|': // 左右有一个是true即可
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '^': // 左右必须相异
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			}
		}
	} else { // 目标渴望是false
		for i := left + 1; i < right; i += 2 { // 遍历每一个逻辑运算符
			switch logicStr[i] {
			case '&': // 左右有一个是false即可
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '|': // 左右必须都是false
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
			case '^': // 左右必须相同
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
			}
		}
	}
	return res
}

func LogicalOptDP(logicStr string, desire bool) int {
	if logicStr == "" {
		return 0
	}
	if !isValid(logicStr) {
		return 0
	}

	return logicalOptDP(logicStr, desire)

}

func logicalOptDP(logicStr string, desire bool) int {
	N := len(logicStr)

	trueDP := make([][]int, N)
	for i := 0; i < N; i++ {
		trueDP[i] = make([]int, N)
	}

	falseDP := make([][]int, N)
	for i := 0; i < N; i++ {
		falseDP[i] = make([]int, N)
	}

	// 1.设置初始条件(矩阵对角线) -- 两个矩阵(一个是最终返回结果是true，一个是最终返回结果为false)
	for i := 0; i < N; i += 2 { // 每次都要访问数字(left和right都在偶数位置上)
		if logicStr[i] == '0' {
			trueDP[i][i] = 0
			falseDP[i][i] = 1
		} else if logicStr[i] == '1' {
			trueDP[i][i] = 1
			falseDP[i][i] = 0
		}
	}

	// 2.根据初始条件获取其他位置(从下往上，从左向右求解)
	for left := N - 3; left >= 0; left -= 2 { // 范围的左边界(需要是数字)  N-1作为对角线是初始条件不用求
		for right := left + 2; right < N; right += 2 { // 范围的右边界(需要是数字)
			for oper := left + 1; oper < right; oper += 2 { // 遍历从左边界到右边界范围内的所有逻辑运算符
				// 2.1 计算trueDP
				switch logicStr[oper] {
				case '&': // logicStr[oper] 左右两侧都必须是true
					trueDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
				case '|':
					trueDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
					trueDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					trueDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				case '^':
					trueDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					trueDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				}
				// 2.2 计算falseDP
				switch logicStr[oper] {
				case '&':
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
					falseDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					falseDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				case '|':
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
				case '^':
					falseDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
				}
			}
		}
	}

	// 3.返回结果(左边界为0，右边界为N-1的矩阵元素即是要求解的目标)
	if desire {
		return trueDP[0][N-1]
	} else {
		return falseDP[0][N-1]
	}

}

// 判断一个逻辑字符串是否有效
func isValid(logicStr string) bool {
	length := len(logicStr)

	if length%2 == 0 { // 字符串必须是奇数长度
		return false
	}

	for i := 0; i < length; i += 2 { // 偶数位置必须是0或1
		if logicStr[i] != '0' && logicStr[i] != '1' {
			return false
		}
	}

	for i := 1; i < length; i += 2 { // 奇数位置必须是逻辑运算符 & | ^
		if logicStr[i] != '&' && logicStr[i] != '|' && logicStr[i] != '^' {
			return false
		}
	}

	return true
}
