package binaryTree

func GetBinaryTreeMaxWidthImproved(tree *Tree) int {

	queue := NewQueue()
	queue.Push(tree.Root)
	curEnd := tree.Root  // 记录当前层的最后一个节点
	nextEnd := tree.Root // 记录下一层的最后一个节点

	curLevelNode := 0   // 循环变量，当前层的节点数
	max := curLevelNode // 最大宽度

	for {
		if queue.Len == 0 {
			return max
		}
		data := queue.Pop()
		node := data.(*Node)

		if node.Left != nil {
			queue.Push(node.Left)
			nextEnd = node.Left
		}
		if node.Right != nil {
			queue.Push(node.Right)
			nextEnd = node.Right
		}
		if node == curEnd {
			curLevelNode++
			max = getMax(max, curLevelNode)
			curEnd = nextEnd
			curLevelNode = 0
		} else {
			curLevelNode++
		}
	}
}
