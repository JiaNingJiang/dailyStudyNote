package linkList

// 直接在原始链表上操作。与其说是复制，不如说是剪切
func CopyLinkListWithRandomPImproved(sl *SLinkList) {
	if sl.Len == 0 {
		return
	}

	current := sl.Head
	for { // 第一个循环，为链表的每一个节点创建克隆节点(只有Data字段被复制),并将该克隆节点插入到自己的后继节点之前
		if current == nil {
			break
		}
		snode := new(RNode)       // 复制当前链表节点
		snode.Data = current.Data // 第一次循环仅复制Data字段

		oldNext := current.Next // 原始链表中当前节点的后继节点
		current.Next = snode    // 更改当前节点的后继节点为克隆节点
		snode.Next = oldNext    // 克隆节点的后继节点为oldNext

		sl.Len++

		current = oldNext // 循环变量指向原后继节点oldNext
	}

	index1 := sl.Head      // 指向奇数节点，也就是原链表上的节点
	index2 := sl.Head.Next // 指向偶数节点，也就是克隆出来的节点
	for {                  // 第二个循环，让克隆节点完成random指针的克隆
		if index2.Next == nil {
			break
		}
		index2.Random = index1.Random.Next // 重点：注意这里是 = index1.Random.Next，而不是index1.Random
		index1 = index1.Next.Next
		index2 = index2.Next.Next
	}

	// 第三次循环，如果是奇数节点则从链表删除
	current = sl.Head
	index := 1 // 区分奇偶节点
	for {
		if current == nil { // 由于循环变量一次递增两次，因此这里是current == nil ,而不是current.Next == nil
			break
		}
		if index == 2 { // 跳过第一个奇数节点
			sl.Head = current
			sl.Len--
		}
		if index%2 == 0 { // 跳过奇数节点
			current.Next = current.Next.Next
			sl.Len--
		}
		current = current.Next // 重要：这里实际上是一次跳跃了两个
		index++
	}
}
