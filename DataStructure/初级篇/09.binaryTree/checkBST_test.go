package binaryTree

import (
	"fmt"
	"testing"
)

func TestCheckBST(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := NewTreeWithArr(arr)

	InOrderTraversal(tree)
	fmt.Println()

	//PreValue = math.MinInt
	//flag := CheckBST(tree.Root)
	//if flag {
	//	fmt.Println("当前二叉树是搜索二叉树")
	//} else {
	//	fmt.Println("当前二叉树不是搜索二叉树")
	//}

	flag1 := CheckBSTByDP(tree)
	if flag1 {
		fmt.Println("当前二叉树是搜索二叉树")
	} else {
		fmt.Println("当前二叉树不是搜索二叉树")
	}

}
