package bruteRecursionPromote

import (
	"fmt"
	"testing"
)

func TestMinCoinSP(t *testing.T) {
	arr := []int{5, 9, 4}
	target := 13

	fmt.Println("最少的硬币数：", MinCoinSP(arr, target))
	fmt.Println("最少的硬币数：", MinCoinSPMemoryCache(arr, target))
	fmt.Println("最少的硬币数：", MinCoinSPStrictTable(arr, target))

	Comparator1(MinCoinSP, MinCoinSPMemoryCache)
	fmt.Println()
	Comparator1(MinCoinSPMemoryCache, MinCoinSPStrictTable)
}
