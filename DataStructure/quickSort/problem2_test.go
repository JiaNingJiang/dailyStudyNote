package quickSort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestDutchFlag(t *testing.T) {
	arr := []int{49, 61, 23, 10, 1, 1, 1}
	target := 23
	//arr := []int{2, 4, 4, 3, 7, 1, 4, 6}
	//target := 4
	DutchFlag(arr, target)
	//force(arr, target)
	fmt.Println("排序后：", arr)

	Comparator(DutchFlag, force)
}

func force(arr []int, target int) {
	temp := make([]int, 0)
	for _, v := range arr { // 第一个循环，把所有小于target的存到temp
		if v < target {
			temp = append(temp, v)
		}
	}
	for _, v := range arr { // 第二个循环，把所有等于target的存到temp
		if v == target {
			temp = append(temp, v)
		}
	}
	for _, v := range arr { // 第三个循环，把所有大于target的存到temp
		if v > target {
			temp = append(temp, v)
		}
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = temp[i]
	}
}

func Comparator(f1, f2 func([]int, int)) bool {
	var testTime int = 5000 // 比较次数
	var maxSize int = 8     // 测试用输入数组的最大大小
	var maxValue int = 100  // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue) // 生成一个长度随机，元素值也随机的数组
		randomArray := make([]int, 0, len(arr1))
		arr2 := make([]int, 0, len(arr1))

		for _, v := range arr1 {
			arr2 = append(arr2, v)
			randomArray = append(randomArray, v)
		}

		f1(arr1, arr1[len(arr1)/2])
		f2(arr2, arr2[len(arr2)/2])
		if !reflect.DeepEqual(arr1, arr2) {
			succeed = false
			fmt.Printf("arr1: %v , arr2:%v \n randomArray:%v ,Target:%v\n",
				arr1, arr2, randomArray, arr1[len(arr1)/2])
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
