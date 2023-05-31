package binaryTree

import "math"

type Node struct {
	Data  interface{}
	Left  *Node // 左孩子节点
	Right *Node // 右孩子节点
}

type Tree struct {
	Root *Node // 根节点
	Size int   // 大小
}

// 通过数组构建一棵完全二叉树
func NewTreeWithArr(arr []int) *Tree {
	tree := new(Tree)
	tree.Size = len(arr)

	index := 0
	root := new(Node)
	root.Data = arr[index]
	root.Left = newTreeWithArr(arr, index*2+1)  // 递归创建左子树
	root.Right = newTreeWithArr(arr, index*2+2) // 递归创建右子树

	tree.Root = root
	return tree
}

func newTreeWithArr(arr []int, index int) *Node {
	if index >= len(arr) { // 重要：防止越界访问数组
		return nil
	}

	if arr[index] == math.MinInt { // 值为math.MinInt意味当前节点不存在(按照完全二叉树的顺序)
		return nil
	}

	node := new(Node)
	node.Data = arr[index]
	node.Left = newTreeWithArr(arr, index*2+1)
	node.Right = newTreeWithArr(arr, index*2+2)
	return node
}
