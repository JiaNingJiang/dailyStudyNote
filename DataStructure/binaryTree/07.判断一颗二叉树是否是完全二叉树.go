package binaryTree

func CheckCBT(tree *Tree) bool {
	lastNonLeaf := false // 如果当前节点为最后一个非叶子结点，此变量变为true

	queue := NewQueue()
	queue.Push(tree.Root)

	for {
		if queue.Len == 0 { // 整颗二叉树都已经完成遍历，依然没有出现任何问题。则必然是完全二叉树
			return true
		}
		data := queue.Pop()
		node := data.(*Node)

		if node.Left == nil && node.Right != nil { // 如果一个节点只有右孩子，没有左孩子。必然不是完全二叉树
			return false
		}
		if lastNonLeaf { // 在前一个节点是最后一个非叶子节点的前提之下，依然拥有左孩子节点。必然不是完全二叉树
			if node.Left != nil {
				return false
			}
		}
		if node.Left != nil && node.Right == nil { // 判断当前节点是否是最后一个非叶子结点
			lastNonLeaf = true
		}
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}

	}
}
