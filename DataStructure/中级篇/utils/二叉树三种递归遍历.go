package utils

import "fmt"

func PreOrderTraversal(tree *Tree) {
	root := tree.Root
	preOrderTraversal(root)
}

func preOrderTraversal(node *Node) {
	if node == nil { // 重要：这个是边界条件
		return
	} else {
		fmt.Printf("%d ", node.Data)  // 第一次访问当前节点时打印
		preOrderTraversal(node.Left)  // 左递归函数返回当前节点时就意味着第二次访问当前节点
		preOrderTraversal(node.Right) // 右递归函数返回当前节点时就意味着第三次访问当前节点
	}

}

func InOrderTraversal(tree *Tree) {
	root := tree.Root
	inOrderTraversal(root)
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	} else {
		inOrderTraversal(node.Left)
		fmt.Printf("%d ", node.Data) // 第二次访问当前节点时打印
		inOrderTraversal(node.Right)
	}

}

func PostOrderTraversal(tree *Tree) {
	root := tree.Root
	postOrderTraversal(root)
}

func postOrderTraversal(node *Node) {
	if node == nil {
		return
	} else {
		postOrderTraversal(node.Left)
		postOrderTraversal(node.Right)
		fmt.Printf("%d ", node.Data) // 第三次访问当前节点时打印
	}

}
