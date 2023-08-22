package binaryTree

import (
	"fmt"
	"testing"
)

func TestFindSuccessor(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := NewTreeWithArr(arr)

	InOrderTraversal(tree)

	node := tree.Root.Right.Right
	if successor := FindSuccessor(tree, node); successor != nil {
		fmt.Printf("\n节点%d的后继节点为:%d", node.Data, successor.Data)
	} else {
		fmt.Printf("\n节点%d的后继节点不存在", node.Data)
	}

}
