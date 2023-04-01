package main

import (
	"fmt"
	"sync"
)

type DoubleList struct {
	head *ListNode  // 指向链表头部
	tail *ListNode  // 指向链表尾部
	len  int        // 列表长度
	lock sync.Mutex // 为了进行并发安全pop操作
}

// 列表节点
type ListNode struct {
	pre   *ListNode // 前驱节点
	next  *ListNode // 后驱节点
	value string    // 值
}

// 获取节点值
func (node *ListNode) GetValue() string {
	return node.value
}

// 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在后驱节点
func (node *ListNode) HashNext() bool {
	return node.pre != nil
}

// 是否存在前驱节点
func (node *ListNode) HashPre() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

// 添加节点到链表头部的第N个元素之前，N=0表示新节点成为新的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {

}

func main() {
	list := new(DoubleList)
	list.AddNodeFromHead(0, "hello")
	list.AddNodeFromHead(0, "world")
	list.AddNodeFromHead(0, "ccc")

	for {
		node := list.head
		fmt.Printf("%v\n", node.value)
		node = node.next
		if node.next == nil {
			break
		}
	}

}
