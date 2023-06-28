package lesson1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSheepEatGrass(t *testing.T) {
	grass := 20

	fmt.Println("胜利者: ", SheepEatGrass(grass))

	Comparator2(SheepEatGrass, SheepEatGrass2)
}

func Comparator2(f1, f2 func(int) string) bool {
	var testTime int = 70 // 比较次数
	//var maxValue int = 100
	var succeed bool = true

	for i := 0; i < testTime; i++ {
		grass := i
		winner1 := f1(grass)
		winner2 := f2(grass)
		if !reflect.DeepEqual(winner1, winner2) {
			succeed = false
			fmt.Printf("winner1: %v , winner2:%v\n", winner1, winner2)
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
