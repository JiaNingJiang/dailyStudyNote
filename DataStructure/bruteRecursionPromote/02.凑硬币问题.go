package bruteRecursionPromote

func MinCoin(coinSet []int, target int) int {
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
	dpMemory := make([][]int, len(coinSet)+1) //
	for i := 0; i < len(coinSet)+1; i++ {     // 横坐标表示当前本轮选中的硬币(因为可能为0，因此要+1)
		dpMemory[i] = make([]int, target+1) // 纵坐标表示当前剩余需要的钱数(必然在 0 ~ target范围内)，不能使用已搜集钱数作为纵坐标(因为范围不止0~target)
		for j := 0; j <= target; j++ {
			dpMemory[i][j] = -1
		}
	}
	return minCoinMemoryCache(coinSet, target, 0, 0, target, dpMemory)
}

func minCoinMemoryCache(coinSet []int, target int, index int, curCoin, rest int, dpMemory [][]int) int {
	if index >= len(coinSet) && rest != 0 { // coinSet访问越界(先返回，防止dpMemory的横坐标越界)  （rest!=0是为了防止在最后一个硬币出收集齐target数额）
		return -1
	}
	if rest < 0 { // 金币超额(先返回，防止dpMemory的纵坐标越界)
		return -1
	}
	if dpMemory[index][rest] != -1 { // 缓存命中
		return dpMemory[index][rest]
	}
	// 缓存未命中，也就是新情况
	if rest == 0 { // 金币刚好够
		dpMemory[index][0] = curCoin // 记录缓存(当前路径硬币数)
	} else { // 金币不够(即rest > 0)
		yao := minCoinMemoryCache(coinSet, target, index+1, curCoin+1, rest-coinSet[index], dpMemory)
		buyao := minCoinMemoryCache(coinSet, target, index+1, curCoin, rest, dpMemory)

		if yao == -1 && buyao == -1 {
			dpMemory[index][rest] = -1
		} else if buyao == -1 {
			dpMemory[index][rest] = yao
		} else if yao == -1 {
			dpMemory[index][rest] = buyao
		} else {
			if yao < buyao {
				dpMemory[index][rest] = yao
			} else {
				dpMemory[index][rest] = buyao
			}
		}
	}
	return dpMemory[index][rest]
}
