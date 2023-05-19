package graph

import (
	"fmt"
	"testing"
)

func TestKruskalMST(t *testing.T) {
	matrix := [][]int{
		{0, 1, 5}, {1, 2, 3}, {2, 0, 8}}
	graph := CreateGraph(matrix)

	mst := KruskalMST(graph) // 获取最小生成树的边

	for _, edge := range mst {
		fmt.Println(edge)
	}

}
