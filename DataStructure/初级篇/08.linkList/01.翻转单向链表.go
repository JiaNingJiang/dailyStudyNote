package linkList

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkList struct {
	Head *Node // 头指针
	Len  int   // 链表长度
}

func NewLinkList() *LinkList {
	return &LinkList{
		Head: nil,
		Len:  0,
	}
}

// 尾插法
func (ll *LinkList) TailAdd(data interface{}) *Node {
	node := new(Node)
	node.Data = data
	node.Next = nil

	if ll.Head == nil {
		ll.Head = node
		ll.Len++
		return node
	}
	var current *Node = ll.Head
	for {
		if current.Next == nil {
			current.Next = node
			ll.Len++
			break
		}
		current = current.Next
	}
	return node
}

// 删除链表尾部元素,并且返回该元素值
func (ll *LinkList) TailDel() interface{} {
	if ll.Len == 1 {
		data := ll.Head.Data
		ll.Head = nil
		ll.Len--
		return data
	}
	var current *Node = ll.Head
	for {
		if current.Next.Next == nil {
			data := current.Next.Data
			current.Next = nil
			ll.Len--
			return data
		}
		current = current.Next
	}
}

func (ll *LinkList) HeadAdd(data interface{}) {
	node := new(Node)
	node.Data = data
	node.Next = nil

	if ll.Head == nil {
		ll.Head = node
		ll.Len++
		return
	}
	node.Next = ll.Head
	ll.Head = node
	ll.Len++
}

func (ll *LinkList) Print() {
	if ll.Head == nil {
		fmt.Println("当前链表为空....")
		return
	}
	var current *Node = ll.Head
	for {
		if current == nil {
			fmt.Println("\n链表长度为:", ll.Len)
			break
		} else {
			fmt.Printf("%v ", current.Data)
			current = current.Next
		}

	}
}

func CloneLinkList(ll *LinkList) *LinkList {
	newll := NewLinkList()
	current := ll.Head
	if current == nil {
		return newll
	}
	for {
		if current == nil {
			break
		}
		newll.TailAdd(current.Data)
		current = current.Next
	}
	return newll
}

// 1.迭代翻转链表
func (ll *LinkList) Reverse1() {
	if ll.Len <= 1 { // 没有翻转的必要
		return
	}
	var beg *Node = nil
	mid := ll.Head
	end := ll.Head.Next

	for {
		mid.Next = beg // 链表指向反转

		if end == nil {
			break
		}

		beg = mid // 三个指针依次往后移动一位
		mid = end
		end = end.Next
	}
	ll.Head = mid // 重新修正头指针指向
}

// 递归翻转链表
func (ll *LinkList) Reverse2() {
	if ll.Len <= 1 { // 没有翻转的必要
		return
	}
	newHead := Recursive_reverse(ll.Head)
	ll.Head = newHead
}

func Recursive_reverse(head *Node) *Node {
	if head.Next == nil { // 边界条件,找到尾节点
		return head
	}
	newHead := Recursive_reverse(head.Next) // 不断递归，直到当前节点的后继结点是尾节点,停止递归

	head.Next.Next = head // 让当前节点的后继节点指向自己
	head.Next = nil       // 因为后继节点已经指向自己，所以自己不必再指向后继节点

	return newHead // newHead在递归过程中不会发生变化，永远指向尾节点(反转后的新头结点)
}

// 头插法反转链表
func (ll *LinkList) Reverse3() {
	if ll.Len <= 1 { // 没有翻转的必要
		return
	}
	newLinkList := NewLinkList()
	current := ll.Head
	for { // 从头到尾遍历原始链表，读取到的节点以头插法插入到新的反向链表中
		if current == nil {
			break
		}
		newLinkList.HeadAdd(current.Data)
		current = current.Next
	}
	ll.Head = newLinkList.Head
}

// 就地逆置法反转链表
func (ll *LinkList) Reverse4() {
	if ll.Len <= 1 { // 没有翻转的必要
		return
	}

	beg := ll.Head
	end := beg.Next

	for {
		if beg.Next == nil {
			break
		}
		beg.Next = end.Next // 将end剥离，移动到首部
		end.Next = ll.Head  // end的next更新为头结点
		ll.Head = end       // 头指针指向end

		// beg = beg.Next // 重点：beg总是指向一个节点，那就是最初的头结点
		end = beg.Next // 更新循环变量end
	}
}
