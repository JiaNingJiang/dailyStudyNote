package bruteRecursionPromote

func MinCoin(coinSet []int, target int) int {
	if target < 1 {
		return -1
	}

	return minCoin(coinSet, target, 0, 0, 0)
}

func minCoin(coinSet []int, target int, index int, curWorth int, curCoin int) int {
	if curWorth == target { // 边界条件一:一旦凑齐target金额，立刻返回总共所需的硬币数（重要：必须在边界条件二之前）
		return curCoin
	}
	if index >= len(coinSet) { // 边界条件二:遍历完所有的硬币，也没能凑成target金额
		return -1
	}
	if curWorth > target { // 边界条件三:一旦当前金额大于target，立即返回-1，表明此分支不可用
		return -1
	}
	yao := minCoin(coinSet, target, index+1, curWorth+coinSet[index], curCoin+1)
	buyao := minCoin(coinSet, target, index+1, curWorth, curCoin)

	if yao == -1 && buyao == -1 { // 两种情况都返回-1，表明都无法凑齐target金额
		return -1
	} else if buyao == -1 { // 只有"不要"这种情况无法凑齐target金额，那么就返回"要"这种情况所需的硬币总数
		return yao
	} else if yao == -1 { // 只有"要"这种情况无法凑齐target金额，那么就返回"不要"这种情况所需的硬币总数
		return buyao
	} else { // 两种情况都可以凑齐target金额，那么返回所需硬币数最小的情况
		if yao < buyao {
			return yao
		} else {
			return buyao
		}
	}
}

func MinCoinMemoryCache(coinSet []int, target int) int {
	if target < 1 {
		return -1
	}
	dpMemory := make([][]int, len(coinSet)+1) //
	for i := 0; i < len(coinSet)+1; i++ {     // 横坐标表示当前本轮选中的硬币
		dpMemory[i] = make([]int, target+1) // 纵坐标表示当前剩余需要的钱数(必然在 0 ~ target范围内)，不能使用已搜集钱数作为纵坐标(因为范围不止0~target)
		for j := 0; j <= target; j++ {
			dpMemory[i][j] = -1
		}
	}
	return minCoinMemoryCache(coinSet, target, 0, target, dpMemory)
}

func minCoinMemoryCache(coinSet []int, target int, index int, rest int, dpMemory [][]int) int {
	if rest == 0 { // 边界条件一:一旦凑齐target金额，立刻返回总共所需的硬币数（重要：必须在边界条件二之前）
		dpMemory[index][0] = 0 // 等于0表示当凑齐target之和，不会额外再花任何一枚硬币
		return 0
	}
	if index >= len(coinSet) { // 边界条件二:遍历完所有的硬币，也没能凑成target金额
		return -1
	}
	if rest < 0 { // 边界条件三:一旦当前金额大于target，立即返回-1，表明此分支不可用
		return -1
	}
	// target金额没有凑满，同时也不是边界条件
	if dpMemory[index][rest] != -1 { // 缓存命中
		return dpMemory[index][rest]
	}
	// 缓存未命中，也就是新情况
	// 金币不够(即rest > 0)
	yao := minCoinMemoryCache(coinSet, target, index+1, rest-coinSet[index], dpMemory)
	buyao := minCoinMemoryCache(coinSet, target, index+1, rest, dpMemory)

	if yao == -1 && buyao == -1 {
		dpMemory[index][rest] = -1
	} else if buyao == -1 {
		dpMemory[index][rest] = yao + 1 // 注意：必须+1，表示将yao这种情况要的硬币也加上
	} else if yao == -1 {
		dpMemory[index][rest] = buyao
	} else {
		if yao+1 < buyao { // 注意：必须是 yao+1 和 buyao 这两种情况进行比较
			dpMemory[index][rest] = yao + 1
		} else {
			dpMemory[index][rest] = buyao
		}
	}
	return dpMemory[index][rest]
}

func MinCoinStrictTable(coinSet []int, target int) int {
	if target < 1 {
		return -1
	}

	table := make([][]int, len(coinSet)+1) // 横坐标是coinSet中硬币的下标
	for i := 0; i < len(coinSet)+1; i++ {
		table[i] = make([]int, target+1) // 纵坐标是剩余钱数 (0~target)
	}

	for line := 0; line < len(coinSet)+1; line++ { // 第一列都为0，表示当剩余钱数为零时，任何硬币都不会被选
		table[line][0] = 0
	}
	for col := 1; col < target+1; col++ { // 最后一行(除了最后一行的第一列)都为-1，表示：剩余钱数不为0，但是已经没有硬币了，说明是失败分支
		table[len(coinSet)][col] = -1
	}

	for line := len(coinSet) - 1; line >= 0; line-- { // 从倒数第二行开始，向上计算
		for col := 1; col <= target; col++ {
			yao := -1
			if col-coinSet[line] >= 0 { // 第二维索引可能为负(硬币数值超额)
				yao = table[line+1][col-coinSet[line]] // 如果要当前硬币，那么当前结果取决于这个(左下方)
			}
			buyao := table[line+1][col] // 如果不要当前硬币，那么当前结果取决于这个(正下方)

			if yao == -1 && buyao == -1 {
				table[line][col] = -1
			} else if buyao == -1 {
				table[line][col] = yao + 1
			} else if yao == -1 {
				table[line][col] = buyao
			} else {
				if yao+1 > buyao { // 重要：这里是 yao+1 和 buyao 这两种情况选择小的（而不是yao和buyao进行比较）
					table[line][col] = buyao
				} else {
					table[line][col] = yao + 1 // 重要：一定要+1，作用是加上当前这个要的硬币
				}
			}
		}
	}
	return table[0][target] // 目标探测值(横坐标为0表示当前还没有遍历任何硬币，纵坐标为target表示当前剩余钱数没有任何减少)
}
