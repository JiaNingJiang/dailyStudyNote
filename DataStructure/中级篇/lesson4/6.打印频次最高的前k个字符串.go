package lesson4

import "DataStructure2/utils"

type Word struct {
	Str   string // 单词字符串
	Count int    // 单词出现的次数
}

func LessWord(a, b interface{}) bool {
	aWord := a.(Word)
	bWord := b.(Word)

	if aWord.Count < bWord.Count {
		return true
	} else {
		return false
	}
}

type DymMostFreq struct {
	FreqTable     map[string]int // 词频表
	SmallRootHeap []interface{}  // 小根堆
	SRHeapLoc     map[string]int // 记录单词在小根堆中的位置

	K int // 只打印出现次数最多的k个单词
}

func NewDymMostFreq(k int) *DymMostFreq {
	return &DymMostFreq{
		FreqTable:     make(map[string]int),
		SmallRootHeap: utils.NewHeap(make([]interface{}, 0, k), false, LessWord),
		SRHeapLoc:     make(map[string]int),
		K:             k,
	}
}

// 小根堆完成一次heapify后,需要更新SRHeapLoc
func Hook(srHeapLoc *map[string]int, a, b *interface{}) {
	aWord := (*a).(Word)
	bWord := (*b).(Word)
	temp := (*srHeapLoc)[aWord.Str]
	(*srHeapLoc)[aWord.Str] = (*srHeapLoc)[bWord.Str]
	(*srHeapLoc)[bWord.Str] = temp
}

func (dmf *DymMostFreq) AddWord(word string) *DymMostFreq {
	if _, ok := dmf.FreqTable[word]; ok {
		dmf.FreqTable[word]++
	} else {
		dmf.FreqTable[word] = 1
	}

	if loc, ok := dmf.SRHeapLoc[word]; ok && loc != -1 { // 单词存在于小根堆上. 更新节点的出现次数后重新heapify
		node := dmf.SmallRootHeap[loc].(Word)
		node.Count++
		dmf.SmallRootHeap[loc] = node
		Heapify(dmf.SmallRootHeap, 0, len(dmf.SmallRootHeap), false, LessWord, Hook, &dmf.SRHeapLoc) // 需要重新调整SRHeapLoc
	} else { // 单词不在小根堆上
		if len(dmf.SmallRootHeap) < dmf.K { // 小根堆未满
			newLoc := 0
			newLoc = utils.HeapInsert(&dmf.SmallRootHeap, Word{Str: word, Count: 1}, false, LessWord)
			dmf.SRHeapLoc[word] = newLoc
		} else { // 小根堆已经满了，需要看堆顶元素是否可以被替换掉
			topCount := dmf.SmallRootHeap[0].(Word).Count
			if dmf.FreqTable[word] <= topCount {
				dmf.SRHeapLoc[word] = -1
			} else {
				dmf.SRHeapLoc[dmf.SmallRootHeap[0].(Word).Str] = -1
				dmf.SmallRootHeap[0] = Word{word, dmf.FreqTable[word]}
				dmf.SRHeapLoc[word] = 0
				Heapify(dmf.SmallRootHeap, 0, len(dmf.SmallRootHeap), false, LessWord, Hook, &dmf.SRHeapLoc)
			}
		}
	}
	return dmf
}

func (dmf *DymMostFreq) Pop() string {
	var data interface{}
	dmf.SmallRootHeap, data = PopAndheapify(dmf.SmallRootHeap, false, LessWord, Hook, &dmf.SRHeapLoc)
	word := data.(Word)

	dmf.FreqTable[word.Str]--
	dmf.SRHeapLoc[word.Str] = -1

	return word.Str
}

func Heapify(heap []interface{}, start, end int, form bool, less func(interface{}, interface{}) bool,
	hook func(*map[string]int, *interface{}, *interface{}), srHeapLoc *map[string]int) {
	currentIndex := start // 对以 heap[start]为根节点的子树进行heapify
	leftChildIndex := currentIndex*2 + 1
	rightChildIndex := currentIndex*2 + 2

	for {
		if leftChildIndex > end { // 没有任何孩子节点(最多到heap[end-1])
			break
		}
		newRootIndex := currentIndex
		if form { // 大根堆
			// 获得当前节点 左右孩子 中较大节点的下标
			maxIndex := leftChildIndex
			max := heap[leftChildIndex]
			if rightChildIndex <= len(heap)-1 { // 可能只有左孩子，没有右孩子
				max = BackMax(heap[leftChildIndex], heap[rightChildIndex], less)
				if max == heap[rightChildIndex] {
					maxIndex = rightChildIndex
				}
			}
			// 如果较大节点比根节点还要大，则交换两者
			if less(heap[currentIndex], max) {
				swap(&heap[maxIndex], &heap[currentIndex], hook, srHeapLoc)
			}
			newRootIndex = maxIndex
		} else { // 小根堆
			// 获得当前节点 左右孩子 中较小节点的下标
			minIndex := leftChildIndex
			min := heap[leftChildIndex]
			if rightChildIndex <= len(heap)-1 { // 可能只有左孩子，没有右孩子
				min = BackMin(heap[leftChildIndex], heap[rightChildIndex], less)
				if min == heap[rightChildIndex] {
					minIndex = rightChildIndex
				}
			}
			// 如果较小节点比根节点还要小，则交换两者
			if less(min, heap[currentIndex]) {
				swap(&heap[minIndex], &heap[currentIndex], hook, srHeapLoc)
			}
			newRootIndex = minIndex
		}
		// 更新循环变量
		currentIndex = newRootIndex
		leftChildIndex = currentIndex*2 + 1
		rightChildIndex = currentIndex*2 + 2
	}
}

func PopAndheapify(heap []interface{}, form bool, less func(interface{}, interface{}) bool,
	hook func(*map[string]int, *interface{}, *interface{}), srHeapLoc *map[string]int) ([]interface{}, interface{}) {

	num := heap[0]              // 每次总是返回根堆的根节点
	heap[0] = heap[len(heap)-1] // 让末尾的叶子结点替换掉根节点
	heap = heap[:len(heap)-1]

	// 将新的根节点下沉到合适的位置
	Heapify(heap, 0, len(heap)-1, form, less, hook, srHeapLoc) // 注意：end必须是heapIndex - 1，作用是相当于heap[heapIndex]被删除
	return heap, num
}

func BackMax(a, b interface{}, less func(interface{}, interface{}) bool) interface{} {
	if less(a, b) {
		return b
	} else {
		return a
	}
}

func BackMin(a, b interface{}, less func(interface{}, interface{}) bool) interface{} {
	if less(a, b) {
		return a
	} else {
		return b
	}
}

func swap(a, b *interface{}, hook func(*map[string]int, *interface{}, *interface{}), srHeapLoc *map[string]int) {

	hook(srHeapLoc, a, b)
	temp := *a
	*a = *b
	*b = temp
}
