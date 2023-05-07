package comparator

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// 参数1：需要进行测试的排序函数
func Comparator(f func([]int)) bool {
	var testTime int = 50000 // 比较次数
	var maxSize int = 100    // 测试用输入数组的最大大小
	var maxValue int = 100   // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue) // 生成一个长度随机，元素值也随机的数组
		arr2 := make([]int, 0, len(arr1))

		for _, v := range arr1 {
			arr2 = append(arr2, v)
		}

		f(arr1)
		sort.Ints(arr2)
		if !reflect.DeepEqual(arr1, arr2) {
			succeed = false
			fmt.Printf("arr1: %v , arr2:%v\n", arr1, arr2)
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

func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	len := rand.Intn(maxSize) //

	arr := make([]int, 0, len)

	for i := 0; i < len; i++ {
		arr = append(arr, rand.Intn(maxValue))
	}
	return arr
}
