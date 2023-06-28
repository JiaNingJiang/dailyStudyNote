package lesson1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinBags(t *testing.T) {
	Comparator(MinBags, minBags2)
}

func Comparator(f1, f2 func(int) int) bool {
	var testTime int = 5000 // 比较次数
	//var maxValue int = 100
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		//rand.Seed(time.Now().UnixNano())
		//time.Sleep(time.Nanosecond)
		//apples := rand.Intn(maxValue)
		apples := i
		bags1 := f1(apples)
		bags2 := f2(apples)
		if !reflect.DeepEqual(bags1, bags2) {
			succeed = false
			fmt.Printf("arr1: %v , arr2:%v\n", bags1, bags2)
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
