package binaryTree

var PreValue interface{} // 初始要设置为math.MinInt，也就是要小于BST最小节点的数值

func CheckBST(root *Node) bool {
	if root == nil { // 边界条件。当前子树为空，返回true
		return true
	}
	leftBST := CheckBST(root.Left) // 检查左子树是否是BST
	if !leftBST {
		return false
	}
	// 确保当前子树根节点大于左子树的最右侧节点（PreValue保存）
	currentData := root.Data.(int)
	preData := PreValue.(int)
	if currentData < preData {
		return false
	} else {
		PreValue = root.Data // 确保成功，更新PreValue为当前子树根节点值，以便与右子树的最左节点进行比较
	}
	return CheckBST(root.Right) // 检查右子树是否是BST,以及右子树的最左侧节点是否大于子树的根节点
}
