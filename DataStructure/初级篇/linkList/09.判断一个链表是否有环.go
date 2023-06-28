package linkList

// 为链表添加环（将尾节点的Next指针指向entryPoint个节点上）
func AddCircleForLinkList(ll *LinkList, entryPoint int) {
	if ll.Len <= 1 {
		return
	}
	currentNode := ll.Head
	currentIndex := 1
	var entryNode *Node // 记录entryPoint个节点的地址
	for {
		if currentNode.Next == nil {
			break
		}
		if currentIndex == entryPoint {
			entryNode = currentNode
		}
		currentNode = currentNode.Next
		currentIndex++
	}
	currentNode.Next = entryNode
}

func IsHaveCircle(ll *LinkList) (bool, *Node) {
	if ll.Len <= 1 {
		return false, nil
	}
	if ll.Len == 2 { // 有两个节点的情况
		if ll.Head == ll.Head.Next.Next { // 成环
			return true, ll.Head
		} else { //无环
			return false, nil
		}
	}
	slow := ll.Head // 慢指针，一次走一步（重要：slow与fast必须都从头结点开始）
	fast := ll.Head // 快指针，一次走两步

	for {
		if fast == nil { // 无环的结束条件，fast指向nil
			return false, nil
		}
		if fast == slow && fast != ll.Head { // 有环的结束条件,fast与slow在环上相遇
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 有环，继续寻找入环节点
	fast = ll.Head // 快节点重新回到头结点
	for {
		if fast == slow { //
			return true, fast
		}
		fast = fast.Next // 从头结点开始，每次向后移动一位
		slow = slow.Next // 从相遇节点开始，每次向后移动一位
	}
}
