package greedy

import (
	"fmt"
	"testing"
)

func TestFindMedian(t *testing.T) {

	smallHeap, bigHeap := InitHeap()

	fmt.Println("当前中位数：", FindMedian(smallHeap, bigHeap, 1))

	fmt.Println("当前中位数：", FindMedian(smallHeap, bigHeap, 2))

	fmt.Println("当前中位数：", FindMedian(smallHeap, bigHeap, 3))

	fmt.Println("当前中位数：", FindMedian(smallHeap, bigHeap, 4))

	fmt.Println("当前中位数：", FindMedian(smallHeap, bigHeap, 5))
}
