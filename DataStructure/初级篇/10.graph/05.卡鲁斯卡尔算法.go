package graph

import (
	"DataStructure/09.binaryTree"
	"sort"
)

func KruskalMST(graph *Graph) []Edge {
	// 1.将图的边集合按照从大到小的顺序排序
	edgeSet := make([]Edge, 0, len(graph.Edges))
	for edge, _ := range graph.Edges {
		edgeSet = append(edgeSet, edge)
	}
	sort.Slice(edgeSet, func(i, j int) bool { //升序排序
		if edgeSet[i].weight < edgeSet[j].weight {
			return true
		} else {
			return false
		}
	})
	nodeSet := NewNodeSet(graph)

	queue := binaryTree.NewQueue()
	for _, edge := range edgeSet { // 将所有边加入到队列中
		queue.Push(edge)
	}
	result := make([]Edge, 0) // 最后要返回的边集合
	for {
		if queue.Len == 0 {
			break
		}
		data := queue.Pop()
		edge := data.(Edge)
		if nodeSet.IsSameSet(edge.from, edge.to) {
			continue
		} else {
			result = append(result, edge)
			nodeSet.Union(edge.from, edge.to)
		}
	}
	return result
}
