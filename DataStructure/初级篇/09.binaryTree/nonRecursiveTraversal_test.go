package binaryTree

import (
	"fmt"
	"testing"
)

func TestNonRecursiveTraversal(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	tree := NewTreeWithArr(arr)

	PreOrderNoRecursive(tree)
	fmt.Println()

	InOrderNoRecursive(tree)
	fmt.Println()

	PostOrderNoRecursive(tree)
	fmt.Println()
}
