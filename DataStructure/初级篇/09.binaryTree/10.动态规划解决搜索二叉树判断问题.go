package binaryTree

import "math"

type BSTInfo struct {
	IsBST bool
	Max   int
	Min   int
}

func CheckBSTByDP(tree *Tree) bool {
	info := checkBSTByDP(tree.Root)
	return info.IsBST
}

func checkBSTByDP(root *Node) *BSTInfo {
	if root == nil {
		return &BSTInfo{true, math.MinInt, math.MinInt} // 重要：这里将min设置为math.MinInt其实是不合适的，因为对于叶子节点会直接出现 根节点 < 右子树.Min
	}
	leftInfo := checkBSTByDP(root.Left)
	rightInfo := checkBSTByDP(root.Right)

	greatLeft := leftInfo.Max <= root.Data.(int)  // 根节点是否大于左子树全部节点
	lessRight := root.Data.(int) <= rightInfo.Min // 根节点是否小于右子树全部节点

	if leftInfo.Max == math.MinInt { // 说明没有左孩子结点
		greatLeft = true
		leftInfo.Min = root.Data.(int)
		leftInfo.Max = root.Data.(int)
	}

	if rightInfo.Min == math.MinInt { // 重要：说明没有右孩子节点（此步非常关键，用来弥补root == nil时，Min被设置成math.MinInt的缺陷）
		lessRight = true
		rightInfo.Min = root.Data.(int)
		rightInfo.Max = root.Data.(int)
	}

	currentInfo := new(BSTInfo)
	if leftInfo.IsBST && rightInfo.IsBST && greatLeft && lessRight {
		currentInfo.IsBST = true
	} else {
		currentInfo.IsBST = false
	}
	currentInfo.Max = getMax(getMax(leftInfo.Max, rightInfo.Max), root.Data.(int))
	currentInfo.Min = getMin(getMin(leftInfo.Min, rightInfo.Min), root.Data.(int))

	return currentInfo
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
