package lesson8

// 根据二叉树的先序和中序序列获取后序序列
func GetPostOrder(pre, in []int) []int {

	length := len(pre)
	post := make([]int, length)

	getPostOrder(pre, in, &post, 0, length-1, 0, length-1, 0, length-1)

	return post
}

func getPostOrder(pre, in []int, post *[]int, preS, preE, inS, inE, postS, postE int) {
	if preS > preE { // 递归到最大深度，先序序列无法分得更小
		return
	}
	if preS == preE { // 如果当前先序序列只剩一个节点，那么这个节点必然是子树的根节点
		(*post)[postE] = pre[preS]
		return
	}
	(*post)[postE] = pre[preS] // 后序序列最后一个节点和先序序列的第一个节点是一样的，都是二叉树的根节点

	inRootIndex := inS // 找到根节点在中序序列中的位置
	for ; inRootIndex <= inE; inRootIndex++ {
		if in[inRootIndex] == pre[preS] {
			break
		}
	}

	leftArea := inRootIndex - inS // 左子树节点个数
	// rightArea := inE - inRootIndex // 右子树节点个数

	// 基于先序和中序设置后序的左子树区域
	getPostOrder(pre, in, post, preS+1, preS+leftArea, inS, inRootIndex-1, postS, postS+leftArea-1)

	// 基于先序和中序设置后序的右子树区域
	getPostOrder(pre, in, post, preS+leftArea+1, preE, inRootIndex+1, inE, postS+leftArea, postE-1)

}
