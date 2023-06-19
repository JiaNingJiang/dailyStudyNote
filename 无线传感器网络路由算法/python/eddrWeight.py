import math

E0 = 0.5  # 初始化时每个传感器节点的能量(单位是焦耳J)
packetBit = 4000  # 假设节点间通信的数据包的平均大小为4000bit
packetLen = 5  # 每次发送的数据包个数

epsilon_fs = 10.0 * math.pow(10, -12)
epsilon_mp = 0.0013 * math.pow(10, -12)

e_elec = 50.0 * math.pow(10, -9)
e_da = 5.0 * math.pow(10, -9)

d0 = math.sqrt(epsilon_fs / epsilon_mp)

# 发送 Lbit 数据到相距为 dis 的接收装置时，消耗的能量(单位是焦耳J)
def EDDR_Etx(lbit, dis):
    if float(dis) < d0:
        return float(lbit) * e_elec + float(lbit) * epsilon_fs * math.pow(dis, 2)
    else:
        return float(lbit) * e_elec + float(lbit) * epsilon_mp * math.pow(dis, 4)

# 接收 Lbit 数据消耗的能量
def EDDR_Erx(lbit):
    return float(lbit) * e_elec

# 融合 Lbit 数据消耗的能量
def EDDR_Eda(lbit):
    return float(lbit) * e_da

# 计算簇头到基站的权值（根据簇头到基站的距离dis和簇头当前的能量curE）
def ClusterToBS(dis, curE):
    K = E0 / curE
    return K * dis

# 接收方的Erx和Eda，发送方发送的数据包个数len，接收方的那个剩余能量RcurE, 双方的距离dis ，簇间的最短距离minDis
# 计算两个簇节点之间边的权值
def ClusterToCluster(RErx, REda, length, RcurE, dis, minDis):
    Era = (RErx + REda) * length
    K = RcurE / Era
    return K * math.pow(dis, 2) / minDis

if __name__ == "__main__":
    # 1.测试能量函数是否能够正常运行
    lbit = 4000
    dis = 99.0
    print("d0 = ", d0)
    print("发送消耗的能量为: ", packetLen*EDDR_Etx(lbit, dis))

    print("接收消耗的能量为: ", packetLen*EDDR_Erx(lbit))

    print("融合消耗的能量为: ", packetLen*EDDR_Eda(lbit))

    print()

    # 2.测试权重计算函数是否能够正常运行
    curE = 0.44
    print("簇点到基站边的权重为: ", ClusterToBS(dis, curE))

    # 2.1 权重与距离成正比
    minDis = 7.7
    dis = 7.7
    print("簇点到簇点的权重为: ", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis))
    dis = 8.9
    print("簇点到簇点的权重为: ", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis))

    print()
    # 2.2 权重与接收方能量成正比
    minDis = 7.7
    dis = 7.7
    print("簇点到簇点的权重为: ", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis))
    print("簇点到簇点的权重为: ", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0/2, dis, minDis))

