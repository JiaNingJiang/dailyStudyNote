package main

import (
	"fmt"
	"sync"
)

// TODO: 数组实现的栈
type arrayStack struct {
	stack []interface{} // 底层切片
	size  int           // 元素个数
	lock  sync.Mutex
}

func NewArrayStack() *arrayStack {
	return &arrayStack{
		stack: make([]interface{}, 0),
	}
}

func (as *arrayStack) Push(v interface{}) {
	as.lock.Lock()
	defer as.lock.Unlock()

	if as.stack == nil {
		as.stack = make([]interface{}, 0)
	}

	as.stack = append(as.stack, v)
	as.size += 1
}

func (as *arrayStack) Pop() interface{} {

	as.lock.Lock()
	defer as.lock.Unlock()

	if as.stack == nil {
		return nil
	}

	v := as.stack[as.size-1]
	as.size -= 1
	as.stack = as.stack[:as.size]

	return v
}

func (as *arrayStack) Len() int {
	return as.size
}

func (as *arrayStack) Top() interface{} {
	return as.stack[as.size-1]
}

// TODO:链表实现的栈
type LinkStack struct {
	root *Node
	size int
	mux  sync.Mutex
}

type Node struct {
	next  *Node
	value interface{}
}

func NewLinkStack() *LinkStack {
	return &LinkStack{
		root: new(Node),
	}
}

func (ls *LinkStack) Push(v interface{}) {
	ls.mux.Lock()
	defer ls.mux.Unlock()

	if ls.root == nil {
		newNode := new(Node)
		newNode.value = v
		ls.root = newNode
		ls.size += 1

		return
	}

	newNode := new(Node)
	newNode.value = v
	newNode.next = ls.root

	ls.root = newNode
	ls.size += 1
}

func (ls *LinkStack) Pop() interface{} {
	ls.mux.Lock()
	defer ls.mux.Unlock()

	if ls.root == nil {
		return nil
	}
	newRoot := ls.root.next

	v := ls.root.value

	ls.size -= 1

	ls.root = newRoot

	return v
}

func (ls *LinkStack) Len() int {
	return ls.size
}

func (ls *LinkStack) Top() interface{} {
	return ls.root.value
}

func main() {
	//arrayStack := NewArrayStack()
	//arrayStack.Push("cat")
	//arrayStack.Push("dog")
	//arrayStack.Push("hen")
	//fmt.Println("size:", arrayStack.Len())
	//fmt.Println("pop:", arrayStack.Pop())
	//fmt.Println("pop:", arrayStack.Pop())
	//fmt.Println("size:", arrayStack.Len())
	//arrayStack.Push("drag")
	//fmt.Println("pop:", arrayStack.Pop())

	linkStack := new(LinkStack)
	linkStack.Push("cat")
	linkStack.Push("dog")
	linkStack.Push("hen")
	fmt.Println("size:", linkStack.Len())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("size:", linkStack.Len())
	linkStack.Push("drag")
	fmt.Println("pop:", linkStack.Pop())
	fmt.Println("pop:", linkStack.Pop())
}
