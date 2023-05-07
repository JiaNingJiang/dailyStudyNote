package heapSort

import (
	"DataStructure/comparator"
	"fmt"
	"testing"
)

func TestSmallRootHeapSort(t *testing.T) {

	src := []int{54, 71, 12, 52, 32, 57, 68, 83}

	SmallRootHeapSort(src)

	fmt.Println("堆排序后：", src)
	comparator.Comparator(SmallRootHeapSort)
}
