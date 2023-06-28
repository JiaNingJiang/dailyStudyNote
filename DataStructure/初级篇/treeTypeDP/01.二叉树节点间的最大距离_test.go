package treeTypeDP

import (
	"DataStructure/binaryTree"
	"fmt"
	"math"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{1, 2, math.MinInt, 4, 5, math.MinInt, math.MinInt, 8, math.MinInt, 9,
		math.MinInt, math.MinInt, math.MinInt, math.MinInt, math.MinInt, 15}
	tree := binaryTree.NewTreeWithArr(arr)

	fmt.Println("最大距离: ", MaxDistance(tree))
}
