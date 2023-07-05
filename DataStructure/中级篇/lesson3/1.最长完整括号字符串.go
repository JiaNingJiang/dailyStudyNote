package lesson3

import "math"

func LongestPstr(pstr string) int {
	if len(pstr) == 0 {
		return 0
	}
	dp := make([]int, len(pstr))

	for i := 0; i < len(pstr); i++ {
		if pstr[i] == '(' { // pstr[i] == '('
			dp[i] = 0
		} else { // pstr[i] == ')'
			last := i - 1
			if last < 0 { // dp[i-1]是越界的
				continue
			}
			skipInterval := dp[i-1] + 1
			if i-skipInterval < 0 || pstr[i-skipInterval] != '(' { // 跳跃点不存在或者不是'('
				dp[i] = 0
				continue
			}
			dp[i] = dp[i-1] + 2 // dp[i] 至少是 dp[i-1] + 2

			aSkipInterval := skipInterval + 1 // 再往前多跳一个位置
			if i-aSkipInterval >= 0 && dp[i-aSkipInterval] > 0 {
				dp[i] += dp[i-aSkipInterval]
			}
		}
	}
	maxLen := math.MinInt
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
