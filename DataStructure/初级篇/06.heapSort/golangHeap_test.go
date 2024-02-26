package heapSort

import (
	heap2 "container/heap"
	"fmt"
	"testing"
)

func TestGolangHeap(t *testing.T) {
	src := []int{14, 12, 10, 8, 7, 6}
	form := true

	heap := NewIntHeap(src, form)
	heap2.Init(heap) // 初始化为大(小)根堆

	fmt.Println("heapInsert之后: ", heap.heap)

	for {
		if heap.Len() == 0 {
			break
		}
		fmt.Printf("%d ", heap2.Pop(heap))
	}

}
