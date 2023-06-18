package bruteRecursionPromote

// coins必须是一个没有重复元素的数组
func CoinCase(coins []int, target int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}
	if target == 0 {
		return -1
	}

	return coinCase(coins, target, 0)
}

func coinCase(coins []int, res int, cur int) int {
	if res == 0 {
		return 1
	}
	if cur == len(coins) {
		return 0
	}
	ways := 0
	for zhang := 0; coins[cur]*zhang <= res; zhang++ {
		ways += coinCase(coins, res-coins[cur]*zhang, cur+1)
	}
	return ways
}

func CoinCaseStrictTable(coins []int, target int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}
	if target == 0 {
		return -1
	}

	table := make([][]int, len(coins)+1) // 行的含义是当前coins集合的索引 ( 0 ~ len(coins) )
	for line := 0; line <= len(coins); line++ {
		table[line] = make([]int, target+1) // 列的含义是当前所需的余额( 0 ~ target)
	}

	table[len(coins)][0] = 1 // 最后一行只有第一个元素 存在一条可能的way，因此为1。其余元素都为0
	for col := 1; col <= target; col++ {
		table[len(coins)][col] = 0
	}

	for line := len(coins) - 1; line >= 0; line-- { // 从倒数第二行向上计算，因为上层的结果依赖于下层
		for col := 0; col <= target; col++ {
			ways := 0 // 当前(line,col)元素对应的可能性
			for zhang := 0; coins[line]*zhang <= col; zhang++ {
				ways += table[line+1][col-coins[line]*zhang]
			}
			table[line][col] = ways
		}
	}
	return table[0][target]
}
