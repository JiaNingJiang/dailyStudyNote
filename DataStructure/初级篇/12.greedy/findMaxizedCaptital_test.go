package greedy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindMaxizedCaptital(t *testing.T) {

	count := 2   // 最多做几个项目
	capital := 9 // 启动资金

	costs := []int{2, 7, 12, 10}
	profits := []int{10, 8, 10, 4}

	fmt.Println("贪心算法，最大项目收益：", FindMaxizedCaptital(count, capital, profits, costs))
}

func comparator2(f1, f2 func(int, int, []int, []int) int) bool {
	var testTime int = 5000 // 比较次数
	var maxSize int = 5     // 测试用输入数组的最大长度
	var maxValue int = 15   // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		cost1, profit1 := generateRandomArray1(maxSize, maxValue) // 生成一个长度随机，元素值也随机的数组
		cost2 := make([]int, 0, len(cost1))
		profit2 := make([]int, 0, len(profit1))

		for j := 0; j < len(cost1); j++ {
			cost2 = append(cost2, cost1[j])
			profit2 = append(profit2, profit1[j])
		}

		rand.Seed(time.Now().UnixNano())
		count := len(cost1) - 2  // 最多项目数
		capital := rand.Intn(10) // 初始启动资金

		result1 := f1(count, capital, profit1, cost1)
		result2 := f2(count, capital, profit2, cost2)
		if result1 != result2 {
			succeed = false
			fmt.Println("启动资金:", capital, " 项目上限:", count)
			fmt.Printf("cost1: %v , cost2:%v\n", cost1, cost2)
			fmt.Printf("profit1: %v , profit2:%v\n", profit1, profit2)
			fmt.Printf("result1: %v , result2:%v\n", result1, result2)
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

func generateRandomArray1(maxSize, maxValue int) ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize)
	costs := make([]int, 0, length)
	profits := make([]int, 0, length)

	for i := 0; i < length; i++ {
		cost := rand.Intn(maxValue)
		profit := rand.Intn(maxValue)
		costs = append(costs, cost)
		profits = append(profits, profit)
	}

	return costs, profits
}
