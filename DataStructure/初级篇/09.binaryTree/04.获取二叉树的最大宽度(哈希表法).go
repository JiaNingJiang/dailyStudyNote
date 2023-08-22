package binaryTree

func GetBinaryTreeMaxWidth(tree *Tree) int {
	hashMap := make(map[*Node]int) // 记录每个节点在第几层

	queue := NewQueue()
	queue.Push(tree.Root)
	hashMap[tree.Root] = 1

	currentLevel := 1   // 循环变量，当前所在层数(1~n)
	curLevelNode := 0   // 循环变量，当前层的节点数
	max := curLevelNode // 最大宽度

	for {
		if queue.Len == 0 {
			max = getMax(max, curLevelNode)
			break
		}
		data := queue.Pop()
		node := data.(*Node)

		if hashMap[node] != currentLevel { // 弹出的节点在下一层
			max = getMax(max, curLevelNode) // 更新最大宽度(总结上一层的节点数)
			currentLevel++                  // 层数++
			curLevelNode = 1
		} else { // 弹出的节点还在当前层
			curLevelNode++
		}

		if node.Left != nil {
			queue.Push(node.Left)
			hashMap[node.Left] = currentLevel + 1 // 孩子节点肯定在下一层
		}
		if node.Right != nil {
			queue.Push(node.Right)
			hashMap[node.Right] = currentLevel + 1 // 孩子节点肯定在下一层
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
