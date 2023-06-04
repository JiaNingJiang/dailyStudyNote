package bruteRecursionPromote

import (
	"fmt"
	"testing"
)

func TestMinCoin(t *testing.T) {
	arr := []int{2, 3, 5, 9, 4, 7}
	target := 10

	fmt.Println("最少的硬币数：", MinCoin(arr, target))

	fmt.Println("最少的硬币数：", MinCoinMemoryCache(arr, target))
}
