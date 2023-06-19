package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDijkstra(t *testing.T) {
	matrix := generateMatrix()

	graph := CreateGraph(matrix)

	for round := 0; round < 300; round++ {
		res := Dijkstra(graph.Source)

		for node, energy := range res {
			fmt.Printf("第(%d)轮 -- 节点(%f)剩余的能量为:%f\n", round, node.key, energy)
		}
		fmt.Println()
	}
}

func generateMatrix() [][]float64 {
	edgeMax := 1000                              // 总共的上限边数
	nodeMax := 300                               // 总共的节点(簇点)上限数
	headShare := 0.3                             // 簇头节点在所有簇节点中的份额
	headMax := int(headShare * float64(nodeMax)) // 簇头节点的上限个数

	clToBSMin := d0     // 簇头到基站的最近距离
	clToBSMax := 2 * d0 // 簇头到基站的最远距离

	clToClMin := 5.0 // 簇节点到簇节点的最近距离
	clToClMax := d0  // 簇节点到簇节点的最远距离

	headCount := 0 // 循环变量，已有的簇头数

	matrix := make([][]float64, edgeMax)
	// 1.先构建簇节点到基站的边
	for i := 0; i < edgeMax; i++ {
		matrix[i] = make([]float64, 5)
		curRole := nodeRole(headShare) // 当前源节点的身份(簇头还是普通簇节点)
		if curRole == Cluster {        // 成为簇头节点，需要与基站相连
			// 1.成为簇头节点，需要与基站相连
			if headCount < headMax {
				rand.Seed(time.Now().UnixNano())
				matrix[i] = []float64{float64(headCount), 0, float64(Cluster), float64(BaseStation), clToBSMin + float64(rand.Intn(int(clToBSMax)-int(clToBSMin)))}
				headCount++
			}

		}
	}
	// 2.构建簇节点到簇节点的边
	for i := 0; i < edgeMax; i++ {
		if matrix[i][0] > 0 { // 此边已经存在，是簇头到基站的边
			continue
		} else {
			rand.Seed(time.Now().UnixNano())
			sender := rand.Intn(headMax-1) + 1
			receiver := rand.Intn(headMax-1) + 1
			for {
				if sender != receiver {
					break
				} else {
					receiver = rand.Intn(headMax-1) + 1
				}
			}
			matrix[i] = []float64{float64(sender), float64(receiver), float64(Cluster), float64(Cluster), clToClMin + float64(rand.Intn(int(clToClMax)-int(clToClMin)))}
			time.Sleep(2 * time.Nanosecond) // 防止相邻两项sender和receiver相同
		}
	}
	return matrix
}

func nodeRole(p float64) role {
	count := int(p * 10)
	random := 0
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		random += rand.Intn(2) // 要么是1，要么是0
	}

	if random <= count {
		return Cluster
	} else {
		return Common
	}

}
