package binaryTree

func FindSuccessor(tree *Tree, n *Node) *Node {

	// 情况一：当前节点有右子树，那么后继节点就是右子树的最左节点
	if n.Right != nil {
		return findSuccessor1(n.Right)
	}

	// 情况二：当前节点没有右子树，仅当当前节点位于某祖先节点的的左子树上时，该祖先节点是节点的后继节点；其他情况下节点无后继节点
	hashMap := make(map[*Node]*Node, 0)
	hashMap[tree.Root] = tree.Root
	ancestorMapForm(hashMap, tree.Root) // 构建父子关系表

	return findSuccessor2(n, hashMap)

}

func findSuccessor1(n *Node) *Node {
	var current *Node = n
	for {
		if current.Left == nil {
			return current
		}
		current = current.Left
	}
}

func findSuccessor2(n *Node, ancestorMap map[*Node]*Node) *Node {
	current := n
	for {
		if ancestorMap[current] == current { // 已经遍历到根节点，说明节点n不在任何祖先节点的左子树上
			return nil
		}
		if ancestorMap[current].Left == current { // 当前节点是其父节点的左孩子，那么返回其父节点
			return ancestorMap[current]
		}
		current = ancestorMap[current]
	}
}
