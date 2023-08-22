package bruteRecursionPromote

import (
	"fmt"
	"testing"
)

func TestCoinCase(t *testing.T) {

	arr := []int{5, 9, 4}
	target := 13

	fmt.Println("凑齐的可能性：", CoinCase(arr, target))
	fmt.Println("凑齐的可能性：", CoinCaseStrictTable(arr, target))

	Comparator1(CoinCase, CoinCaseStrictTable)

}
