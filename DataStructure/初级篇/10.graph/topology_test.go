package graph

import "testing"

func TestTopology(t *testing.T) {
	matrix := [][]int{
		{0, 1, 5}, {0, 2, 5}, {0, 3, 4}, {0, 4, 5},
		{1, 5, 3}, {2, 6, 7}, {3, 7, 11}, {4, 8, 3},
		{5, 0, 6}}
	graph := CreateGraph(matrix)

	Topology(graph)
}
