package lesson8

import (
	"DataStructure2/utils"
	"fmt"
	"testing"
)

func TestTreeTotalNode(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8}
	tree := utils.NewTreeWithArr(arr)

	fmt.Println("节点总数: ", TreeTotalNode(tree.Root))
}
