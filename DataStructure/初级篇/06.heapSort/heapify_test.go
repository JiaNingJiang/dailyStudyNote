package heapSort

import (
	"fmt"
	"testing"
)

func TestGetSortArrFromHeap(t *testing.T) {
	src := []int{14, 12, 10, 8, 7, 6}
	form := false
	arr := NewHeap(src, form)
	fmt.Println("heapInsert之后: ", arr)
	GetSortArrFromHeap(arr, form)
	fmt.Println("堆排序后：", arr)

}
