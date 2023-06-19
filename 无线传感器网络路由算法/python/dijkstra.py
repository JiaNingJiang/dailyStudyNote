import math
import eddrWeight as ew
import graph
import sys

def Dijkstra(node):
    energyMap = {}
    weightMap = {node:0}  # (距离全部初始化为0)如果节点没有出现在distanceMap中，则说明其距离node为∞ ; 距离自己的距离就是0
    selectedMap = {} # 已经求过距离的节点，存在selectedNodes中，以后再也不碰

    minNode = getMinDistanceAndUnselectedNode(node, energyMap, weightMap, selectedMap)
    if minNode is None:
        return energyMap

    while True:
        if len(selectedMap) == len(weightMap): # 所有节点都已经走过一遍
            return energyMap

        weight = weightMap[minNode]  # 当前路径的积攒权重

        for edge in minNode.edges:
            toNode = edge.to
            if toNode not in weightMap: # 再走新一步时，遇到新的节点(从未在distanceMap中出现)，则记录(新建)到该新节点的距离
                weightMap[toNode] = weight + edge.weight
            else:
                weightMap[toNode] = min(weightMap[toNode], weight + edge.weight)

        selectedMap[minNode] = None

        minNode = getMinDistanceAndUnselectedNode(minNode, energyMap, weightMap, selectedMap)
        if minNode is None:  # 节点因为能量不够，走不下去了
            return energyMap

# 找到最优的下一节点，如果找不到则返回nil(并结束程序)
def getMinDistanceAndUnselectedNode(cur, energyMap, weightMap, selectedMap):
    minWeight = math.inf
    minNode = None

    for toNode, weight in weightMap.items():
        if toNode in selectedMap:
            continue
        else:
            if weight < minWeight:
                if weight > 0: # weight == 0 时找到的下一个节点是自己
                    factDistance = 0 
                    for edge in cur.edges:
                        if edge.from_ == cur and edge.to == minNode:
                            factDistance = edge.dis
                            break
                    sendNeed = ew.packetLen * ew.EDDR_Etx(ew.packetBit, factDistance)
                    # 1.需要确保发送节点有足够的能量完成本次发送
                    if cur.energy < sendNeed:
                        print("发送节点能量不够, 需要的能量: {} sender.Energy: {}" .format(sendNeed, cur.energy))
                        print()
                        sys.exit()
                        # continue
                    receiveNeed = ew.packetLen * (ew.EDDR_Erx(ew.packetBit) + ew.EDDR_Eda(ew.packetBit))
                    # 2.需要确保接收节点有足够的能量完成本次接收
                    if toNode.energy < receiveNeed:
                        print("接收节点能量不够, 需要的能量: {} receiver.Energy: {}" .format(receiveNeed, toNode.energy))
                        print()
                        sys.exit()
                        # continue

                minWeight = weight
                minNode = toNode

    if minNode is not None:
        if minWeight > 0: # 如果挑选出了作为下一起点的新节点，那么新旧节点都需要消耗能量
            cur.energy -= ew.packetLen * ew.EDDR_Etx(ew.packetBit, factDistance) # 发送节点发送需要消耗能量
            minNode.energy -= ew.packetLen * ew.EDDR_Erx(ew.packetBit) # 接收节点需要减去接收消耗的能量
            minNode.energy -= ew.packetLen * ew.EDDR_Eda(ew.packetBit) # 接收节点需要减去融合需要的能量

            if graph.Way == 1:
                # 更新边的权重
                minDis = math.inf  # 记录当前节点所有临边中距离最短的长度
                for edge in cur.edges:
                    if edge.dis < minDis:
                        minDis = edge.dis

                if cur.role == "BaseStation" and minNode.role == "Cluster":
                    weight = math.inf
                elif cur.role == "Cluster" and minNode.role == "BaseStation":
                    weight = ew.ClusterToBS(factDistance, cur.energy)
                elif cur.role == "Cluster" and minNode.role == "Cluster":
                    weight = ew.ClusterToCluster(
                        ew.EDDR_Erx(ew.packetBit),
                        ew.EDDR_Eda(ew.packetBit),
                        ew.packetLen,
                        minNode.energy,
                        factDistance,
                        minDis,
                    )

                for edge in cur.edges:  # 更新该边(cur --> minNode )的权重
                    if edge.from_ == cur and edge.to == minNode:
                        edge.weight = weight

        energyMap[cur] = cur.energy
        energyMap[minNode] = minNode.energy
        return minNode
    else:
        return None



