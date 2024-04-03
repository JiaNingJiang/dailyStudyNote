## 一、题目一

给你一个字符串类型的数组arr，比如：

string[] arr = {  "`b` \\\\ `cst`", "`d` \\\\", "a\\\\d\\\\e" , "a\\\\b\\\\c"  }

需要你把这些路径蕴含的目录结构画出来，子目录直接列在父目录下面，并比父目录向右进两个，就像下面这样：

```
a
  b
    c
  d
    e
b
  cst
d    
```

同一级的需要按照字母顺序排列，不能乱。

```go
1.先将整个字符串数组构建为前缀树的形式（为了方便打印，让每一个节点再额外存储一个字符串）
2.打印的时候，按照深度优先遍历算法进行打印。每一层在打印字符前需要先打印 (layer-1)*2个空格
```

```go
package lesson7

import (
	"fmt"
	"strings"
)

type Node struct {
	pass int              // 经过此节点的路径数
	end  int              // 以此节点为终点的路径数
	data string           // 为了方便打印，让节点存储其头顶路径的内容
	next map[string]*Node // 当前节点的 下一 节点
}

// 返回一颗前缀树的根节点
func NewPrefixTree() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

// 向前缀树中插入一条文件路径
func (ptree *Node) Insert(file string) {
	filePath := strings.Split(file, "\\\\")
	if filePath[len(filePath)-1] == "" {
		filePath = filePath[:len(filePath)-1]
	}

	ptree.pass++     // 所有的路径都起源于前缀树根节点
	curNode := ptree // 从根节点出发，拓展前缀树

	for _, dir := range filePath {
		if _, ok := curNode.next[dir]; !ok { // 下一个节点不存在，则需要创建
			curNode.next[dir] = &Node{next: make(map[string]*Node)}
		}
		curNode = curNode.next[dir] // 跳转到下一个节点
		curNode.data = dir
		curNode.pass++
	}
	curNode.end++
}

// 采用深度优先遍历方式打印整颗前缀树
func (ptree *Node) DFSPrintFilePath() {
	ptree.dFSPrintFilePath(0)
}

func (ptree *Node) dFSPrintFilePath(layer int) {
	if ptree == nil {
		return
	}

	for _, child := range ptree.next {
		dir := child.data
		// 1.先打印空格
		for i := 0; i < 2*layer; i++ {
			fmt.Printf(" ")
		}
		// 2.再打印孩子节点的内容(也就是目录名)
		fmt.Println(dir)
		child.dFSPrintFilePath(layer + 1) // 深度优先遍历的方式继续往下打印
	}
}
```



## 二、题目二

双向链表节点结构和二叉树节点结构是一样的，如果你把last认为是left，next认为是right的话。

给定一个搜索二叉树的头结点head，请转化为一条有序的双向链表，并返回链表的头结点和尾节点。

```go
采用递归的方法：
对于当前节点X：
1.递归处理当前节点的左子树，将其转化为一个双向链表，并返回转化后的头和尾节点
2.递归处理当前节点的右子树，将其转化为一个双向链表，并返回转化后的头和尾节点
3.左子树的尾节点的right指向当前节点X,当前节点X的left指向左子树的尾节点
4.右子树的头节点的left指向当前节点X,当前节点X的right指向右子树的头节点
5.如果左子树不为空，链表头结点就是左子树头结点；否则链表头结点为当前节点X
6.如果右子树不为空，链表尾结点就是右子树尾结点；否则链表尾结点为当前节点X
```

```go
package lesson7

import "DataStructure2/utils"

// 将一颗搜索二叉树转化为双向链表，返回链表的头结点和尾节点
func BSTtoLinkList(root *utils.Node) (*utils.Node, *utils.Node) {

	return bSTtoLinkList(root)
}

func bSTtoLinkList(root *utils.Node) (*utils.Node, *utils.Node) {
	if root == nil {
		return nil, nil
	}
	leftHead, leftTail := bSTtoLinkList(root.Left)    // 将左子树变成双向链表
	rightHead, rightTail := bSTtoLinkList(root.Right) // 将右子树变成双向链表

	if leftTail != nil { // 左子树变的双向链表不为空
		leftTail.Right = root
		root.Left = leftTail
	}
	if rightHead != nil { // 右子树变的双向链表不为空
		root.Right = rightHead
		rightHead.Left = root
	}

	curHead := root
	curTail := root
	if leftHead != nil { // 左子树变的双向链表不为空,那么左链表的头结点就是整个链表的头结点（默认是当前节点）
		curHead = leftHead
	}
	if rightTail != nil { // 右子树变的双向链表不为空,那么右链表的尾结点就是整个链表的尾结点（默认是当前节点）
		curTail = rightTail
	}

	return curHead, curTail
}
```



