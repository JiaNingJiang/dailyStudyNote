package binaryTree

import (
	"fmt"
	"testing"
)

func TestFindMinAncestor(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := NewTreeWithArr(arr)

	BinaryTreeBFS(tree)
	fmt.Println()

	ancestor1 := FindMinAncestor(tree, tree.Root.Left.Left, tree.Root.Left.Right)
	fmt.Println(ancestor1)

	ancestor2 := FindMinAncestorDP(tree, tree.Root.Left.Left, tree.Root.Left.Right)
	fmt.Println(ancestor2)
}
