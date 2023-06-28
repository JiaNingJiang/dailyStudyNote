package lesson1

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMinPain(t *testing.T) {
	arr := []string{"R", "G", "R", "G", "R"}

	fmt.Println("染色的最小代价为: ", MinPain(arr))

	Comparator3(MinPain, MinPain2)
}

func Comparator3(f1, f2 func([]string) int) bool {
	var testTime int = 50 // 比较次数
	var maxSize int = 15  // 测试用输入数组的最大大小
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		colorArr := generateColorLattice(maxSize)

		cost1 := f1(colorArr)
		cost2 := f2(colorArr)
		if !reflect.DeepEqual(cost1, cost2) {
			succeed = false
			fmt.Printf("cost1: %v , cost2:%v\n", cost1, cost2)
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

func generateColorLattice(maxSize int) []string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize) //

	arr := make([]string, 0, length)

	for i := 0; i < length; i++ {
		time.Sleep(time.Nanosecond)
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(2)
		if random == 0 {
			arr = append(arr, "R")
		} else {
			arr = append(arr, "G")
		}
	}
	return arr
}
