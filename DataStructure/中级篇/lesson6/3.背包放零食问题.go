package lesson6

// snacks: 记录各个零食的重量
// bagSize: 背包的大小
// 返回值: 可以的零食放法
func BagSnack(snacks []int, bagSize int) int {
	return bagSnack(snacks, bagSize, 0)
}

func bagSnack(snacks []int, remainBag, index int) int {
	if index >= len(snacks) {
		return -1
	}
	if remainBag == 0 {
		return 1
	}
	if remainBag < 0 {
		return -1
	}

	yao := bagSnack(snacks, remainBag-snacks[index], index+1)
	buyao := bagSnack(snacks, remainBag, index+1)

	if yao == -1 && buyao == -1 {
		return -1
	} else if yao == -1 {
		return buyao
	} else if buyao == -1 {
		return yao
	} else {
		return yao + buyao
	}
}

// snacks: 记录各个零食的重量
// bagSize: 背包的大小
// 返回值: 可以的零食放法
func BagSnackDP(snacks []int, bagSize int) int {
	return bagSnackDP(snacks, bagSize)
}

func bagSnackDP(snacks []int, bagSize int) int {

	// 行表示背包内零食重量，从0~bagSize，共bagSize+1行
	// 列表示当前遍历到零食下标，0~len(snacks)-1, 共len(snacks)列
	// 初始已知：最后一行全1，因为零食重量 == bagSize，且零食数组没有越界
	// 目标求解值: (0,0)
	// 依赖关系： matrix[size][i] = matrix[size+snacks[i]][i+1] + matrix[size][i+1]

	dp := make([][]int, bagSize+1)
	for i := 0; i <= bagSize; i++ {
		dp[i] = make([]int, len(snacks))
	}

	// 初始条件
	for col := 0; col < len(snacks); col++ {
		dp[bagSize][col] = 1
	}

	for row := bagSize - 1; row >= 0; row-- { // 从下往上(因为dp[i][j]依赖于下方)
		for col := len(snacks) - 1; col >= 0; col-- { // 从右向左(因为dp[i][j]依赖于右方)
			lowerRow := row + snacks[col]
			rightCol := col + 1
			if lowerRow <= bagSize && rightCol < len(snacks) { // size+snacks[i]和i+1都没有越界
				dp[row][col] = dp[lowerRow][col+1] + dp[row][col+1]
			} else if rightCol >= len(snacks) { //i+1都越界
				dp[row][col] = 0
			} else if lowerRow > bagSize && rightCol < len(snacks) { // 只有size+snacks[i]越界
				dp[row][col] = dp[row][col+1]
			}
		}
	}
	return dp[0][0]
}
