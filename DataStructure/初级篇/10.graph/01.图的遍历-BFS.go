package graph

import (
	"DataStructure/09.binaryTree"
	"fmt"
)

func BFS(graph *Graph) {
	if len(graph.Nodes) == 0 {
		return
	}
	queue := binaryTree.NewQueue()
	queue.Push(graph.Source)

	seenList := make(map[*Node]struct{}) // 防止有环图的死循环
	seenList[graph.Source] = struct{}{}

	for {
		if queue.Len == 0 {
			break
		}
		data := queue.Pop()
		node := data.(*Node)
		fmt.Printf("%d ", node.value)
		for _, next := range node.nexts {
			if _, ok := seenList[next]; !ok { // 防止死循环
				queue.Push(next)
				seenList[next] = struct{}{}
			}

		}
	}
}
