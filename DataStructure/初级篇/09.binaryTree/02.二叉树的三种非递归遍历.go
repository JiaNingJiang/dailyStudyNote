package binaryTree

import (
	"DataStructure/08.linkList"
	"fmt"
)

func PreOrderNoRecursive(tree *Tree) {
	stack := linkList.NewStack() // 创建一个栈
	stack.Push(tree.Root)

	for {
		if stack.Len == 0 { // 栈中再无节点，这意味遍历结束
			return
		}
		data := stack.Pop() // 子树的根节点出栈，并打印内容
		node := data.(*Node)
		fmt.Printf("%d ", node.Data)
		if node.Right != nil { // 先让右孩子进栈
			stack.Push(node.Right)
		}
		if node.Left != nil { // 再让左孩子出现
			stack.Push(node.Left)
		}
	}
}

func InOrderNoRecursive(tree *Tree) {
	stack := linkList.NewStack() // 创建一个栈
	stack.Push(tree.Root)

	current := tree.Root
	iterationLeftPush(current, stack) // 第一个循环，不断将当前节点的左孩子节点入栈，直到没有左孩子
	for {
		if stack.Len == 0 {
			break
		}
		data := stack.Pop()
		node := data.(*Node)
		fmt.Printf("%d ", node.Data) // 打印这个出栈的节点
		if node.Right != nil {       // 如果当前被弹出节点有右孩子，则采取同样的措施(子树的左孩子节点不断入栈)
			stack.Push(node.Right) // 右孩子先进栈
			iterationLeftPush(node.Right, stack)
		}
	}
}

// 从当前子树根节点开始，不断让左孩子节点入栈
func iterationLeftPush(root *Node, stack *linkList.Stack) {
	current := root
	for {
		if current.Left == nil {
			return
		}
		current = current.Left
		stack.Push(current)
	}
}

func PostOrderNoRecursive(tree *Tree) {
	stack1 := linkList.NewStack()
	stack2 := linkList.NewStack()

	stack1.Push(tree.Root)
	for { // 第一个循环，让所有节点按照后序遍历顺序存入到栈2中(根节点再栈2的最下层)
		if stack1.Len == 0 {
			break
		}
		data := stack1.Pop() // 从栈1弹出的元素存到栈2中
		node := data.(*Node)
		stack2.Push(node)

		if node.Left != nil { // 先让左孩子进栈
			stack1.Push(node.Left)
		}
		if node.Right != nil { // 再让右孩子进栈
			stack1.Push(node.Right)
		}
	}

	for { // 第二个循环，打印后序遍历
		if stack2.Len == 0 {
			break
		}
		data := stack2.Pop()
		node := data.(*Node)
		fmt.Printf("%d ", node.Data) // 打印这个出栈的节点
	}

}
