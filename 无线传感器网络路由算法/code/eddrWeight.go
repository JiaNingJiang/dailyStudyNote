package main

import (
	"math"
)

const (
	E0 = 0.5 // 初始化时每个传感器节点的能量(单位是焦耳J)

	packetBit = 4000 // 假设节点间通信的数据包的平均大小为4000bit

	packetLen = 5 // 每次发送的数据包个数
)

var (
	epsilon_fs = 10.0 * math.Pow(10, -12)
	epsilon_mp = 0.0013 * math.Pow(10, -12)

	e_elec = 50.0 * math.Pow(10, -9)
	e_da   = 5.0 * math.Pow(10, -9)

	d0 = math.Sqrt(epsilon_fs / epsilon_mp)
)

// 发送 Lbit 数据到相距为 dis 的接收装置时，消耗的能量(单位是焦耳J)
func EDDR_Etx(lbit int, dis float64) float64 {
	if dis < d0 {
		return (float64(lbit)*e_elec + float64(lbit)*epsilon_fs*math.Pow(dis, float64(2)))
	} else {
		return (float64(lbit)*e_elec + float64(lbit)*epsilon_mp*math.Pow(dis, float64(4)))
	}
}

// 接收 Lbit 数据消耗的能量
func EDDR_Erx(lbit int) float64 {
	return (float64(lbit) * e_elec)
}

// 融合 Lbit 数据消耗的能量
func EDDR_Eda(lbit int) float64 {
	return (float64(lbit) * e_da)
}

// 计算簇头到基站的权值（根据簇头到基站的距离dis和簇头当前的能量curE）
func ClusterToBS(dis float64, curE float64) float64 {
	K := E0 / curE
	return K * dis
}

// 接收方的Erx和Eda，发送方发送的数据包个数len，接收方的那个剩余能量RcurE, 双方的距离dis ，簇间的最短距离minDis
// 计算两个簇节点之间边的权值
func ClusterToCluster(RErx float64, REda float64, len int, RcurE float64, dis float64, minDis float64) float64 {
	Era := (RErx + REda) * float64(len)
	//fmt.Printf("Era = %f\n", Era)
	K := RcurE / Era
	return K * math.Pow(dis, float64(2)) / minDis
}
