package binaryTree

func FindMinAncestorDP(tree *Tree, n1, n2 *Node) *Node {
	return lowestAncestor(tree.Root, n1, n2)
}

func lowestAncestor(head, n1, n2 *Node) *Node {
	if head == nil { // 当前分支遍历到底也没有找到n1或n2
		return nil
	}
	if head == n1 || head == n2 { // 当前分支遍历到了n1或n2
		return head
	}
	left := lowestAncestor(head.Left, n1, n2)   // 左子树上是否有n1或n2
	right := lowestAncestor(head.Right, n1, n2) // 右子树上是否有n1或n2

	if left != nil && right != nil { // 左右子树上分别是n1和n2，那么当前节点就是n1和n2的最小公共祖先
		return head
	}
	// 左右子树上只能找到n1、n2其中的一个
	if left != nil { // 只有左子树能找到
		return left
	}
	if right != nil { // 只有右子树能找到
		return right
	}
	// n1和n2不在当前分支路线上
	return nil
}
