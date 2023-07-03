package lesson2

import (
	"fmt"
	"testing"
)

func TestBinaryTreeCount(t *testing.T) {
	n := 4

	fmt.Println("可以存在的二叉树个数: ", BinaryTreeCount(n))

	fmt.Println("可以存在的二叉树个数: ", BinaryTreeCount2(n))
}
