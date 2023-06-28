package greedy

import (
	"DataStructure/heapSort"
	heap2 "container/heap"
)

func SlicingGoldBar(arr []int) int {
	h := heapSort.NewIntHeap(arr, false)
	heap2.Init(heapSort.NewIntHeap(arr, false)) // 将h变为一个小根堆

	newGoldBar := 0
	sumPrice := 0
	for {
		if h.Len() <= 1 {
			return sumPrice
		}
		newGoldBar = heap2.Pop(h).(int) + heap2.Pop(h).(int) // 弹出两个最小的元素，合成一个金块
		heap2.Push(h, newGoldBar)                            // 合成一个金块后，再次加入到小根堆中
		sumPrice += newGoldBar                               // 本次合并(或者说切割)花费的金额
	}
}
