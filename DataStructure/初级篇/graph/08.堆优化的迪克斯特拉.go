package graph

func Dijkstral2(node *Node) map[*Node]int {
	nodeHeap := NewNodeHeap()

	nodeHeap.addOrUpdateOrIgnore(node, 0)
	result := make(map[*Node]int)
	for {
		if nodeHeap.isEmpty() {
			break
		}
		record := nodeHeap.pop() // 小根堆出根节点
		cur := record.Node
		distance := record.Distance
		for _, edge := range cur.edge {
			nodeHeap.addOrUpdateOrIgnore(edge.to, edge.weight+distance)
		}
		result[cur] = distance // 遍历过的节点放入result中（从小根堆中弹出的节点不可能再回到小根堆中）
	}
	return result
}

// 专用的小根堆
type NodeHeap struct {
	nodes        []*Node       // 底层数组
	heapIndexMap map[*Node]int // 可以根据Node查询其在数组中的位置(对于那些曾经进入过小根堆，但后续又出堆的节点其对应的int值为-1)
	distanceMap  map[*Node]int // Node与head的当前最小距离
	size         int           // 堆节点的数量
}

func NewNodeHeap() *NodeHeap {
	nodeHeap := new(NodeHeap)
	nodeHeap.nodes = make([]*Node, 0)
	nodeHeap.heapIndexMap = make(map[*Node]int)
	nodeHeap.distanceMap = make(map[*Node]int)
	nodeHeap.size = 0
	return nodeHeap
}

// 判断小根堆是否为空
func (nh *NodeHeap) isEmpty() bool {
	return nh.size == 0
}

// 判断Node是否曾经进入过小根堆
func (nh *NodeHeap) isEntered(node *Node) bool {
	if _, ok := nh.heapIndexMap[node]; ok {
		return true
	} else {
		return false
	}
}

// 判断Node当前是否还在小根堆中
func (nh *NodeHeap) inHeap(node *Node) bool {
	return nh.isEntered(node) && nh.heapIndexMap[node] != -1
}

// 在小根堆上交换两个节点
func (nh *NodeHeap) swap(index1, index2 int) {
	// 1.交换历史存在位置
	nh.heapIndexMap[nh.nodes[index1]] = index2
	nh.heapIndexMap[nh.nodes[index2]] = index1

	// 2.交换当前小根堆底层数组中的位置
	tmp := nh.nodes[index1]
	nh.nodes[index1] = nh.nodes[index2]
	nh.nodes[index2] = tmp
}

func (nh *NodeHeap) heapify(end int) {
	currentIndex := 0
	leftChildIndex := currentIndex*2 + 1
	rightChildIndex := currentIndex*2 + 2

	for {
		if leftChildIndex > end { // 没有任何孩子节点(最多到heap[end-1])
			break
		}
		newRootIndex := currentIndex

		// 获得当前节点 左右孩子 中较小节点的下标
		minIndex := leftChildIndex
		minDis := nh.distanceMap[nh.nodes[leftChildIndex]]
		if rightChildIndex <= len(nh.nodes)-1 { // 可能只有左孩子，没有右孩子
			minDis = getMin(nh.distanceMap[nh.nodes[leftChildIndex]], nh.distanceMap[nh.nodes[rightChildIndex]])
			if minDis == nh.distanceMap[nh.nodes[rightChildIndex]] {
				minIndex = rightChildIndex
			}
		}
		// 如果较小节点比根节点还要小，则交换两者
		if minDis < nh.distanceMap[nh.nodes[currentIndex]] {
			nh.swap(minIndex, currentIndex)
		}
		newRootIndex = minIndex

		// 更新循环变量
		currentIndex = newRootIndex
		leftChildIndex = currentIndex*2 + 1
		rightChildIndex = currentIndex*2 + 2
	}
}

// 只有当节点的距离变小时，才会调用此方法，更新整个小根堆(节点上移)
func (nh *NodeHeap) rewriteAndRecovery(index int) {
	if index < 0 || index >= len(nh.nodes) { // 不允许越界修改
		return
	}
	currentIndex := index // 当前节点在底层数组中的位置
	parentIndex := (currentIndex - 1) / 2

	for {
		if currentIndex == 0 { // 已到边界
			break
		}
		currentNodeDis := nh.distanceMap[nh.nodes[currentIndex]] // 当前节点距离目标节点的距离
		parentNodeDis := nh.distanceMap[nh.nodes[parentIndex]]   // 父节点距离目标节点的距离
		if currentNodeDis < parentNodeDis {
			nh.swap(currentIndex, parentIndex) // 节点上移
		} else {
			break
		}
		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2
	}

}

// addOrUpdateOrIgnore()方法是最核心的：
// 参数分别是节点node以及head到该node的距离
// 如果node从未加入过小根堆，则新增即可
// 如果node之前加入过,但是距离变小了，则更新整个小根堆
// 如果node之前加入过,但是距离没变小,则忽略
func (nh *NodeHeap) addOrUpdateOrIgnore(node *Node, distance int) {
	if nh.inHeap(node) { // 节点当前仍在小根堆上
		if distance >= nh.distanceMap[node] { // 距离没变小，则退出
			return
		} else { // 距离变小了，则需要更新(实际就是将距离更近的节点在底层数组中往前移动)
			nh.distanceMap[node] = distance              // 更新距离
			nh.rewriteAndRecovery(nh.heapIndexMap[node]) //更新该节点的值并重新堆化
		}

	}
	if !nh.isEntered(node) { // 节点从未出现过
		nh.nodes = append(nh.nodes, node)
		nh.heapIndexMap[node] = nh.size
		nh.distanceMap[node] = distance
		nh.heapify(nh.size)
		nh.size++
	}
}

type NodeRecord struct {
	Node     *Node
	Distance int
}

func (nh *NodeHeap) pop() NodeRecord {
	nodeRecord := NodeRecord{Node: nh.nodes[0], Distance: nh.distanceMap[nh.nodes[0]]}
	//待返回的根节点
	nh.swap(0, nh.size-1) // 交换根节点和最后一个节点在堆上的位置
	nh.heapIndexMap[nh.nodes[nh.size-1]] = -1
	delete(nh.distanceMap, nh.nodes[nh.size-1])
	nh.nodes[nh.size-1] = nil
	nh.size--
	nh.heapify(nh.size)
	return nodeRecord
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
