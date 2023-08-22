package greedy

import (
	"DataStructure/06.heapSort"
	heap2 "container/heap"
)

func InitHeap() (*heapSort.IntHeap, *heapSort.IntHeap) {
	smallBase := make([]int, 0)
	bigBase := make([]int, 0)

	smallHeap := heapSort.NewIntHeap(smallBase, false) // 小根堆后N/2个数
	bigHeap := heapSort.NewIntHeap(bigBase, true)      // 大根堆,存储前N/2个数

	heap2.Init(smallHeap)
	heap2.Init(bigHeap)

	return smallHeap, bigHeap
}

// 每次从数据流中取出一个数,并返回当前已有数字的中位数
func FindMedian(smallHeap, bigHeap *heapSort.IntHeap, num int) float64 {
	if smallHeap.Len() == 0 && bigHeap.Len() == 0 { // 第一个数加入到大根堆中
		heap2.Push(bigHeap, num)
		return float64(num)
	}
	bigTop := heap2.Pop(bigHeap).(int) // 获取大根堆堆顶元素(大根堆永远不可能再次为空，因此不需要判断堆是否为空)
	heap2.Push(bigHeap, bigTop)
	if num >= bigTop { // 放入到小根堆(作为后N/2个数)
		heap2.Push(smallHeap, num)
	} else { // 放入小根堆(作为前N/2个数)
		heap2.Push(bigHeap, num)
	}

	diff := smallHeap.Len() - bigHeap.Len()
	sum := smallHeap.Len() + bigHeap.Len()

	if diff >= 2 {
		moveNum := heap2.Pop(smallHeap)
		heap2.Push(bigHeap, moveNum)
	} else if diff <= -2 {
		moveNum := heap2.Pop(bigHeap)
		heap2.Push(smallHeap, moveNum)
	}

	// 重点：注意各种情况的判别，返回正确的中位数
	if sum%2 != 0 { // 奇数个数时，有唯一的中位数
		if diff > 0 { // 小根堆数多，那么中位数就是小根堆的堆顶元素
			median := heap2.Pop(smallHeap).(int) // 获取小根堆堆顶元素,此数即为中位数
			heap2.Push(smallHeap, median)
			return float64(median)
		} else { // 大根堆数多，那么中位数就是大根堆的堆顶元素
			median := heap2.Pop(bigHeap).(int) // 获取大根堆堆顶元素,此数即为中位数
			heap2.Push(bigHeap, median)
			return float64(median)
		}
	} else { // 偶数个数时，中位数为平均值
		median1 := heap2.Pop(smallHeap).(int) // 获取小根堆堆顶元素,此数即为中位数
		heap2.Push(smallHeap, median1)

		median2 := heap2.Pop(bigHeap).(int) // 获取大根堆堆顶元素,此数即为中位数
		heap2.Push(bigHeap, median2)

		return float64((median1 + median2)) / 2
	}

}
