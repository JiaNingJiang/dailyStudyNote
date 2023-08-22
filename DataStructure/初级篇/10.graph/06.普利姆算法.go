package graph

import (
	"DataStructure/09.binaryTree"
	"sort"
)

func PrimMST(graph *Graph) []Edge {

	knownNode := make(map[*Node]struct{})       // 已知点集合
	result := make([]Edge, 0, len(graph.Edges)) // 结果边集合(MST)

	queue := binaryTree.NewQueue()

	for _, node := range graph.Nodes { // 保证所有点都能遍历一遍(防止非连通图)
		if _, ok := knownNode[node]; !ok { // 新节点
			knownNode[node] = struct{}{} // 要记得把节点变为已知节点
			edges := make([]*Edge, 0, len(node.edge))
			edges = append(edges, node.edge...)
			sort.Slice(edges, func(i, j int) bool {
				if edges[i].weight < edges[j].weight {
					return true
				} else {
					return false
				}
			})
			for _, edge := range edges { // 新节点的所有边(按照从小到大)都加入到队列中
				queue.Push(edge)
			}

			for {
				if queue.Len == 0 {
					break
				}
				data := queue.Pop()
				edge := data.(*Edge)
				toNode := edge.to
				if _, ok := knownNode[toNode]; !ok { // 边的终点是新节点,则将改边加入到result
					knownNode[toNode] = struct{}{}
					result = append(result, *edge) // 该边将作为MST中的一条边

					// 将该终点的所有边排序后添加到queue中
					newEdges := make([]*Edge, 0, len(toNode.edge))
					newEdges = append(newEdges, toNode.edge...)
					sort.Slice(newEdges, func(i, j int) bool {
						if newEdges[i].weight < newEdges[j].weight {
							return true
						} else {
							return false
						}
					})
					for _, newEdge := range newEdges {
						queue.Push(newEdge)
					}
				} else { // 已读的节点，直接跳过
					continue
				}

			}
		}
	}
	return result
}
