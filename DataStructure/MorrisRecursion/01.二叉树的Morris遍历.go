package MorrisRecursion

import (
	"DataStructure/binaryTree"
	"fmt"
	"reflect"
)

const (
	PreOrder int = iota
	InOrder
	PostOrder
)

func Morris(tree *binaryTree.Tree, form int) {
	if tree.Root == nil {
		return
	}

	cur := tree.Root
	mostRight := new(binaryTree.Node)

	for {
		if cur == nil { // 完成遍历
			break
		}
		mostRight = cur.Left  // 根据是否有左子树分成两种情况
		if mostRight != nil { // 有左子树,那么mostRight就指向左子树的左右节点
			for {
				if mostRight.Right == nil { // 意味着这是cur第一次寻找了自己左子树的mostRight
					break
				}
				if mostRight.Right == cur { // 这是cur第二次寻找了自己左子树的mostRight
					break
				}
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil { // 因为第一次找到自己的mostRight而退出的
				mostRight.Right = cur
				if form == PreOrder { // 前序遍历，对于遍历两次的节点，需要在第一次访问的时候进行打印
					fmt.Printf("%d ", cur.Data)
				}
				cur = cur.Left
				continue //  重新寻找当前新节点的mostRight
			}
			if mostRight.Right == cur { // 因为第二次找到自己的mostRight而退出的（此时cur已经遍历完整个左子树，重新回到了根节点位置）
				mostRight.Right = nil
				if form == InOrder { // 中序遍历，对于遍历两次的节点，需要在第二次访问的时候进行打印
					fmt.Printf("%d ", cur.Data)
				}
				if form == PostOrder { // 对于后序遍历，对于遍历两次的节点，需要逆序打印左子树的右边界
					printRightEdge(cur.Left)
				}
				cur = cur.Right
				continue // 遍历右子树
			}
		} else { // 没有左子树,直接遍历当前节点的右子树
			if form == PreOrder || form == InOrder { // 前中序遍历,对于第一次遍历到的节点,直接打印
				fmt.Printf("%d ", cur.Data)
			}
			cur = cur.Right
		}
	}
	if form == PostOrder { // 对于后序遍历，当完成对整棵树的Morris遍历后，需要逆序打印整棵树的右边界
		printRightEdge(tree.Root)
	}
}

// 逆序打印右边界
func printRightEdge(root *binaryTree.Node) {
	tail := reverseEdge(root) // 获取逆序后的右边界的新起点(原本的最右叶子结点)
	cur := tail
	for {
		if cur == nil { // 逆序遍历完整个右边界
			break
		}
		fmt.Printf("%d ", cur.Data)
		cur = cur.Right // 指向原本的父节点，现在的右孩子节点
	}
	reverseEdge(tail) // 重要：恢复之前的顺序
}

func reverseEdge(root *binaryTree.Node) *binaryTree.Node {
	pre := new(binaryTree.Node)
	cur := root
	next := new(binaryTree.Node)

	for {
		if cur == nil {
			break
		}
		next = cur.Right                  // 实现保留自己的右孩子的地址
		if reflect.DeepEqual(cur, root) { // 重要：右边界的起点，必须需要让其右孩子指针为nil，而不能是Node{}
			cur.Right = nil
		} else {
			cur.Right = pre // 自己的右孩子指针指向父节点
		}
		pre = cur  // 记录当前节点(下一次循环中成为自己孩子节点的右孩子节点)
		cur = next // 更新循环变量
	}
	return pre // pre最终指向最右叶子结点
}
