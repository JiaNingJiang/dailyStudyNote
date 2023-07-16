package lesson9

import "math"

// 将str1编辑成str2，消耗的最小代价
func StrEditDistance(str1, str2 string, icost, dcost, rcost int) int {
	rowCount := len(str1) + 1
	colCount := len(str2) + 1

	dp := make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		dp[i] = make([]int, colCount)
	}

	// 1.初始条件（第一行和第一列是已知的）
	for col := 0; col < colCount; col++ {
		dp[0][col] = col * icost // 第一行，str1为空，str1变成str2只能插入
	}
	for row := 1; row < rowCount; row++ {
		dp[row][0] = row * dcost // 第一行，str2为空，str1变成str2只能删除
	}

	// 2.普通dp求解
	for row := 1; row < rowCount; row++ {
		for col := 1; col < colCount; col++ {
			cost := math.MaxInt
			// 2.1 结尾字符相等的情况，依赖于左上角的dp[row-1][col-1]
			if str1[row-1] == str2[col-1] { // 注意：字符串索引下标要比矩阵中的行与列小1(因为矩阵引入了0长度字符串概念)
				cost = getMin(cost, dp[row-1][col-1])
			}
			// 2.2 将str1[0……i-1]变成str2[0……j]，然后将str1[i]删除
			cost = getMin(cost, dp[row-1][col]+dcost)

			// 2.3 将str1[0……i]变成str2[0……j-1]，然后在str1[i]之后新加一个str2[j]
			cost = getMin(cost, dp[row][col-1]+icost)

			// 2.4 将str1[0……i-1]变成str2[0……j-1]，然后将str1[i]变成str2[j]
			cost = getMin(cost, dp[row-1][col-1]+rcost)

			dp[row][col] = cost
		}
	}

	// 3.返回结果
	return dp[rowCount-1][colCount-1]
}
