package binaryTree

import (
	"fmt"
	"testing"
)

func TestCheckFBT(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8}
	tree := NewTreeWithArr(arr)

	//BinaryTreeBFS(tree)
	//fmt.Println()

	flag := CheckFBT(tree)
	if flag {
		fmt.Println("当前二叉树是满二叉树")
	} else {
		fmt.Println("当前二叉树不是满二叉树")
	}
}
