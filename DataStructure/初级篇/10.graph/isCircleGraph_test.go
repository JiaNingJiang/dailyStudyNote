package graph

import (
	"fmt"
	"testing"
)

func TestIsCirCleGraph(t *testing.T) {
	matrix := [][]int{
		{0, 1, 5}, {1, 2, 3}, {2, 0, 8}}
	//matrix := [][]int{
	//	{0, 1, 5}, {0, 2, 5}, {0, 3, 4}, {0, 4, 5},
	//	{1, 5, 3}, {2, 6, 7}, {3, 7, 11}, {4, 8, 3},
	//	{5, 0, 6}}
	graph := CreateGraph(matrix)

	if flag := IsCirCleGraph(graph); flag {
		fmt.Println("当前图有环")
	} else {
		fmt.Println("当前图无环")
	}
}
