import random
import eddrWeight as ew
import graph
import dijkstra as dk

edgeMax = 1000  # 总共的上限边数
nodeMax = 300  # 总共的节点(簇点)上限数
headShare = 0.3  # 簇头节点在所有簇节点中的份额
headMax = int(headShare * float(nodeMax))  # 簇头节点的上限个数

clToBSMin = ew.d0  # 簇头到基站的最近距离
clToBSMax = 2 * ew.d0  # 簇头到基站的最远距离

clToClMin = 5.0  # 簇节点到簇节点的最近距离
clToClMax = ew.d0  # 簇节点到簇节点的最远距离

headCount = 0  # 循环变量，已有的簇头数

matrix = [[] for _ in range(edgeMax)]

# 根据指定的概率返回簇节点的身份(簇头还是普通节点)
def nodeRole(p):
    count = int(p * 10)
    random1 = 0
    for i in range(10):
        random1 += random.randint(0, 1)  # 要么是1，要么是0

    if random1 <= count:
        return graph.Role.Cluster
    else:
        return graph.Role.Common



# 1.先构建簇节点到基站的边
for i in range(edgeMax):
    matrix[i] = [0.0] * 5
    curRole = nodeRole(headShare)  # 当前源节点的身份(簇头还是普通簇节点)
    if curRole == graph.Role.Cluster:  # 成为簇头节点，需要与基站相连
        # 1.成为簇头节点，需要与基站相连
        if headCount < headMax:
            random.seed()
            matrix[i] = [float(headCount), 0.0, float(graph.Role.Cluster), float(graph.Role.BaseStation), clToBSMin + float(random.randint(0, int(clToBSMax) - int(clToBSMin)))]
            headCount += 1

# 2.构建簇节点到簇节点的边
for i in range(edgeMax):
    if matrix[i][0] > 0:  # 此边已经存在，是簇头到基站的边
        continue
    else:
        random.seed()
        sender = random.randint(1, headMax)
        receiver = random.randint(1, headMax)
        while sender == receiver:
            receiver = random.randint(1, headMax)
        matrix[i] = [float(sender), float(receiver), float(graph.Role.Cluster), float(graph.Role.Cluster), clToClMin + float(random.randint(0, int(clToClMax) - int(clToClMin)))]



if __name__ == "__main__":
    g = graph.create_graph(matrix)

    i = 0
    while i < 300:
        res = dk.Dijkstra(g.source)

        for node, energy in res.items():
            print("第{}轮 -- 节点{}剩余的能量为:{}".format(i, node.key, energy))
        i=i+1
        print()