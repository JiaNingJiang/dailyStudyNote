package binaryTree

import (
	"fmt"
	"testing"
)

func TestCheckAVL(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := NewTreeWithArr(arr)

	BinaryTreeBFS(tree)
	fmt.Println()

	flag := CheckAVL(tree)
	if flag {
		fmt.Println("当前二叉树是平衡二叉树")
	} else {
		fmt.Println("当前二叉树不是平衡二叉树")
	}

}
