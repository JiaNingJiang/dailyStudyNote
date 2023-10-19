package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 2, 2, 5}

	index := sort.SearchInts(arr, 3)
	arr[index] = 3

	fmt.Println(arr)
}
