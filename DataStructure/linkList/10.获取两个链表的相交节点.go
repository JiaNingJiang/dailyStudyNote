package linkList

func GetIntersectionPoint(l1, l2 *LinkList) *Node {

	flag1, junction1 := IsHaveCircle(l1)
	flag2, junction2 := IsHaveCircle(l2)

	if !flag1 && !flag2 { // 情况一：两个链表都无环
		// 需要获取两个链表的尾节点
		tail1 := GetTailNodeWithinLinkList(l1)
		tail2 := GetTailNodeWithinLinkList(l2)
		if tail1 != tail2 { // 两条线性链表的尾节点不一样，则必然不相交
			return nil
		} else { // 尾节点是一个，必然相交
			return GetLinearJunction(l1, l2)
		}

	} else if flag1 && flag2 { // 情况二：两个链表都有环
		if junction1 == junction2 { // 入环点是同一个
			return GetLinearJunction(l1, l2) // 相交点在两个链表的线性部分，因此可用此方法求相交点
		} else { // 入环点不是同一个，有两种情况：1.两条链表不相交   2.两条链表有两个交点
			return GetCircleJunction(l1, l2, junction1, junction2)
		}
	} else { // 情况三：一个有环，一个无环(不可能相交)
		return nil
	}

}

// 获取链表的尾节点
func GetTailNodeWithinLinkList(ll *LinkList) *Node {
	if ll.Len == 0 {
		return nil
	}
	current := ll.Head
	for {
		if current.Next == nil {
			return current
		}
		current = current.Next
	}
}

// 相交点在两个链表的线性部分
func GetLinearJunction(l1, l2 *LinkList) *Node {
	l1current := l1.Head
	l2current := l2.Head
	// 第一次循环，让长链表的指针先向后移动 abs(l1.Len - l2.Len)步
	if l1.Len > l2.Len {
		for i := 0; i < l1.Len-l2.Len; i++ {
			l1current = l1current.Next
		}
	} else if l1.Len < l2.Len {
		for i := 0; i < l2.Len-l1.Len; i++ {
			l2current = l2current.Next
		}
	}
	// 第二个循环，找到两个线性链表的交点
	for {
		if l1current == l2current {
			return l1current
		}
		l1current = l1current.Next
		l2current = l2current.Next
	}
}

// 相交点在两个链表的环部分
func GetCircleJunction(l1, l2 *LinkList, j1, j2 *Node) *Node {

	l1current := j1.Next
	for {
		if l1current == j1 { //两个入环点不在同一个环上，说明两链表没有交点
			return nil
		}
		if l1current == j2 {
			return j1 // 两个入环点在同一个环上，说明相交。可以随机返回两个入环节点中的任意一个，因为这两个节点都是相交点。
		}
		l1current = l1current.Next
	}
}
