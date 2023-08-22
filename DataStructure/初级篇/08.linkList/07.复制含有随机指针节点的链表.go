package linkList

import "fmt"

type RNode struct {
	Data   interface{}
	Next   *RNode
	Random *RNode // 指向一个随机的节点(也可能是nil)
}

type SLinkList struct {
	Head *RNode
	Len  int
}

func NewSLinkList() *SLinkList {
	return &SLinkList{
		Head: nil,
		Len:  0,
	}
}

func (sl *SLinkList) TailAdd(data interface{}) *RNode {
	rnode := new(RNode)
	rnode.Data = data
	if sl.Head == nil {
		sl.Head = rnode
		sl.Len++
		return rnode
	}
	current := sl.Head
	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	current.Next = rnode
	sl.Len++
	return rnode
}

func (sl *SLinkList) RandomAdd(src, dest int) *RNode {
	if src > sl.Len || sl.Len == 0 { // 要添加Random指针的节点不存在
		return nil
	}
	srcNode := sl.Head
	for i := 1; i < src; i++ { // 第一个循环，找到要添加Random指针的节点
		srcNode = srcNode.Next
	}
	if dest > sl.Len { // dest节点不存在
		srcNode.Random = nil
		return srcNode
	}
	destNode := sl.Head
	for i := 1; i < dest; i++ { // 第二个循环，找到Random要指向的目标节点
		destNode = destNode.Next
	}
	srcNode.Random = destNode
	return srcNode
}

func (sl *SLinkList) Print() {
	if sl.Len == 0 {
		fmt.Println("链表为空")
	}
	current := sl.Head
	for {
		if current == nil {
			fmt.Println("\n链表长度为:", sl.Len)
			break
		}
		fmt.Printf("%v(rand->%v) ", current.Data, current.Random)
		current = current.Next
	}
}

func CopyLinkListWithRandomP(sl *SLinkList) *SLinkList {
	if sl.Len == 0 {
		return nil
	}

	newsl := NewSLinkList()
	hashMap := make(map[*RNode]*RNode, 0)
	current := sl.Head
	for { // 第一个循环，创建复制节点(只有Data字段被复制)，记录原始节点与复制节点的映射关系
		if current == nil {
			break
		}
		snode := new(RNode)       // 复制当前链表节点
		snode.Data = current.Data // 第一次循环仅复制Data字段
		hashMap[current] = snode  // 记录原始节点与复制节点的映射关系

		current = current.Next
	}
	// 第二个循环，让克隆节点完成next/random指针的复制
	for base, clone := range hashMap {
		clone.Next = hashMap[base.Next]
		clone.Random = hashMap[base.Random]
	}
	newsl.Len = sl.Len
	newsl.Head = hashMap[sl.Head]

	return newsl
}
