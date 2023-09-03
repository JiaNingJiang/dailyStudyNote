package lesson8

import (
	"DataStructure2/utils"
	"math"
)

// 返回一颗完全二叉树总的节点个数
func TreeTotalNode(root *utils.Node) int {

	leftDepth := getTreeDepth(root.Left)   // 获取左子树的深度
	rightDepth := getTreeDepth(root.Right) // 获取右子树的深度

	return treeTotalNode(root, leftDepth, rightDepth)
}

func treeTotalNode(root *utils.Node, leftDepth, rightDepth int) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right == nil {
		return 2
	}

	if rightDepth == leftDepth { // 右子树深度 == 左子树深度，意味着左子树必然是满二叉树，但右子树不一定是
		leftTotal := math.Pow(float64(2), float64(leftDepth)) - 1 + 1 // 左子树节点个数+根节点
		// leftTotal + 右子树个数(递归求)
		return int(leftTotal) + treeTotalNode(root.Right, getTreeDepth(root.Right.Left), getTreeDepth(root.Right.Right))
	}
	if rightDepth < leftDepth { // 右子树深度 < 左子树深度，意味着右子树必然是满二叉树，但左子树不一定是
		rightTotal := math.Pow(float64(2), float64(rightDepth)) - 1 + 1 // 右子树节点个数+根节点
		// rightTotal + 左子树个数(递归求)
		return int(rightTotal) + treeTotalNode(root.Left, getTreeDepth(root.Left.Left), getTreeDepth(root.Left.Right))
	}
	panic("二叉树并非完全二叉树") // 右子树深度 > 左子树深度(不可能出现这种情况)
}

// 获取一颗二叉树的深度
func getTreeDepth(root *utils.Node) int {
	depth := 0

	cur := root
	for {
		if cur == nil {
			return depth
		}
		depth++
		cur = cur.Left
	}
}
