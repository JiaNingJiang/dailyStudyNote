package lesson9

import (
	"fmt"
	"testing"
)

func TestLookUpMissing(t *testing.T) {
	// arr := []int{3, 2, 1, 6, 2, 7, 5}

	// arr := []int{3, 3, 2, 6, 2, 5, 5}

	// arr := []int{3, 1, 2, 6, 2, 5, 5}

	arr := []int{2, 2, 5, 5, 5, 5}
	fmt.Println("缺少的数字有: ", LookUpMissing(arr))

	arr = []int{2, 2, 5, 5, 5, 5}
	fmt.Println("缺少的数字有: ", matchAbsence(arr))
}

func matchAbsence(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for _, ele := range arr { // 争做让 每一个 i位置放的数字是 i+1
		modify(ele, arr)
	}
	absenceSet := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			absenceSet = append(absenceSet, i+1)
		}
	}
	return absenceSet
}

func modify(value int, arr []int) {
	for {
		if arr[value-1] == value {
			return
		}
		tmp := arr[value-1]
		arr[value-1] = value
		value = tmp
	}
}
