package bruteRecursionPromote

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMinCoin(t *testing.T) {
	arr := []int{4, 5, 10, 3}
	//arr := []int{10, 5, 10, 10, 10, 10, 10}
	target := 13

	fmt.Println("最少的硬币数：", MinCoin(arr, target))

	fmt.Println("最少的硬币数：", MinCoinMemoryCache(arr, target))

	fmt.Println("最少的硬币数：", MinCoinStrictTable(arr, target))

	Comparator1(MinCoin, MinCoinMemoryCache)
	fmt.Println()

	Comparator1(MinCoin, MinCoinStrictTable)
}

func Comparator1(f1, f2 func([]int, int) int) bool {
	var testTime int = 5000 // 比较次数
	var maxCoinNumber int = 20
	var maxCoinValue int = 15
	var maxTargetValue int = 15

	var succeed bool = true

	for i := 0; i < testTime; i++ {
		coinSet, target := generateRandomArray1(maxCoinNumber, maxCoinValue, maxTargetValue)

		coinSetCopy := make([]int, 0)
		for j := 0; j < len(coinSet); j++ {
			coinSetCopy = append(coinSetCopy, coinSet[j])
		}

		cointSet1 := f1(coinSet, target)
		cointSet2 := f2(coinSetCopy, target)
		if !reflect.DeepEqual(cointSet1, cointSet2) {
			succeed = false
			fmt.Printf("coinSet:(%v) target:(%v) res1:(%v)   res2:(%v)\n", coinSet, target, cointSet1, cointSet2)
			break
		}
	}

	if succeed {
		fmt.Printf("insertionSort run successfully!")
		return true
	} else {
		fmt.Printf("insertionSort run Faultily!")
		return false
	}
}

func generateRandomArray1(maxCoinNumber, maxCoinValue, maxTargetValue int) ([]int, int) {
	rand.Seed(time.Now().UnixNano())
	coinNum := rand.Intn(maxCoinNumber)

	coinSet := make([]int, coinNum)
	for i := 0; i < coinNum; i++ {
		coinSet[i] = rand.Intn(maxCoinValue) + 1

	}

	target := rand.Intn(maxTargetValue)

	return coinSet, target
}
