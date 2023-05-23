package greedy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBestArrage(t *testing.T) {
	meetings := []Meeting{
		{6, 8},
		{11, 15},
		{8, 10},
		{9, 12},
		{7, 9},
	}

	result := BestArrage(meetings, 6)
	fmt.Println("贪心算法结果：", result)

	resultForce := BestArrageForce(meetings, 6)
	fmt.Println("暴力递归算法结果：", resultForce)

	comparator1(BestArrage, BestArrageForce)
}
func comparator1(f1, f2 func([]Meeting, int) []Meeting) bool {
	var testTime int = 5000 // 比较次数
	var maxSize int = 20    // 测试用输入数组的最大大小
	var maxValue int = 24   // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue) // 生成一个长度为2，元素值也随机的数组
		arr2 := make([]Meeting, 0, len(arr1))

		for _, v := range arr1 {
			arr2 = append(arr2, v)
		}

		result1 := f1(arr1, 0)
		result2 := f2(arr2, 0)
		if len(result1) != len(result2) {
			succeed = false
			fmt.Printf("arr1: %v , arr2:%v\n", arr1, arr2)
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

func generateRandomArray(maxSize, maxValue int) []Meeting {
	rand.Seed(time.Now().UnixNano())
	len := rand.Intn(maxSize)
	arr := make([]Meeting, 0, len)

	for i := 0; i < len; i++ {
		startTime := rand.Intn(maxValue)              // 区间开始时间(0h~23h)
		interval := 1 + rand.Intn(maxValue-startTime) // 每次占用的最长使用时间（1h ~ 24h）
		meeting := Meeting{
			Start: startTime,
			End:   startTime + interval,
		}
		arr = append(arr, meeting)
	}

	return arr
}
