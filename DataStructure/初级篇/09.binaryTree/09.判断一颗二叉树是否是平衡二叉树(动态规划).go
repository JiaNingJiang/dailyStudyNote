package binaryTree

import "math"

type AVLInfo struct {
	IsBalance bool // 是否平衡
	Height    int  // 当前子树高度
}

func CheckAVL(tree *Tree) bool {
	info := checkAVL(tree.Root)
	return info.IsBalance
}

func checkAVL(root *Node) *AVLInfo {
	if root == nil {
		return &AVLInfo{true, 0}
	}

	leftInfo := checkAVL(root.Left)   // 递归左子树
	rightInfo := checkAVL(root.Right) // 递归右子树

	currentInfo := new(AVLInfo)
	currentInfo.Height = getMax(leftInfo.Height, rightInfo.Height) + 1
	if leftInfo.IsBalance && rightInfo.IsBalance && math.Abs(float64(leftInfo.Height-rightInfo.Height)) < 2 {
		currentInfo.IsBalance = true
	} else {
		currentInfo.IsBalance = false
	}
	return currentInfo
}
