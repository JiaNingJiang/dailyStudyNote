package graph

import "testing"

func TestCreateGraph(t *testing.T) {
	matrix := [][]int{{0, 1, 5},
		{1, 2, 3},
		{2, 3, 6},
		{3, 0, 1}}
	CreateGraph(matrix)
}
