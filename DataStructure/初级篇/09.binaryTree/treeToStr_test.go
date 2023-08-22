package binaryTree

import (
	"fmt"
	"testing"
)

func TestSerializeByPre(t *testing.T) {
	arr := []int{5, 3, 7, 2, 4, 6, 8, 1}
	tree := NewTreeWithArr(arr)

	PreOrderTraversal(tree)
	fmt.Println()

	str := SerializeByPre(tree.Root)
	fmt.Printf("将二叉树序列化(前序)后：%s\n", str)

	newTree := DeserializationByPre(str)

	PreOrderTraversal(newTree)
	fmt.Println()
}
