package graph

import (
	"DataStructure/08.linkList"
	"fmt"
)

func DFS(graph *Graph) {
	if len(graph.Nodes) == 0 {
		return
	}
	stack := linkList.NewStack()
	seenList := make(map[*Node]struct{}) // 可以防止有环图的死循环

	stack.Push(graph.Source)
	seenList[graph.Source] = struct{}{}
	fmt.Printf("%d ", graph.Source.value) // 每个节点都是入栈后即可打印

	for {
		if stack.Len == 0 {
			break
		}
		data := stack.Pop() // 节点出栈
		node := data.(*Node)

		for _, next := range node.nexts { // 遍历这个出栈节点的邻居节点是否已经被遍历过
			if _, ok := seenList[next]; !ok { // 此邻居节点还没有被遍历过
				stack.Push(node) // 将此节点和邻居节点先后入栈
				stack.Push(next)

				seenList[next] = struct{}{} // 记录该邻居节点已经被遍历
				fmt.Printf("%d ", next.value)
				break // 直接break，不再继续处理当前节点的其他邻居节点。作用是保留当前深度搜索路径。
			}
		}
	}
}
