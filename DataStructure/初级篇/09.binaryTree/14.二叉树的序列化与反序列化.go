package binaryTree

import (
	"fmt"
	"strconv"
	"strings"
)

// 此问题的关键所在是: 如果二叉树的一条路径到底了,需要为其加上一个独特的"结束符"。
// 反序列化时就可以从左子树转移向右子树，从右子树转移向根节点

// 按照前序遍历将二叉树变为字符串
func SerializeByPre(head *Node) string {
	if head == nil { // 一条路径的结束符
		return "#_"
	}
	var str string
	str += fmt.Sprintf("%d_", head.Data) // 加上当前子树的根节点
	str += SerializeByPre(head.Left)     // 加上左子树
	str += SerializeByPre(head.Right)    // 加上右子树

	return str
}

func DeserializationByPre(str string) *Tree {
	nodeList := strings.Split(str, "_") // 得到由二叉树各节点数值组成的数组
	queue := NewQueue()
	treeLen := 0
	for _, node := range nodeList {
		queue.Push(node)
		if node != "#" {
			treeLen++
		}
	}
	root := reconPreOrder(queue)
	newTree := new(Tree)
	newTree.Root = root
	newTree.Size = treeLen
	return newTree
}

func reconPreOrder(queue *Queue) *Node {
	data := queue.Pop()
	if data == "#" { // 表示原始的二叉树并不存在该节点（当前递归分支结束，已到达最后的叶子结点）
		return nil
	}
	node := new(Node)
	node.Data, _ = strconv.Atoi(data.(string)) // 构建当前根节点
	node.Left = reconPreOrder(queue)           // 构建左子树
	node.Right = reconPreOrder(queue)          // 构建右子树

	return node
}
