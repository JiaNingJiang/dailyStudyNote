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
			if l1.Len > l2.Len {
				
			}
		}

	} else if flag1 && flag2 { // 情况二：两个链表都有环

	} else { // 情况三：一个有环，一个无环(不可能相交)
		return nil
	}

}

func GetTailNodeWithinLinkList(ll *LinkList) *Node {
	if ll.Len == 0 {
		return nil
	}
	if ll.Len == 1 {
		return ll.Head
	}
	current := ll.Head
	for {
		if current.Next == nil {
			return current
		}
		current = current.Next
	}
}
