package heapSort

import (
	"fmt"
	"testing"
)

func TestRewriteAndRecovery(t *testing.T) {

	src := []int{54, 71, 12, 52, 32, 57, 68, 83} // 原始数组
	heap := NewHeap(src, false)                  // 变为小根堆
	fmt.Println("破坏前小根堆：", heap)
	RewriteAndRecovery(heap, 3, 11, false) // 破坏并恢复

	fmt.Println("破坏并恢复后：", heap)

	//GetSortArrFromHeap(heap, false) //
	//
	//fmt.Println("破坏并恢复后：", heap)
}
