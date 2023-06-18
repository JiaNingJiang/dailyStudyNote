package bruteRecursionPromote

import "math"

// 同一面值硬币可以重复使用
func MinCoinSP(coins []int, target int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}
	if target <= 0 {
		return -1
	}
	minCoinCount := math.MaxInt
	minCoinSP(coins, &minCoinCount, target, 0)

	if minCoinCount == math.MaxInt {
		minCoinCount = -1
	}

	return minCoinCount
}

func minCoinSP(coins []int, minCoinCount *int, res int, count int) {
	if res < 0 {
		return
	}
	if res == 0 {
		*minCoinCount = getMin(*minCoinCount, count)
	}

	for i := 0; i < len(coins); i++ { // 每次都只挑选一种面值的硬币，每一种都有被选的可能
		minCoinSP(coins, minCoinCount, res-coins[i], count+1) // 不考虑任何策略，只是给出所有的可能性
	}
}

func MinCoinSPMemoryCache(coins []int, target int) int {
	if coins == nil || len(coins) == 0 {
		return -1
	}
	if target <= 0 {
		return -1
	}
	dpMemory := make([]int, target) // dpMemory[rest] 总是存储 凑齐rest金额所需要的最少硬币数

	return minCoinSPMemoryCache(coins, target, dpMemory)
}

func minCoinSPMemoryCache(coins []int, res int, dpMemory []int) int {
	if res == 0 { // 刚好凑齐
		return 0
	}
	if res < 0 { // 超额
		return -1
	}
	if dpMemory[res-1] != 0 { // -1 的意义是因为数组的下标不是从1开始，而是从0开始。因此rest与下标的关系就是 rest-1 == index
		return dpMemory[res-1]
	}
	minCoinCount := math.MaxInt // 记录当前深度(同一个res)下所需要的最少硬币数
	for i := 0; i < len(coins); i++ {
		result := minCoinSPMemoryCache(coins, res-coins[i], dpMemory)
		if result >= 0 && result+1 < minCoinCount {
			minCoinCount = result + 1 // +1 的意义是加上当前使用的这个硬币，即 coins[i]
		}
	}
	if minCoinCount == math.MaxInt { // 所有的分支都不能凑齐rest，那么意味着当剩余钱数为rest时是无法从凑齐的
		dpMemory[res-1] = -1
	} else {
		dpMemory[res-1] = minCoinCount // 记录所有分支中需要硬币数最少的
	}
	return dpMemory[res-1]
}

func MinCoinSPStrictTable(coins []int, target int) int {
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

	// 当前元素从 自己 i张分支中获取分支的返回值(各分支所需的硬币数) + 自己的张数，从分支的硬币数中选择一个最小的进行保存
	table[len(coins)][0] = 0 // 最后一行只有第一个元素可以不拿任何硬币凑齐target，其他的元素都凑不齐target
	for col := 1; col <= target; col++ {
		table[len(coins)][col] = -1
	}

	for line := len(coins) - 1; line >= 0; line-- { // 从倒数第二行向上计算，因为上层的结果依赖于下层
		for col := 0; col <= target; col++ {
			minCoinCount := math.MaxInt
			for zhang := 0; coins[line]*zhang <= col; zhang++ {
				if table[line+1][col-coins[line]*zhang] != -1 { // 下面的分支能够凑齐
					count := table[line+1][col-coins[line]*zhang] + zhang // 下层所需的硬币数+当前层使用的相同面值的硬币数
					if count < minCoinCount {
						minCoinCount = count // 保存本层所有分支中需要的最少硬币数目
					}
				}
			}
			if minCoinCount == math.MaxInt {
				table[line][col] = -1
			} else {
				table[line][col] = minCoinCount
			}

		}
	}

	return table[0][target]
}
