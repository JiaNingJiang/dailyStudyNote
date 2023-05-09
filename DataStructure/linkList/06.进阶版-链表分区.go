package linkList

func SeparateLinkListImproved(ll *LinkList, pivot int) {
	// 准备6个指针：1.小于区域头/尾指针 LH/LT  2.等于区域头/尾指针 EH/ET  3.大于区域头/尾指针 GH/GT
	var LT, ET, GT *Node
	LH := NewLinkList()
	EH := NewLinkList()
	GH := NewLinkList()

	_ = GT // 此算法实现中没有用到

	current := ll.Head
	for { // 第一个循环，遍历原始链表，生成三条新链(小于链，等于链，大于链)
		if current == nil {
			break
		}
		flag := NodeDataCompare(current.Data, pivot, TypeInt)
		if flag == -1 { // 加入到小于区域
			LT = LH.TailAdd(current.Data)
		} else if flag == 0 { // 加入到等于区域
			ET = EH.TailAdd(current.Data)
		} else if flag == 1 { // 加入到大于区域
			GT = GH.TailAdd(current.Data)
		}
		current = current.Next
	}
	// 将三条链连接起来
	lessLen := LH.Len
	equalLen := EH.Len
	greatLen := GH.Len

	if lessLen != 0 && equalLen != 0 && greatLen != 0 { // 全部存在
		LT.Next = EH.Head
		ET.Next = GH.Head
		ll.Head = LH.Head // 更新原始链表
	} else if lessLen != 0 && equalLen != 0 { // 只有小于等于
		LT.Next = EH.Head
		ll.Head = LH.Head // 更新原始链表
	} else if lessLen != 0 && greatLen != 0 { // 没有等于
		LT.Next = GH.Head
		ll.Head = LH.Head // 更新原始链表
	} else if equalLen != 0 && greatLen != 0 { // 只有大于等于
		ET.Next = GH.Head
		ll.Head = EH.Head // 更新原始链表
	} else if lessLen != 0 { // 全部小于
		ll.Head = LH.Head // 更新原始链表
	} else if equalLen != 0 { // 全部等于
		ll.Head = EH.Head // 更新原始链表
	} else { // 全部大于
		ll.Head = GH.Head // 更新原始链表
	}

}