## 三、题目三

找到一颗二叉树中，最大的搜索二叉子树，返回最大搜索二叉子树的头结点。

```go
最大的二叉搜索子树，可能的来源：
1.来自当前节点的左子树
2.来自当前节点的右子树
3.以当前节点为根节点的整颗二叉树(左子树和右子树都是BST，而且左子树最大值 < 根节点 < 右子树最小值)

因此，每个节点需要返回的递归变量右5个：（用一个Info结构体包含）
1.当前子树是否是搜索二叉树 isBST
2.当前子树的最大值max
3.当前子树的最小值min
4.最大搜索二叉子树的头结点head
5.最大搜索二叉子树的节点个数count

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
```

## 四、题目四

假设现在有一个数组用来记录某一个帖子的得分记录，我们要求解出在这个得分记录中连续打分数据之和的最大值，这个值被认为是该帖子的最高分。

假设某帖子近期的打分记录为：[1 1 -1 -10 11 4 -6 9 20 -10 -2]，那么该帖子曾经达到过的最高分数为：

11+ 4 + (-6) + 9 +20 =38。

```go
子数组累加和最大的可能性不止一种,长度不同的子数组累加和也可能是相同的，这里我们假设我们要求解的是：累加和最大且最长的子数组。
假设arr[0……N-1]数组最终找到的是 i~j 的子数组符合条件，那么会有下面的结论：
1. i~k(k<=j) 这一段位置上的累加和不可能小于0，因为如果小于零，i~j位置就不可能是累加和最大的，k~j的子数组累加和明显会更大
2. m(m>=0) ~i-1 的子数组累加和不可能 >=0。因为如果>0，i~j位置就不可能是累加和最大的；如果==0,i~j位置就不可能是最长的

根据以上特点，总结出求解代码为：

package lesson7

import "math"

func PostMaxScore(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	maxScore := math.MinInt

	curScore := 0
	for i := 0; i < len(scores); i++ {
		curScore += scores[i]
		if curScore > maxScore {
			maxScore = curScore
		}
		if curScore < 0 { // 如果连续分数之和 < 0, 那么下一次从零继续开始
			curScore = 0
		}
	}
	return maxScore
}

```

## 五、题目五

给定一个整型矩阵，返回子矩阵的最大累加和。

假设矩阵为：

| 第0行 |  -5  |  3   |  6   |  4   |
| :---: | :--: | :--: | :--: | :--: |
| 第1行 |  -7  |  9   |  -5  |  3   |
| 第2行 | -10  |  1   | -200 |  4   |

```go
求解的思路是：
先求解0~0行的子矩阵，再求0~1行 ，再求0~2行，再求1~1行，再求1~2行，再求2~2行
1.0~0行： -5 3 6 4 求最大子矩阵累加和，相当于求解数组的最大累加和，采取题目四的算法求解
2.0~1行： 将矩阵压缩，每一列的元素相加得到一个压缩值： -12 12 1 7，采取题目四的算法求解数组的最大累加和就是0~1行的最大子矩阵累加和
3.0~2行：压缩+题目四算法：  -22 12 -199 11  求最大累加和
…………………………

比较每一种情况下的最大累加和，得出最终的最大累加和
```

```go
package lesson7

func MaxMatrixSum(matrix [][]int) int {
	rowCount := len(matrix)    // 行数
	colCount := len(matrix[0]) // 列数

	res := 0
	compressRow := make([]int, colCount)                 // 压缩行
	for startRow := 0; startRow < rowCount; startRow++ { // 起始行号
		for endRow := startRow; endRow < rowCount; endRow++ { // 终止行号
			// 1.进行列压缩
			for col := 0; col < colCount; col++ {
				for row := startRow; row <= endRow; row++ {
					compressRow[col] += matrix[row][col]
				}
			}
			// 2.获取压缩行的最大累加和
			compressMax := PostMaxScore(compressRow)
			// 3.更新子矩阵最大累加和
			res = getMax(res, compressMax)
			// 4.重新清空压缩行
			compressRow = make([]int, colCount)
		}
	}
	return res
}
```

