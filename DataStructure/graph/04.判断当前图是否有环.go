package graph

import (
	"DataStructure/binaryTree"
	"reflect"
)

type NodeSet struct {
	// key为当前Node节点，value为该Node节点所在集合下的全部节点
	//(注意：绝对不能是切片！需要用map或数组这些扩充时地址不会变化的数据结构)
	SetMap map[*Node]map[*Node]struct{}
}

func NewNodeSet(graph *Graph) *NodeSet {
	nodeList := graph.Nodes
	nodeSet := &NodeSet{
		SetMap: make(map[*Node]map[*Node]struct{}),
	}

	for _, node := range nodeList { // 初始化，每个节点所在的集合中只有该节点本身
		nodeSet.SetMap[node] = make(map[*Node]struct{})
		nodeSet.SetMap[node][node] = struct{}{}
		// nodeSet.SetMap[node] = append(nodeSet.SetMap[node], node)
	}
	return nodeSet
}

// 判断一条边的from点和to点是否位于同一个节点集合中
func (ns *NodeSet) IsSameSet(from, to *Node) bool {
	fromSet := ns.SetMap[from]
	toSet := ns.SetMap[to]

	return reflect.DeepEqual(fromSet, toSet)
}

// 将from点和to点归到一个节点集合中(重点，尤其要注意此函数)
func (ns *NodeSet) Union(from, to *Node) {
	//fromSet := ns.SetMap[from]   不能是map集合的复制，必须对两个节点所在的map直接进行操作
	//toSet := ns.SetMap[to]

	for toNode, _ := range ns.SetMap[to] {
		// 不能使用切片，因为一旦append，切片的地址就会变化，导致节点实际上不能合并到内存中同一区域的数组中
		//fmt.Printf("追加元素前,ns.SetMap[from]地址为:%p\n", ns.SetMap[from])
		//ns.SetMap[from] = append(ns.SetMap[from], toNode) // 把to所在集合中的所有点归入到from所在集合中
		//fmt.Printf("追加元素后,ns.SetMap[from]地址为:%p\n", ns.SetMap[from])

		ns.SetMap[from][toNode] = struct{}{} // 把to所在集合中的所有点归入到from所在集合中
		ns.SetMap[toNode] = ns.SetMap[from]  // to节点所在集合的所有节点，其所在集合更改为from所在集合
	}
}

func IsCirCleGraph(graph *Graph) bool {
	edgeSet := graph.Edges // 取出所有边
	nodeSet := NewNodeSet(graph)

	queue := binaryTree.NewQueue()
	for edge, _ := range edgeSet { // 将所有边加入到队列中
		queue.Push(edge)
	}

	for {
		if queue.Len == 0 {
			return false
		}
		data := queue.Pop()
		edge := data.(Edge)
		if nodeSet.IsSameSet(edge.from, edge.to) {
			return true // 一旦发现当前边的 from和to在同一个点集合中，则必然有环
		}
		nodeSet.Union(edge.from, edge.to) // 将遍历完的边的from和to归并到一个点集合中
	}
}
