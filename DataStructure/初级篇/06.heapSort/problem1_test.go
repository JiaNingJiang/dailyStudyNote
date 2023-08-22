package heapSort

import (
	"fmt"
	"testing"
)

func TestProcess(t *testing.T) {
	arr := []int{4, 1, 3, 2}
	k := 3
	Process(arr, k)

	fmt.Println("排序后: ", arr)
}
