package lesson7

import "DataStructure2/utils"

type Info struct {
	IsBST    bool        // 是否是BST
	MaxV     int         // 当前子树的最大值
	MinV     int         // 当前子树的最小值
	Head     *utils.Node // 最大BST的头结点
	BSTCount int         // 最大BST上的节点数
}

func MaxBST(root *utils.Node) *Info {
	res := maxBST(root)
	//if res != nil {
	//	if res.Head != nil {
	//		return res.Head
	//	}
	//}
	//return nil
	return res
}

func maxBST(root *utils.Node) *Info {
	if root == nil {
		return nil
	}

	leftRes := maxBST(root.Left)
	rightRes := maxBST(root.Right)

	resInfo := new(Info)
	curV := root.Data.(int)

	if leftRes == nil && rightRes == nil { // 左右都为空节点
		resInfo.IsBST = true
		resInfo.MaxV = root.Data.(int)
		resInfo.MinV = root.Data.(int)
		resInfo.BSTCount = 1
		resInfo.Head = root
		return resInfo
	} else if leftRes == nil { // 只有左子树为空节点,则完全取决于右子树
		if curV < rightRes.MinV {
			resInfo.IsBST = true
			resInfo.BSTCount = rightRes.BSTCount + 1
			resInfo.Head = root
		} else {
			resInfo.IsBST = false
			resInfo.BSTCount = rightRes.BSTCount
			resInfo.Head = rightRes.Head
		}

		resInfo.MaxV = getMax(rightRes.MaxV, curV)
		resInfo.MinV = getMin(rightRes.MinV, curV)

		return resInfo
	} else if rightRes == nil { // 只有右子树为空节点,则完全取决于左子树
		if curV > leftRes.MaxV {
			resInfo.IsBST = true
			resInfo.BSTCount = leftRes.BSTCount + 1
			resInfo.Head = root
		} else {
			resInfo.IsBST = false
			resInfo.BSTCount = leftRes.BSTCount
			resInfo.Head = leftRes.Head
		}

		resInfo.MaxV = getMax(leftRes.MaxV, curV)
		resInfo.MinV = getMin(leftRes.MinV, curV)

		return resInfo
	}
	// 剩下的只有左右子树都不为空的情况

	if leftRes.IsBST && rightRes.IsBST { // 左子树和右子树都是BST(最为理想特殊的情况)
		if curV > leftRes.MaxV && curV < rightRes.MinV { // 左子树最大值 < 当前节点值 < 右子树最小值 --> 整棵树都是BST
			resInfo.IsBST = true
			resInfo.MaxV = rightRes.MaxV
			resInfo.MinV = leftRes.MinV
			resInfo.Head = root
			resInfo.BSTCount = leftRes.BSTCount + rightRes.BSTCount + 1
		} else { // 整棵树构不成BST
			resInfo.IsBST = false
			if leftRes.BSTCount >= rightRes.BSTCount {
				resInfo.MaxV = leftRes.MaxV
				resInfo.MinV = leftRes.MinV
				resInfo.BSTCount = leftRes.BSTCount
				resInfo.Head = leftRes.Head
			} else {
				resInfo.MaxV = rightRes.MaxV
				resInfo.MinV = rightRes.MinV
				resInfo.BSTCount = rightRes.BSTCount
				resInfo.Head = rightRes.Head
			}
		}
		return resInfo
	}
	if leftRes.BSTCount >= rightRes.BSTCount { // 左子树的BST具有更多的节点
		resInfo.IsBST = false
		resInfo.MaxV = leftRes.MaxV
		resInfo.MinV = leftRes.MinV
		resInfo.BSTCount = leftRes.BSTCount
		resInfo.Head = leftRes.Head
	} else { // 右子树的BST具有更多的节点
		resInfo.IsBST = false
		resInfo.MaxV = rightRes.MaxV
		resInfo.MinV = rightRes.MinV
		resInfo.BSTCount = rightRes.BSTCount
		resInfo.Head = rightRes.Head
	}
	return resInfo
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
