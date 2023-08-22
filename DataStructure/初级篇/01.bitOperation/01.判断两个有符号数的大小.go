package bitOperation

// 此函数，如果输入非负数则返回1；如果输入负数则返回0 （需要注意：只有比较的数字都是int32, >>31才能提取出符号位）
func sign(n int32) int32 {
	signal := (n >> 31) & 1
	return flip(signal)
}

// 该函数的参数，只能是1和0
// 如果输入1，返回0；如果输入0，则返回1
func flip(n int32) int32 {
	return n ^ 1
}

// 不考虑a-b溢出的情况
func GetMax1(a, b int32) int32 {
	c := a - b
	sca := sign(c)   // 获取c的符号位(非负数为1，负数为0)
	scb := flip(sca) // 总是与sca相异
	return sca*a + scb*b
}

// 考虑了a-b可能溢出的情况
func GetMax2(a, b int32) int32 {
	c := a - b
	sa := sign(a)
	sb := sign(b)
	sc := sign(c)

	difSab := sa ^ sb
	sameSab := flip(difSab)

	returnA := difSab*sa + sameSab*sc // a、b异号的时候根据两数的符号决定返回值；a、b同号的时候根据c的符号决定返回值
	returnB := flip(returnA)

	return returnA*a + returnB*b
}
