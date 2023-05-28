package mergeTable

import (
	"fmt"
	"testing"
)

func TestCountIsland(t *testing.T) {
	chart := [][]int{
		{0, 0, 1, 0, 1, 0},
		{1, 1, 1, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	fmt.Println("岛屿数量: ", CountIsland(chart))
}
