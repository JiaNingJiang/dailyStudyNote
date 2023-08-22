package binaryTree

import "fmt"

// 采用中序遍历的方式
func MricosoftFold(i, N int, down bool) {
	if i > N { // 折纸次数i超过了总次数N
		return
	}
	MricosoftFold(i+1, N, true) // 左子树必然是凹痕

	// 打印当前根节点的折痕（凹痕或凸痕）
	if down {
		fmt.Printf("凹 ")
	} else {
		fmt.Printf("凸 ")
	}

	MricosoftFold(i+1, N, false) // 右子树必然是凸痕
}
