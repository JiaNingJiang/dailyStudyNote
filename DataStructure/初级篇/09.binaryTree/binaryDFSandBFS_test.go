package binaryTree

import "testing"

func TestBinaryTreeBFS(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewTreeWithArr(arr)

	BinaryTreeBFS(tree)
}

func TestBinaryTreeDFS(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewTreeWithArr(arr)

	BinaryTreeDFS(tree)
}
