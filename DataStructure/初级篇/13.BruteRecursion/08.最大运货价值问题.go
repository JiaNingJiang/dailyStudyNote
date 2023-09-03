package BruteRecursion

import (
	"math"
)

type Commodity struct {
	price  int
	weight int
}

// 分类递归的问题，一旦发现异常分支(不符合给定条件的分支)，则必须想办法将该分支丢弃掉。比如通过在选优阶段降低该分支的优先度到极低而实现
// 正常分支与异常分支的另一个不同之处：正常分支能走到结束条件(结束条件通常是完成了对数组的遍历)，异常分支不允许走到结束条件

func MaxValue(commodities []Commodity, capacity int) int {
	return maxValue(commodities, capacity, 0, 0)
}

func maxValue(commodities []Commodity, capacity int, index int, load int) int {
	if load > capacity { // 一旦发现超重,立即丢弃该分支
		// 一旦发现要了上一笔货物而超重,则立即返回，导致 profit1 = 0.则上一个货物肯定不能要，profit1分支结束递归，只保留profit2分支继续
		return -commodities[index-1].price
	}
	if index == len(commodities) {
		return 0
	}

	profit1 := maxValue(commodities, capacity, index+1, load+commodities[index].weight) + commodities[index].price
	profit2 := maxValue(commodities, capacity, index+1, load)

	return int(math.Max(float64(profit1), float64(profit2)))
}

func MaxValue2(commodities []Commodity, capacity int) int {
	return maxValue2(commodities, capacity, 0, 0, 0)
}
func maxValue2(commodities []Commodity, capacity int, index int, load int, profit int) int {
	if load > capacity { // 一旦发现超重,立即丢弃该分支
		return 0
	}
	if index == len(commodities) {
		return profit
	}

	profit1 := maxValue2(commodities, capacity, index+1, load+commodities[index].weight, profit+commodities[index].price)
	profit2 := maxValue2(commodities, capacity, index+1, load, profit)

	return int(math.Max(float64(profit1), float64(profit2)))
}
