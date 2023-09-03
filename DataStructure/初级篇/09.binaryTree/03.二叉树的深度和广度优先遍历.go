package binaryTree

import (
	"fmt"
)

func BinaryTreeDFS(tree *Tree) {
	PreOrderTraversal(tree) // 二叉树的深度优先遍历就是先序遍历
}

type Queue struct {
	Items []interface{}
	Len   int
}

func NewQueue() *Queue {

	return &Queue{
		Items: make([]interface{}, 0),
		Len:   0,
	}
}

func (q *Queue) Push(data interface{}) {
	q.Items = append(q.Items, data)
	q.Len++
}

func (q *Queue) Pop() interface{} {
	data := q.Items[0]
	q.Items = q.Items[1:] // 删除出队列的元素
	q.Len--
	return data
}

func BinaryTreeBFS(tree *Tree) {
	queue := NewQueue()

	queue.Push(tree.Root)
	for {
		if queue.Len == 0 {
			return
		}
		data := queue.Pop()
		node := data.(*Node)

		fmt.Printf("%d ", node.Data)
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		//用来生成一颗非完全二叉树(非平衡二叉树)
		//if queue.Len == 0 {
		//	temp := new(Node)
		//	temp.Data = 11
		//	node.Left = temp
		//}
	}
}
