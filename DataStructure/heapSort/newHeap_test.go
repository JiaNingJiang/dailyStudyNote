package heapSort

import (
	"fmt"
	"testing"
)

func TestFastNewHeap(t *testing.T) {
	src := []int{6, 8, 7, 10, 12, 14}
	FastNewHeap(src, false)

	fmt.Println("heapInsert之后: ", src)
}
