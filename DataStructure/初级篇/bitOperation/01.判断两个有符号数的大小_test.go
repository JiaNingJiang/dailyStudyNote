package bitOperation

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestGetMax1(t *testing.T) {
	fmt.Println(GetMax1(-191342107, 2059470033))
	Comparator(GetMax1, func(i1 int32, i2 int32) int32 {
		if i1 >= i2 {
			return i1
		} else {
			return i2
		}
	})

}

func TestGetMax2(t *testing.T) {
	fmt.Println(GetMax2(-191342107, 2059470033))
	Comparator(GetMax2, func(i1 int32, i2 int32) int32 {
		if i1 >= i2 {
			return i1
		} else {
			return i2
		}
	})
}

func Comparator(f1, f2 func(int32, int32) int32) bool {
	var testTime int = 50000 // 比较次数
	var maxValue int = 5000  // 测试用输入数组每个元素的大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxValue) // 生成一个长度随机，元素值也随机的数组

		res1 := f1(arr[0], arr[1])
		res2 := f2(arr[0], arr[1])
		if !reflect.DeepEqual(res1, res2) {
			succeed = false
			fmt.Printf("number1: %v , number2:%v , res1:%v , res2:%v\n", arr[0], arr[1], res1, res2)
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

func generateRandomArray(maxValue int) []int32 {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int32, 0, 2)

	for i := 0; i < 2; i++ {
		arr = append(arr, int32(rand.Intn(maxValue)))
	}
	return arr
}
