package binaryTree

import (
	"fmt"
	"testing"
)

func TestRecursiveTraversal(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	tree := NewTreeWithArr(arr)

	PreOrderTraversal(tree)
	fmt.Println()

	InOrderTraversal(tree)
	fmt.Println()
	
	PostOrderTraversal(tree)
	fmt.Println()
}
