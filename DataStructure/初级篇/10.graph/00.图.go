package graph

type Node struct {
	value interface{} // 存储的数值
	out   int         // 出度
	in    int         // 入度
	nexts []*Node     // 当前节点的所有邻居节点
	edge  []*Edge     // 当前节点的所有边
}

type Edge struct {
	from   *Node // 边的起点
	to     *Node // 边的终点
	weight int   // 边的权重
}

type Graph struct {
	Source *Node                 // 一个用于遍历的起点
	Nodes  map[interface{}]*Node // 点集合
	Edges  map[Edge]struct{}     // 边集合
}

// 用一个二维数组构造图
func CreateGraph(matrix [][]int) *Graph {
	graph := new(Graph)
	graph.Nodes = make(map[interface{}]*Node)
	graph.Edges = make(map[Edge]struct{})
	for _, edge := range matrix { // 每一个元素都是一条边：起点、终点、权重
		from := edge[0]   // 起点
		to := edge[1]     // 终点
		weight := edge[2] // 权重

		// 1. 构造起点
		if _, ok := graph.Nodes[from]; !ok { // 当前起点不存在于点集合中，则新加

			graph.Nodes[from] = &Node{value: from}
		}
		// 2.构造终点
		if _, ok := graph.Nodes[to]; !ok { // 当前终点不存在于点集合中，则新加
			graph.Nodes[to] = &Node{value: to}
		}
		// 3.构造边(默认matrix数组中不会有重复的边)
		fromNode := graph.Nodes[from] // 获取已有的起点(后面进行修改)
		toNode := graph.Nodes[to]     // 获取已有的终点(后面进行修改)

		edge := &Edge{from: fromNode, to: toNode, weight: weight}

		fromNode.out++
		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.edge = append(fromNode.edge, edge)

		toNode.in++

		graph.Edges[*edge] = struct{}{}

		// 4.遇到的第一个点作为source
		if graph.Source == nil { //
			graph.Source = fromNode
		}
	}

	return graph
}
