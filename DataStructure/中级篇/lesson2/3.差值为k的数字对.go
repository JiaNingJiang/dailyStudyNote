package lesson2

import "fmt"

func NumberCouple(arr []int, dvalue int) {
	vSet := make(map[int]struct{})

	for _, v := range arr {
		if _, ok := vSet[v]; !ok {
			vSet[v] = struct{}{}
		}
	}

	for firV, _ := range vSet {
		if _, ok := vSet[firV+dvalue]; ok {
			fmt.Printf("数字对(%d - %d)\n", firV, firV+dvalue)
		}
	}
}
