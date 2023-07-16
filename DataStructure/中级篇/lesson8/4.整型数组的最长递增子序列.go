package lesson8

func LongestIncrease(order []int) int {
	if len(order) == 0 {
		return 0
	}
	dp := make([]int, len(order))

	for i := 0; i < len(order); i++ {
		// 1.在order[i]之前寻找比起更小的数
		maxDP := 0 // 记录比order[i]更小数字的最大dp累计值
		for j := 0; j < i; j++ {
			if order[j] < order[i] {
				maxDP = getMax(maxDP, dp[j])
			}
		}
		// 2.当前i位置的dp值 == 1 + maxDP(可以是0)
		dp[i] = 1 + maxDP
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = getMax(res, dp[i])
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
