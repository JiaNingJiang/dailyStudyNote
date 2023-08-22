package binaryTree

import "math"

type Info struct {
	currentHeight  int
	currentNodeNum int
}

func CheckFBT(tree *Tree) bool {
	info := checkFBT(tree.Root)
	if info.currentNodeNum == int(math.Pow(2, float64(info.currentHeight)))-1 {
		return true
	} else {
		return false
	}
}

func checkFBT(root *Node) *Info {
	if root == nil {
		return &Info{0, 0}
	}

	leftInfo := checkFBT(root.Left)   // 左子树的层数和节点数
	rightInfo := checkFBT(root.Right) // 右子树的层数和节点数

	currentInfo := new(Info)                                                                // 记录当前子树的层数和节点数
	currentInfo.currentHeight = getMax(leftInfo.currentHeight, rightInfo.currentHeight) + 1 // 加上当前根节点所在层
	currentInfo.currentNodeNum = leftInfo.currentNodeNum + rightInfo.currentNodeNum + 1     // 左子树节点+右子树节点+当前根节点

	return currentInfo
}
