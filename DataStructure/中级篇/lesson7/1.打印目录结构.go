package lesson7

import (
	"fmt"
	"strings"
)

type Node struct {
	pass int              // 经过此节点的路径数
	end  int              // 以此节点为终点的路径数
	data string           // 为了方便打印，让节点存储其头顶路径的内容
	next map[string]*Node // 当前节点的 下一 节点
}

// 返回一颗前缀树的根节点
func NewPrefixTree() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

// 向前缀树中插入一条文件路径
func (ptree *Node) Insert(file string) {
	filePath := strings.Split(file, "\\\\")
	if filePath[len(filePath)-1] == "" {
		filePath = filePath[:len(filePath)-1]
	}

	ptree.pass++     // 所有的路径都起源于前缀树根节点
	curNode := ptree // 从根节点出发，拓展前缀树

	for _, dir := range filePath {
		if _, ok := curNode.next[dir]; !ok { // 下一个节点不存在，则需要创建
			curNode.next[dir] = &Node{next: make(map[string]*Node)}
		}
		curNode = curNode.next[dir] // 跳转到下一个节点
		curNode.data = dir
		curNode.pass++
	}
	curNode.end++
}

// 采用深度优先遍历方式打印整颗前缀树
func (ptree *Node) DFSPrintFilePath() {
	ptree.dFSPrintFilePath(0)
}

func (ptree *Node) dFSPrintFilePath(layer int) {
	if ptree == nil {
		return
	}

	for _, child := range ptree.next {
		dir := child.data
		// 1.先打印空格
		for i := 0; i < 2*layer; i++ {
			fmt.Printf(" ")
		}
		// 2.再打印孩子节点的内容(也就是目录名)
		fmt.Println(dir)
		child.dFSPrintFilePath(layer + 1) // 深度优先遍历的方式继续往下打印
	}
}
