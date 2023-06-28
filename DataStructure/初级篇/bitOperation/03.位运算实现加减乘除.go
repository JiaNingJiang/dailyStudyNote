package bitOperation

func Add(a, b int32) int32 {
	noCarryAdd := a
	carryInfo := b
	for {
		if carryInfo == 0 {
			break
		}
		noCarryAddTemp := noCarryAdd ^ carryInfo
		carryInfoTemp := (noCarryAdd & carryInfo) << 1

		noCarryAdd = noCarryAddTemp
		carryInfo = carryInfoTemp
	}
	return noCarryAdd
}

// 将一个数变为其相反数 --> 取反+1
func negNum(a int32) int32 {
	return Add(^a, 1)
}

// a-b == a+(-b)
func Sub(a, b int32) int32 {
	return Add(a, negNum(b))
}

func Mul(a, b int32) int32 {
	var res int32 = 0
	var multiplicand int32 = a // 被乘数
	var bitProduct int32 = b   // 乘数与被乘数某一位的bit乘积(要么为0，要么为b<<n)

	// 遍历被乘数的每一bit位，与乘数相乘
	for {
		if multiplicand == 0 { // 完成被乘数的bit遍历
			break
		}
		if (multiplicand & 1) != 0 { // 获取被乘数当前的最后一位，如果是1的话加上当前位的bit乘积
			res = Add(res, bitProduct)
		}
		multiplicand = multiplicand >> 1 // 当前最后一位完成遍历，再遍历下一位
		bitProduct = bitProduct << 1
	}
	return res
}

// 检查一个int32是否是负数
func isNeg(a int32) int {
	if (a>>31)&1 == 1 {
		return 1
	} else {
		return 0
	}
}

func Div(dividend, divisor int32) int32 {
	dividendTemp := dividend  // 用于保存符号
	divisorTemp := divisor    // 保存符号位
	if isNeg(dividend) == 1 { // 保证参与除法运算的被除数和除数都是整数
		dividend = negNum(dividend)
	}
	if isNeg(divisor) == 1 {
		divisor = negNum(divisor)
	}
	var res int32 = 0
	var i int32
	for i = 31; i >= 0; i = Sub(i, 1) {
		if (dividend >> i) >= divisor { // 如果 被除数右移之和 >= 除数 (除数左移可能会导致除数溢出（符号位被影响）)
			res |= (1 << i)                      // 商对应位置为1
			dividend = Sub(dividend, divisor<<i) // 更新被除数
		}
	}
	if isNeg(dividendTemp)^isNeg(divisorTemp) == 1 { // 被除数与除数符号相反，那么商必然是负数
		return negNum(res)
	} else {
		return res
	}
}
