package BruteRecursion

import (
	"math"
)

type Commodity struct {
	price  int
	weight int
}

func MaxValue(commodities []Commodity, capacity int) int {
	return maxValue(commodities, capacity, 0, 0)
}

func maxValue(commodities []Commodity, capacity int, index int, load int) int {
	if load > capacity {
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
	if load > capacity {
		return 0
	}
	if index == len(commodities) {
		return profit
	}

	profit1 := maxValue2(commodities, capacity, index+1, load+commodities[index].weight, profit+commodities[index].price)
	profit2 := maxValue2(commodities, capacity, index+1, load, profit)

	return int(math.Max(float64(profit1), float64(profit2)))
}
