package heapSort

// 值为int的大(小)根堆
type IntHeap struct {
	heap []int // 底层数组
	form bool  // true为大根堆  false为小根堆
}

// 创建一个大(小)根堆
func NewIntHeap(heap []int, form bool) *IntHeap {
	return &IntHeap{
		heap: heap,
		form: form,
	}
}

// 1.必须实现sort.Interface这个接口(包括less,len,swap三个方法)
func (ih *IntHeap) Len() int {
	return len(ih.heap)
}

// 决定何种情况下将下标为i的元素放到下标为j的元素的签名
func (ih *IntHeap) Less(i, j int) bool {
	if ih.form { // 大根堆
		return ih.heap[i] > ih.heap[j]
	} else { // 小根堆
		return ih.heap[i] < ih.heap[j]
	}
}
func (ih *IntHeap) Swap(i, j int) {
	ih.heap[i], ih.heap[j] = ih.heap[j], ih.heap[i]
}

// 2.实现Push方法(尾插法)
func (ih *IntHeap) Push(x interface{}) {
	ih.heap = append(ih.heap, x.(int))
}

// 3.实现Pop方法（弹出底层数组的末尾值(而非首部值)，让底层数组变为[0:n-1]）
func (ih *IntHeap) Pop() interface{} {
	old := ih
	n := len(old.heap)
	x := old.heap[n-1]
	ih.heap = old.heap[0 : n-1]
	return x
}

// 获取大(小)根堆的堆顶元素(注：这个方法不是heap接口规定要实现的)
func (ih *IntHeap) Top() interface{} {
	return ih.heap[0]
}
