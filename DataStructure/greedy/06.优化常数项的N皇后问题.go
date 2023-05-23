package greedy

func NqueenImproved(n int32) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit int // limit用来做位数的限制，比如8皇后问题,对于辅助数int就只能考虑后八位。通过将辅助数与limit做按位与来实现这种保证
	if n == 32 {
		limit = -1 // bit位为1，表示可以取. bit位为0，会被屏蔽  -1时所有bit位为1，表示都可以取
	} else {
		limit = (1 << n) - 1 // 只有后n位为1，可以取
	}
	return nqueenImproved(limit, 0, 0, 0)
}

func nqueenImproved(limit, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == limit { // 所有的N行都填了皇后，因此该方法可行，解法+1
		return 1
	}

	pos := limit & (^(colLim | leftDiaLim | rightDiaLim)) // 下一行所有可以填皇后的位置都为1（具有记忆性，会不断积累前面i-1行产生的影响）
	res := 0
	for { // 每次循环试一位(最后面的可选的一位)，循环变量是pos
		if pos == 0 { // 下一行无处可放皇后
			break
		}
		mostRightOne := pos & (^pos + 1)                  // 取出pos最右侧的1(二进制位)，在这一位放1
		pos = pos - mostRightOne                          // 消去最右侧的1(因为放了1，所以下一次就不可以在这一列上放置了)
		res += nqueenImproved(limit, colLim|mostRightOne, // 行数限制累积
			(leftDiaLim|mostRightOne)<<1,  // 左对角线限制累积(leftDiaLim也要左移一位是因为这是对刚加的新行的下一行的限制 -- 旧行 | 新行 | 新新行(限制) )
			(rightDiaLim|mostRightOne)>>1) // 右对角线限制累积
	}
	return res
}
