package main

import (
	"fmt"
	"testing"
)

func TestEDDR_Energy(t *testing.T) {
	lbit := packetBit
	dis := 99.0

	fmt.Println("d0 = ", d0)
	fmt.Printf("发送消耗的能量为: %f\n", packetLen*EDDR_Etx(lbit, dis))

	fmt.Printf("接收消耗的能量为: %f\n", packetLen*EDDR_Erx(lbit))

	fmt.Printf("融合消耗的能量为: %f\n", packetLen*EDDR_Eda(lbit))
}

func TestEDDR_Weight(t *testing.T) {
	dis := 89.0
	curE := 0.44
	fmt.Printf("簇点到基站边的权重为: %f\n", ClusterToBS(dis, curE))

	minDis := 7.7
	dis = 7.7
	fmt.Printf("簇点到簇点的权重为: %f\n", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis))
	dis = 8.9
	fmt.Printf("簇点到簇点的权重为: %f\n", ClusterToCluster(EDDR_Erx(packetBit), EDDR_Eda(packetBit), packetLen, E0, dis, minDis))
}
