package treeTypeDP

import "DataStructure/09.binaryTree"

type Info struct {
	maxHeight   int // 当前子树的高度
	maxDistance int // 当前子树记录的最大距离
}

func MaxDistance(tree *binaryTree.Tree) int {
	return maxDistance(tree.Root).maxDistance
}

func maxDistance(root *binaryTree.Node) Info {
	if root == nil { // 叶子结点的下一层，即不存在的节点
		return Info{0, 0}
	}
	leftInfo := maxDistance(root.Left)
	rightInfo := maxDistance(root.Right)
	twoSideDis := leftInfo.maxHeight + rightInfo.maxHeight + 1

	maxDis := getMax(twoSideDis, getMax(leftInfo.maxDistance, rightInfo.maxDistance))
	maxH := getMax(leftInfo.maxHeight, rightInfo.maxHeight) + 1

	return Info{maxHeight: maxH, maxDistance: maxDis}
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
