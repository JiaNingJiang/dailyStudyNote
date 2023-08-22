package linkList

// 空间复杂度优化，不再使用栈，空间复杂度为O(1)
func IsNotPalindromeImproved(ll *LinkList) bool {
	if ll.Len < 2 {
		return false
	}
	if ll.Len == 2 {
		if ll.Head.Data == ll.Head.Next.Data {
			return true
		} else {
			return false
		}
	}

	slow := ll.Head
	fast := ll.Head

	for { // 第一次循环，让slow指向链表中点，fast指向尾节点
		if fast == nil || fast.Next == nil { // 第一种情况对应链表长度为偶数，第二种情况对应链表为奇数
			break
		}
		slow = slow.Next      // 慢指针一次走一步
		fast = fast.Next.Next // 快指针一次走两步
	}
	newHead := Recursive_reverse(slow) // 从slow开始，让后续链表反转

	left := ll.Head  // 左区域从原始链表头结点开始
	right := newHead // 右区域从反转后链表的头结点(原本的尾节点)开始
	for {
		if left == slow && right == slow { // 左右指针同时到达slow才算比较结束
			return true
		}
		leftData := left.Data
		rightData := right.Data
		if NodeDataCompare(leftData, rightData, TypeInt) != 0 { // 一旦左右指针指向的节点数据不相等，则返回false
			return false
		}
		if left != slow {
			left = left.Next
		}
		if right != slow {
			right = right.Next
		}
	}
}
