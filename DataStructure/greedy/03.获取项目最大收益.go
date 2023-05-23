package greedy

import (
	heap2 "container/heap"
)

type Project struct {
	Cost   int // 消费
	Profit int // 利润
}

// k表示你最多只能做k个项目 .  m表示你的初始资金
// profits[i] 表示i号项目在做完之后能够得到的利润(净收益)
// costs[i] 表示i号项目的花费
// 返回值是最后的所得的利润
func FindMaxizedCaptital(k, m int, Profits []int, Costs []int) int {

	costHeap := NewProjectHeap([]project{}, false) // 按照花费大小进行排序的小根堆
	heap2.Init(costHeap)
	profitHeap := NewProjectHeap([]project{}, true) // 按照利润大小排序的大根堆
	heap2.Init(profitHeap)

	for i := 0; i < len(Costs); i++ {
		p := project{cost: Costs[i], profit: Profits[i]}
		heap2.Push(costHeap, p)
	}

	current := m             // 当前可供支配的资金
	for i := 0; i < k; i++ { // 最多寻找k个项目
		for { // 取出当前有能力做的所有项目，添加到按照利润排序的大根堆profitHeap中
			if costHeap.Len() == 0 {
				break
			}
			topProject := heap2.Pop(costHeap).(project) // 弹出一个花费最小的项目
			if topProject.cost > current {              // 没有项目可做了
				heap2.Push(costHeap, topProject) // 再重新回到小根堆中
				break
			}
			heap2.Push(profitHeap, topProject) // 将这个项目加入到按照利润排序的大根堆中
		}
		if profitHeap.Len() == 0 { // 当前没有能力做任何项目了
			return current
		}
		current += heap2.Pop(profitHeap).(project).profit // 做完一个最润最大的项目
	}
	return current
}

type project struct {
	cost   int
	profit int
}

// 值为project的大(小)根堆
type projectHeap struct {
	heap []project // 底层数组
	form bool      // true为按利润排序的大根堆  false为按花费排序的小根堆
}

// 创建一个大(小)根堆
func NewProjectHeap(heap []project, form bool) *projectHeap {
	return &projectHeap{
		heap: heap,
		form: form,
	}
}

// 必须实现sort.Interface这个接口(包括less,len,swap三个方法)
func (ph *projectHeap) Len() int {
	return len(ph.heap)
}
func (ph *projectHeap) Less(i, j int) bool {
	if ph.form { // 大根堆
		return ph.heap[i].profit > ph.heap[j].profit
	} else { // 小根堆
		return ph.heap[i].cost < ph.heap[j].cost
	}
}
func (ph *projectHeap) Swap(i, j int) {
	ph.heap[i], ph.heap[j] = ph.heap[j], ph.heap[i]
}

// 实现Push方法
func (ph *projectHeap) Push(x interface{}) {
	ph.heap = append(ph.heap, x.(project))
}

// 实现Pop方法（弹出底层数组的末尾值，让底层数组变为[0:n-1]）
func (ph *projectHeap) Pop() interface{} {
	old := ph
	n := len(old.heap)
	x := old.heap[n-1]
	ph.heap = old.heap[0 : n-1]
	return x
}

// 获取大(小)根堆的堆顶元素(注：这个方法不是heap接口规定要实现的)
func (ph *projectHeap) Top() interface{} {
	return ph.heap[0]
}
