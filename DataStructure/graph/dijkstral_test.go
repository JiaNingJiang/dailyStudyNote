package graph

import (
	"fmt"
	"testing"
)

func TestDijkstral(t *testing.T) {
	matrix := [][]int{
		{0, 1, 5}, {1, 2, 3}, {0, 2, 6}}
	graph := CreateGraph(matrix)

	distanceMap := Dijkstral(graph.Source)

	for node, distance := range distanceMap {
		fmt.Printf("节点(%d)到节点(%d)的最短距离为:%d\n", graph.Source.value, node.value, distance)
	}
}
