package main

import (
	"fmt"
	"math"
)

// 计算给定点到图（有向图）中各点的最小距离
func Dijkstra(node *Node) map[*Node]float64 {
	energyMap := make(map[*Node]float64, 0)
	weightMap := make(map[*Node]float64, 0) // (距离全部初始化为0)如果节点没有出现在distanceMap中，则说明其距离node为∞
	weightMap[node] = 0                     // 距离自己的距离就是0
	// 已经求过距离的节点，存在selectedNodes中，以后再也不碰
	selectedMap := make(map[*Node]struct{}, 0)

	minNode := getMinDistanceAndUnselectedNode(node, energyMap, weightMap, selectedMap) // 一开始，minNode就是节点自己
	if minNode == nil {                                                                 // 节点因为能量不够，走不下去了
		return energyMap
	}

	for {
		if len(selectedMap) == len(weightMap) { // 所有节点都已经走过一遍
			return energyMap
		}

		weight := weightMap[minNode] // 当前路径的积攒距离

		for _, edge := range minNode.edge {
			toNode := edge.to
			if _, ok := weightMap[toNode]; !ok { // 再走新一步时，遇到新的节点(从未在distanceMap中出现)，则记录(新建)到该新节点的距离
				weightMap[toNode] = weight + edge.weight
			} else { // 遇到已有节点，则更新距离
				weightMap[toNode] = math.Min(float64(weightMap[toNode]), float64(weight+edge.weight))
			}
		}
		selectedMap[minNode] = struct{}{} // 该节点已经被走过

		minNode = getMinDistanceAndUnselectedNode(minNode, energyMap, weightMap, selectedMap)
		if minNode == nil { // 节点因为能量不够，走不下去了
			return energyMap
		}

	}
}

// 找到最优的下一节点，如果找不到则返回nil
func getMinDistanceAndUnselectedNode(cur *Node, energyMap map[*Node]float64, weightMap map[*Node]float64, selectedMap map[*Node]struct{}) *Node {

	minWeight := math.MaxFloat64
	var minNode *Node

	var factDistance float64

	for toNode, weight := range weightMap {
		if _, ok := selectedMap[toNode]; ok {
			continue
		} else {
			if weight < minWeight {
				if weight > 0 { // distance == 0 时找到的下一个节点是自己
					for _, edge := range cur.edge {
						if edge.from == cur && edge.to == minNode {
							factDistance = edge.dis
						}
					}
					// 1.需要确保发送节点有足够的能量完成本次发送
					if cur.energy < packetLen*EDDR_Etx(packetBit, factDistance) {
						fmt.Printf("发送节点能量不够,distance:%f cur.Energy:%f \n", factDistance, cur.energy)
						panic("发送节点能量不够")
					}
					// 2.需要确保接收节点有足够的能量完成本次接收
					if toNode.energy < packetLen*(EDDR_Erx(packetBit)+EDDR_Eda(packetBit)) {
						fmt.Printf("发送节点能量不够,distance:%f cur.Energy:%f \n", factDistance, cur.energy)
						panic("接受节点能量不够")
					}
				}
				minWeight = weight
				minNode = toNode
			}
		}
	}
	if minNode != nil { // 如果挑选出了作为下一起点的新节点，那么新旧节点都需要消耗能量
		if minWeight > 0 {
			cur.energy -= packetLen * EDDR_Etx(packetBit, factDistance) // 发送节点发送需要消耗能量
			minNode.energy -= packetLen * EDDR_Erx(packetBit)           // 接收节点需要减去接收消耗的能量
			minNode.energy -= packetLen * EDDR_Eda(packetBit)           // 接收节点需要减去融合需要的能量

			// 更新边的权重
			var minDis float64 = math.MaxFloat64 // 记录当前节点所有临边中距离最短的长度
			for _, edge := range cur.edge {
				if edge.dis < minDis {
					minDis = edge.dis
				}
			}

			//var weight float64 // 两点间边的权重
			//if cur.role == BaseStation && minNode.role == Cluster {
			//	weight = math.MaxFloat64
			//} else if cur.role == Cluster && minNode.role == BaseStation {
			//	weight = ClusterToBS(factDistance, cur.energy)
			//} else if cur.role == Cluster && minNode.role == Cluster {
			//	weight = ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, minNode.energy, factDistance, minDis)
			//}
			//
			//for _, edge := range cur.edge { // 更新该边(cur --> minNode )的权重
			//	if edge.from == cur && edge.to == minNode {
			//		//fmt.Printf("\n 旧的权重为:%f  新的权重为:%f \n", edge.weight, weight)
			//		edge.weight = weight
			//	}
			//}

		}
		energyMap[cur] = cur.energy
		energyMap[minNode] = minNode.energy
		return minNode
	} else {
		return nil
	}

}
