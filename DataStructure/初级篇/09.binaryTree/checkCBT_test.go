package binaryTree

import (
	"fmt"
	"testing"
)

func TestCheckCBT(t *testing.T) {
	arr := []int{5, 3, 7, 1}
	tree := NewTreeWithArr(arr)

	BinaryTreeBFS(tree)
	fmt.Println()

	flag := CheckCBT(tree)
	if flag {
		fmt.Println("当前二叉树是完全二叉树")
	} else {
		fmt.Println("当前二叉树不是完全二叉树")
	}

}
