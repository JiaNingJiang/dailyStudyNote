package mergeTable

import (
	"fmt"
	"testing"
)

func TestCountIsland2(t *testing.T) {
	arr := [][]int{
		{0, 0, 1, 0, 1, 0},
		{1, 1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0, 1},
		{0, 0, 0, 0, 1, 0},
	}

	chart := NewChart(arr)

	chart.CountIsland()

	fmt.Println("岛屿数量: ", chart.island)
}
