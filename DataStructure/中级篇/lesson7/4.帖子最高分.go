package lesson7

import "math"

func PostMaxScore(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	maxScore := math.MinInt

	curScore := 0
	for i := 0; i < len(scores); i++ {
		curScore += scores[i]
		if curScore > maxScore {
			maxScore = curScore
		}
		if curScore < 0 { // 如果连续分数之和 < 0, 那么下一次从零继续开始
			curScore = 0
		}
	}
	return maxScore
}
