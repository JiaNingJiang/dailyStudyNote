package heapSort

import (
	"fmt"
	"testing"
)

func TestNewHeap(t *testing.T) {
	src := []int{6, 8, 7, 10, 12, 14}
	arr := NewHeap(src, false)

	fmt.Println("heapInsert之后: ", arr)
}
