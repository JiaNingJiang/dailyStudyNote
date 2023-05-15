package binaryTree

func FindMinAncestor(tree *Tree, n1, n2 *Node) *Node {

	hashMap := make(map[*Node]*Node, 0)
	hashMap[tree.Root] = tree.Root
	ancestorMapForm(hashMap, tree.Root)

	hashSet := make(map[*Node]struct{}) // 存储从n1到根节点的路径上的所有节点
	currentNode := n1
	for {
		if hashMap[currentNode] == currentNode {
			hashSet[currentNode] = struct{}{} // 最后加入根节点
			break
		}
		hashSet[currentNode] = struct{}{}
		currentNode = hashMap[currentNode]
	}

	// 查询n2是否在hashSet中存在（是否在n1到根节点的路径上）
	currentNode = n2
	for {
		if hashMap[currentNode] == currentNode { // o2路径已经到达根节点
			return currentNode
		}
		if _, ok := hashSet[currentNode]; ok {
			return currentNode
		}
		currentNode = hashMap[currentNode]
	}
}

// 构建父子hash表，key为孩子节点，value是该孩子节点的父节点
func ancestorMapForm(hashMap map[*Node]*Node, root *Node) {
	if root == nil {
		return
	}
	hashMap[root.Left] = root
	hashMap[root.Right] = root

	ancestorMapForm(hashMap, root.Left)
	ancestorMapForm(hashMap, root.Right)
}
