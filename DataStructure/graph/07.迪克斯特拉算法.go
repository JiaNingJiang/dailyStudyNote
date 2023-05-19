package graph

import "math"

// 计算给定点到图（有向图）中各点的最小距离
func Dijkstral(node *Node) map[*Node]int {
	distanceMap := make(map[*Node]int, 0) // (距离全部初始化为0)如果节点没有出现在distanceMap中，则说明其距离node为∞
	distanceMap[node] = 0                 // 距离自己的距离就是0
	// 已经求过距离的节点，存在selectedNodes中，以后再也不碰
	selectedMap := make(map[*Node]struct{}, 0)

	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedMap) // 一开始，minNode就是节点自己

	for {
		if len(selectedMap) == len(distanceMap) { // 所有节点都已经走过一遍
			return distanceMap
		}
		//if minNode == nil {
		//	return distanceMap
		//}
		distance := distanceMap[minNode] // 当前路径的积攒距离

		for _, edge := range minNode.edge {
			toNode := edge.to
			if _, ok := distanceMap[toNode]; !ok { // 再走新一步时，遇到新的节点(从未在distanceMap中出现)，则记录(新建)到该新节点的距离
				distanceMap[toNode] = distance + edge.weight
			} else { // 遇到已有节点，则更新距离
				distanceMap[toNode] = int(math.Min(float64(distanceMap[toNode]), float64(distance+edge.weight)))
			}
		}
		selectedMap[minNode] = struct{}{} // 该节点已经被走过
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedMap)
	}
}

func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, selectedMap map[*Node]struct{}) *Node {

	minDistance := math.MaxInt
	var minNode *Node

	for node, distance := range distanceMap {
		if _, ok := selectedMap[node]; ok {
			continue
		} else {
			if distance < minDistance {
				minDistance = distance
				minNode = node
			}
		}
	}

	return minNode
}
