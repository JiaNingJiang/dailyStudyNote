package binaryTree

import (
	"fmt"
	"testing"
)

func TestGetBinaryTreeMaxWidth(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	tree := NewTreeWithArr(arr)

	fmt.Println("(哈希表法)当前二叉树的最大宽度为：", GetBinaryTreeMaxWidth(tree))
	fmt.Println("(改进法)当前二叉树的最大宽度为：", GetBinaryTreeMaxWidthImproved(tree))
}
