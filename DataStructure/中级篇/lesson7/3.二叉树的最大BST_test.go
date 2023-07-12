package lesson7

import (
	"DataStructure2/utils"
	"fmt"
	"testing"
)

func TestMaxBST(t *testing.T) {
	left := []int{5, 3, 7, 2, 4, 6, 8, 1}
	leftTree := utils.NewTreeWithArr(left)

	utils.InOrderTraversal(leftTree)
	fmt.Println()

	right := []int{14, 12, 16, 11, 13, 15, 17, 10}
	rightTree := utils.NewTreeWithArr(right)

	utils.InOrderTraversal(rightTree)
	fmt.Println()

	root := &utils.Node{Data: 9}
	root.Left = leftTree.Root
	root.Right = rightTree.Root

	newTree := new(utils.Tree)
	newTree.Root = root
	newTree.Size = leftTree.Size + rightTree.Size + 1
	//newTree.Size = leftTree.Size + 0 + 1
	//newTree.Size = 0 + rightTree.Size + 1

	utils.InOrderTraversal(newTree)
	fmt.Println()

	maxBSTTree := MaxBST(root)

	fmt.Println(maxBSTTree)

}
