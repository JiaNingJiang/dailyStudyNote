package graph

import (
	"DataStructure/09.binaryTree"
	"fmt"
)

// 只能用于有向无环图的遍历(如果有环，环上所有点是无法遍历到的)
func Topology(graph *Graph) {
	inMap := make(map[*Node]int) // 存储所有节点的入度

	zeroQueue := binaryTree.NewQueue() // 用于存零入度节点的队列
	for _, node := range graph.Nodes { // 遍历整个点集合，1.存储所有节点的入度  2.将一个零入度的节点加入到队列
		if node.in == 0 {
			zeroQueue.Push(node)
		}
		inMap[node] = node.in
	}

	result := make([]interface{}, 0) // 存储所有已遍历节点的value
	for {
		if zeroQueue.Len == 0 {
			break
		}
		data := zeroQueue.Pop()
		node := data.(*Node)
		result = append(result, node.value) // 遍历该节点
		//delete(inMap, node)                 // 将该遍历节点从inMap中的记录删除

		for _, next := range node.nexts { // 让该节点的所有邻居节点在inMap中的入度-1
			inMap[next]--
			if inMap[next] == 0 { // 如果出现新的入度为0的点，将其新加到零入度节点队列中
				zeroQueue.Push(next)
			}
		}
	}

	for _, node := range result {
		fmt.Printf("%d ", node)
	}

}
