package lesson9

import "math"

// start,end 分别表示初始人气值和目标人气值
// add 表示人气值+2 需要消耗的金币
// twice 表示人气值*2 需要消耗的金币
// del 表示人气值-2 需要消耗的金币
func PopuValue(start, end int, add, twice, del int) int {
	return popuValue(start, end, add, twice, del, 0, 0)
}

func popuValue(start, end int, add, twice, del int, curPopu int, curCoin int) int {
	if curPopu == end { // 达到目标人气值
		return curCoin
	}
	if curPopu >= 2*end { // 最优情况下，人气值不可能超过end的2倍
		return -1
	}
	if curPopu < 0 { // 最优情况下，人气值不可能小于0
		return -1
	}

	if curCoin > (end-start)/2*add { // 消耗的金币数比只凭点赞消耗的更多，就不可能是最优情况
		return -1
	}

	addRes := popuValue(start, end, add, twice, del, curPopu+2, curCoin+add)
	twiceRes := popuValue(start, end, add, twice, del, curPopu*2, curCoin+twice)
	delRes := popuValue(start, end, add, twice, del, curPopu-2, curCoin+del)

	// 只有当三种分区的结果都为-1的时候，返回值才会是-1
	if addRes == -1 && twiceRes == -1 && delRes == -1 {
		return -1
	}

	// 只要存在有分支！=-1的情况下，返回有效分支中消耗金币数最少的
	addCoin := math.MaxInt
	twiceCoin := math.MaxInt
	delCoin := math.MaxInt

	if addRes != -1 {
		addCoin = addRes
	}
	if twiceRes != -1 {
		twiceCoin = twiceRes
	}
	if delRes != -1 {
		delCoin = delRes
	}

	res := getMin(getMin(addCoin, twiceCoin), delCoin)

	return res
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
