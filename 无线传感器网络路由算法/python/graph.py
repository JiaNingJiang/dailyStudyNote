import math
import eddrWeight as ew
import dijkstra as dk

Way = 1  ## way == 0 表示采用传统的迪杰斯特拉算法(节点距离作为权重)  way == 1 表示采用EDDR中距离和能量的综合作为节点权重


class Role:
    BaseStation = 0
    Cluster = 1
    Common = 2

    def from_float(value):
        if int(value) == 0 :
            return Role.BaseStation
        elif  int(value) == 1 :
            return Role.Cluster

class Node:
    def __init__(self, role, key, energy):
        self.role = role  # 节点的类型
        self.key = key  # 区分节点的标识符
        self.energy = energy  # 存储的能量 (能量可以在初始化阶段设置为一个统一的值)
        self.out = 0  # 出度
        self.in_ = 0  # 入度
        self.nexts = []  # 当前节点的所有邻居节点
        self.edges = []  # 当前节点的所有边

class Edge:
    def __init__(self, frm, to, dis, weight):
        self.from_ = frm # 边的起点
        self.to = to # 边的终点
        self.dis = dis # 边表示的实际距离
        self.weight = weight # 边的权重

class Graph:
    def __init__(self):
        self.source = None # 一个用于遍历的起点
        self.nodes = {} # 点集合
        self.edges = {} # 边集合

# 用一个二维数组(边集合)构造图
def create_graph(matrix):
    graph = Graph()
    graph.nodes = {}
    graph.edges = {}

    min_dis = math.inf
    for edge in matrix:
        dis = edge[4]
        if dis < min_dis:
            min_dis = dis

    for edge in matrix:
        frm = edge[0] # 起点的标识符
        to = edge[1] # 终点的标识符
        from_role = Role.from_float(edge[2]) # 起点的身份
        to_role = Role.from_float(edge[3]) # 终点的身份
        dis = edge[4] # 两点的物理距离

        weight = dis  ## 只使用距离作为权重

        if Way == 1:   ## 不只使用距离作为权重
            if from_role == Role.BaseStation and to_role == Role.Cluster:
                weight = math.inf
            elif from_role == Role.Cluster and to_role == Role.BaseStation:
                weight = ew.ClusterToBS(dis, ew.E0)
            elif from_role == Role.Cluster and to_role == Role.Cluster:
                weight = ew.ClusterToCluster(ew.EDDR_Erx(ew.packetBit), ew.EDDR_Eda(ew.packetBit), ew.packetLen, ew.E0, dis, min_dis)

        # 1. 构造起点
        if frm not in graph.nodes:
            graph.nodes[frm] = Node(from_role, frm, ew.E0)
        # 2.构造终点
        if to not in graph.nodes:
            graph.nodes[to] = Node(to_role, to, ew.E0)
        # 3.构造边(默认matrix数组中不会有重复的边)
        from_node = graph.nodes[frm]
        to_node = graph.nodes[to]
        edge = Edge(from_node, to_node, dis, weight)

        from_node.out += 1
        from_node.nexts.append(to_node)
        from_node.edges.append(edge)

        to_node.in_ += 1

        graph.edges[edge] = None

        # 4.让最后一个点成为source
        graph.source = from_node

    return graph





# if __name__ == "__main__":
#     matrix = [
# 		[1, 0, float(Role.Cluster), float(Role.BaseStation), 99.3],
# 		[2, 0, float(Role.Cluster), float(Role.BaseStation), 54.5],
#         [3, 0, float(Role.Cluster), float(Role.BaseStation), 67.2],

# 		[1, 2, float(Role.Cluster), float(Role.Cluster), 22.4],
# 		[1, 3, float(Role.Cluster), float(Role.Cluster), 7.7],

# 		[2, 1, float(Role.Cluster), float(Role.Cluster), 22.4],
# 		[2, 3, float(Role.Cluster), float(Role.Cluster), 8.9],

# 		[3, 1, float(Role.Cluster), float(Role.Cluster), 7.7],
# 		[3, 2, float(Role.Cluster), float(Role.Cluster), 8.9],
#         ]
#     graph = create_graph(matrix)

#     i = 0
#     while i < 300:
#         res = dk.Dijkstra(graph.source)

#         for node, energy in res.items():
#             print("第{}轮 -- 节点{}剩余的能量为:{}".format(i, node.key, energy))
#         i=i+1
#         print()