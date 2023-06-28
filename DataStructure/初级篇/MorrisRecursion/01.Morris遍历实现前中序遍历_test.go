package MorrisRecursion

import (
	"DataStructure/binaryTree"
	"fmt"
	"testing"
)

func TestMorrisPreOrder(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := binaryTree.NewTreeWithArr(arr)

	Morris(tree, PreOrder)
}

func TestMorrisInOrder(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := binaryTree.NewTreeWithArr(arr)

	Morris(tree, InOrder)
}

func TestMorrisPostOrder(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := binaryTree.NewTreeWithArr(arr)

	Morris(tree, PostOrder)
	fmt.Println()
	Morris(tree, PostOrder)
}
