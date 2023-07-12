package lesson7

import "DataStructure2/utils"

// 将一颗搜索二叉树转化为双向链表，返回链表的头结点和尾节点
func BSTtoLinkList(root *utils.Node) (*utils.Node, *utils.Node) {

	return bSTtoLinkList(root)
}

func bSTtoLinkList(root *utils.Node) (*utils.Node, *utils.Node) {
	if root == nil {
		return nil, nil
	}
	leftHead, leftTail := bSTtoLinkList(root.Left)    // 将左子树变成双向链表
	rightHead, rightTail := bSTtoLinkList(root.Right) // 将右子树变成双向链表

	if leftTail != nil { // 左子树变的双向链表不为空
		leftTail.Right = root
		root.Left = leftTail
	}
	if rightHead != nil { // 右子树变的双向链表不为空
		root.Right = rightHead
		rightHead.Left = root
	}

	curHead := root
	curTail := root
	if leftHead != nil { // 左子树变的双向链表不为空,那么左链表的头结点就是整个链表的头结点（默认是当前节点）
		curHead = leftHead
	}
	if rightTail != nil { // 右子树变的双向链表不为空,那么右链表的尾结点就是整个链表的尾结点（默认是当前节点）
		curTail = rightTail
	}

	return curHead, curTail
}
