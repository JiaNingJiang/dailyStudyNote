package lesson2

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMagicTime(t *testing.T) {
	arr1 := []int{95}
	arr2 := []int{32, 6, 64, 29, 18, 91, 26, 81, 0}
	fmt.Printf("最多可以进行(%d)次magic操作\n", MagicTime(arr1, arr2))

	arr1 = []int{95}
	arr2 = []int{32, 6, 64, 29, 18, 91, 26, 81, 0}
	fmt.Printf("最多可以进行(%d)次magic操作\n", MagicTimeRecursion(arr1, arr2))

	Comparator(MagicTime, MagicTimeRecursion)
}

func Comparator(f1, f2 func([]int, []int) int) bool {
	var testTime int = 50  // 比较次数
	var maxSize int = 12   // 测试用输入数组的最大大小
	var maxValue int = 100 // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue) // 生成一个长度随机，元素值也随机的数组
		arr2 := generateRandomArray(maxSize, maxValue)

		c_arr1 := make([]int, len(arr1))
		c_arr2 := make([]int, len(arr2))

		for index, val := range arr1 {
			c_arr1[index] = val
		}

		for index, val := range arr2 {
			c_arr2[index] = val
		}

		res1 := f1(arr1, arr2)
		res2 := f2(c_arr1, c_arr2)
		if !reflect.DeepEqual(res1, res2) {
			succeed = false
			fmt.Printf("arr1: %v   arr2:%v\n    res1: %v , res2:%v\n", arr1, arr2, res1, res2)
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
	time.Sleep(time.Nanosecond)
	length := rand.Intn(maxSize) + 1 //

	arr := make([]int, 0, length)

	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(maxValue))
	}
	return arr
}
