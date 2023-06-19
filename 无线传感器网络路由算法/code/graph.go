package main

import "math"

type role uint8

const (
	BaseStation role = iota
	Cluster
	Common
)

type Node struct {
	role   role    // 节点的类型
	key    float64 // 区分节点的标识符
	energy float64 // 存储的能量 (能量可以在初始化阶段设置为一个统一的值)
	out    int     // 出度
	in     int     // 入度
	nexts  []*Node // 当前节点的所有邻居节点
	edge   []*Edge // 当前节点的所有边
}

type Edge struct {
	from   *Node   // 边的起点
	to     *Node   // 边的终点
	dis    float64 // 边表示的实际距离
	weight float64 // 边的权重
}

type Graph struct {
	Source *Node                 // 一个用于遍历的起点
	Nodes  map[interface{}]*Node // 点集合
	Edges  map[Edge]struct{}     // 边集合
}

// 用一个二维数组构造图
func CreateGraph(matrix [][]float64) *Graph {
	graph := new(Graph)
	graph.Nodes = make(map[interface{}]*Node)
	graph.Edges = make(map[Edge]struct{})

	var minDis float64 = math.MaxFloat64 // 先遍历一遍matrix，获取簇节点到簇节点的最小的物理距离
	for _, edge := range matrix {
		dis := edge[4] // 两点的物理距离
		if dis < minDis {
			minDis = dis
		}
	}

	for _, edge := range matrix { // 每一个元素都是一条边：起点、终点、权重
		from := edge[0]           // 起点的标识符
		to := edge[1]             // 终点的标识符
		fromRole := role(edge[2]) // 起点的身份
		toRole := role(edge[3])   // 终点的身份
		dis := edge[4]            // 两点的物理距离

		var weight float64 // 两点间边的权重
		weight = dis
		//if fromRole == BaseStation && toRole == Cluster {
		//	weight = math.MaxFloat64
		//} else if fromRole == Cluster && toRole == BaseStation {
		//	weight = ClusterToBS(dis, E0)
		//} else if fromRole == Cluster && toRole == Cluster {
		//	weight = ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis)
		//}

		// 1. 构造起点
		if _, ok := graph.Nodes[from]; !ok { // 当前起点不存在于点集合中，则新加
			graph.Nodes[from] = &Node{role: fromRole, key: from, energy: E0}
		}
		// 2.构造终点
		if _, ok := graph.Nodes[to]; !ok { // 当前终点不存在于点集合中，则新加
			graph.Nodes[to] = &Node{role: toRole, key: to, energy: E0}
		}
		// 3.构造边(默认matrix数组中不会有重复的边)
		fromNode := graph.Nodes[from] // 获取已有的起点(后面进行修改)
		toNode := graph.Nodes[to]     // 获取已有的终点(后面进行修改)

		edge := &Edge{from: fromNode, to: toNode, dis: dis, weight: weight}

		fromNode.out++
		fromNode.nexts = append(fromNode.nexts, toNode)
		fromNode.edge = append(fromNode.edge, edge)

		toNode.in++

		graph.Edges[*edge] = struct{}{}

		//// 4.遇到的第一个点作为source
		//if graph.Source == nil { //
		//	graph.Source = fromNode
		//}
		graph.Source = fromNode
	}

	return graph
}
