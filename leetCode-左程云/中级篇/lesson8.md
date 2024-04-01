## 一、题目一

为一条长度为n的道路设计路灯安置方案，可以把道路视为n个方格，需要照亮的地方用`.`表示，不需要照亮的障碍物格子用`X`表示。小Q现在要在道路上设置一些路灯，对于安置在pos位置的路灯，这站路灯可以照亮pos-1,pos,pos+1这三个位置。

小Q希望能安置尽量少的路灯照亮所有`.`区域，请计算一下最少需要多少盏路灯。

```go
// area是只包含`.`和`X`的字符串
// 求解这个问题可以用贪心策略
package lesson8

func MinLight(area string) int {
	index := 0
	minLight := 0
	for {
		if index >= len(area) {
			return minLight
		}
		if area[index] == 'x' { // 当前位置是路障,则直接跳过
			index++
		} else { // 当前位置是人行道
			minLight++                // 该区域必然会有一个灯(具体位置可以是index或者是index+1，取决于下一个位置是否也是人行道)
			if index+1 == len(area) { // 当前位置就是最后一块人行道，则路灯就安放在当前index位置
				return minLight
			}
			if area[index+1] == 'x' { // 下一个位置是路障，则路灯就安放在当前index位置，然后跳到index+2位置
				index += 2
				continue
			} else { // 下一位置还是人行道，则路灯安放在index+1位置，然后跳到index+3位置
				index += 3
			}
		}
	}
}

```

## 二、题目二

已知一颗完全二叉树的先序跟中序遍历结果（二叉树中每一个元素都是不一样的），请求出后序遍历结果。

```go
先序遍历(根左右)：
[ 根节点  (左子树)   (右子树)  ]

中序遍历(左根右)：
[ (左子树)  根节点   (右子树)  ]

后序遍历(左右根)：
[ (左子树)   (右子树)  根节点  ]

1.先序遍历的第一个元素必然是整颗二叉树的根节点。
2.中序遍历序列根节点将左右子树划分为两个区域，这样左右子树各自的节点数目就已知了
3.根据左右子树各自的节点数目，可以知道先序遍历序列中左右子树各自的区域

func GetPostSet(pre,in []int) []int {
    post := make([]int,len(pre))
    length := len(pre)
    
    set(pre,in,post,0,length,0,length,0,length)
    return post
}

func set(pre,in,post []int,prei,prej,ini,inj,posi,posj int) {
    if prei > prej {
        return 
    }
    if prei == prej {  // 先序遍历序列的左子树只剩下一个数
        pos[posi] = pre[prej]
    }
    pos[posj] = pre[prei]  // 根节点必然在后序遍历序列的最后一个位置
    find := ini
    for ;find<=inj;find++ {  // 在中序遍历序列中找到根节点
        if in[find] == pre[prei] {
            break
        } 
    }
    // 填充后序遍历序列的左子树区域(区域大小:find-ini)
    // 前序遍历序列左子树区域：prei+1 ~ prei+find-ini （prei是根节点）
    // 中序遍历序列左子树区域：ini ~ find-1 (find是根节点)
    // 后序遍历序列左子树区域：posi ~ posi+find-ini-1
    set(pre,in,pos,prei+1,prei+find-ini,ini,find-1,posi,posi+find-ini-1)
    
    // 填充后序遍历序列的右子树区域(区域大小:inj - find)
    // 前序遍历序列右子树区域：prei+find-ini+1 ~ prej 
    // 中序遍历序列右子树区域：find+1 ~ inj (find是根节点)
    // 后序遍历序列右子树区域：posi+find-ini ~ posj-1 (posj是根节点)
    set(pre,in,pos,prei+find-ini+1,prej,find+1,inj,posi+find-ini,posj-1)
}
```

```go
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
```

优化策略：

```go
每次要通过遍历的方式从中序序列中找到当前(子)树的根节点，时间复杂度是O(N)。
可以预先建立一个hashMap,Map中存储中序序列中各种值在序列中的下标，将查找根节点的时间复杂度变为O(1)
```



## 三、题目三

返回一棵完全二叉树的节点总个数？

```go
1.获取根节点X
2.获取X的左子树的深度D(不断访问左孩子节点即可获取)
3.获取X的右子树的深度d
1) d == D,说明X的左子树是一颗满二叉树，左子树节点个数 + X = 2^D - 1 + 1 = 2^D。但是X的右子树不一定是满二叉树，此时将右子树的根节点作为新的X进行新一轮递归。
2) d < D,说明X的右子树是一颗满二叉树，右子树节点个数 + X = 2^d。但是X的左子树不一定是满二叉树，此时将左子树的根节点作为新的X进行新一轮递归。
3）边界条件，当根节点X没有右子树时，意味着X的左子树必定只有一个节点，此时返回节点数为2即可（根节点+左孩子结点）
```

```go
// h是整颗完全二叉树的深度(是固定的全局变量)
// level表示当前的node在第几层
// 返回值表示以node为根节点的完全二叉树，节点的个数
func bs(node *Node,level int,h int) int {
    if level == h {  // 最后一层
        return 1
    }
    if mostLeftLevl(node.right,level+1) == h {  // 右子树的最左节点能到达最底层
        return (1 << (h - level)) + bs(node.right,level+1,h) //左子树和根节点的节点总个数 + 右子树的节点总个数(递归)
    } else {   // 右子树的最左节点不能到达最底层
        return (1 << (h - level - 1)) + bs(node.left,level+1,h) // 右子树和根节点的节点个数 + 左子树的节点个数(递归)
    }
}

// 如果node在第level层
// 返回以node为根节点的子树(该子树必须是完全二叉树)的最深层数
func mostLeftLevel(node *Node,level int) int {
    for {
        if node == nil {
            return level-1
        }
        level++
        node = node.left
    }
}

// 时间复杂度是：O(（logN）^2) : 一共需要遍历 logN层，每一层要遍历最多logN个节点(每轮遍历的节点数-1)
```

```go
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
```



## 四、题目四

求一个整形数组的最长递增子序列的长度

1. 方法一：动态规划 (时间复杂度为 O(N^2) )

```go
假设原始整型数组为 arr = [3 1 2 6 3 4 0]
其最长递增子序列应该是 [1 2 3 4]

1.准备一个与 arr 长度相等的辅助数组 dp = []
2.开始遍历数组arr
3.arr[0] = 3, 3左侧没有比它更小的数字，因此dp[0] = 1
4.arr[1] = 1, 1左侧没有比它更小的数字，因此dp[1] = 1
5.arr[2] = 2, 2左侧有比它小的数字，从左侧选取dp值最大的一个，因为这里只有arr[1]比它小，那么就选择arr[1], 此时dp[2] = 1 + dp[1] = 2
6.arr[3] = 6, 6左侧比它小且dp累计值最大的是arr[2]，因此dp[3] = 1 + dp[2] =3
7.arr[4] = 3, 3左侧比它小且dp累计值最大的是arr[2]，因此dp[4] = 1 + dp[2] =3
7.arr[5] = 4, 4左侧比它小且dp累计值最大的是arr[4]，因此dp[5] = 1 + dp[4] =4

因此，最长递增子序列长度为4
```

2. 优化版 (时间复杂度为 O(`N` * `logN` ) )

```go
额外准备一个ends数组，长度和原始数组arr以及dp数组相等。
ends[i]负责记录所有长度为i+1(因为下标从0开始，而长度从1开始)的递增子序列中数值最小的结尾整数
ends起始阶段所有位都是无效的(无效的位是不能访问的)，每次遍历arr一位，都会对ends进行更新

arr = [3 1 2 6 3 4 0]

1.arr[0] = 3, 此时ends所有位都是无效位，ends = [3 …… ]， dp[0] = ends数组中3以及3左侧元素个数之和，dp[0] = 1

2.arr[1] = 1, 此时在ends有效位置中寻找比1大的最左侧的元素（ 采用二分的方式查找ends数组的有效区域，时间复杂度是O(logN) ），结果是3，那么把3更新为1（更新的目的是为了方便arr[1]后续的元素查找递增子序列，对arr[1]之前的元素没有任何作用），即ends = [1 …… ],  dp[1] = ends数组中1以及1左侧元素个数之和，dp[1] = 1

3.arr[2] = 2, 此时在ends有效位置中寻找比2大的最左侧的元素，发现没有，那么扩充ends，即ends = [1 2 …… ]， dp[2] = ends数组中2以及2左侧元素个数之和，dp[2] = 2

4.arr[3] = 6, 此时在ends有效位置中寻找比6大的最左侧的元素，发现没有，那么扩充ends，即ends = [1 2 6 …… ]， dp[3] = ends数组中6以及6左侧元素个数之和，dp[3] = 3

5.arr[4] = 3, 此时在ends有效位置中寻找比3大的最左侧的元素, 发现是6，那么更新ends，即
ends = [1 2 3 ……]， dp[4] = ends数组中3以及3左侧元素个数之和，dp[4] = 3

6.arr[5] = 4, 此时在ends有效位置中寻找比4大的最左侧的元素, 发现没有，那么扩充ends，即ends = [1 2 3 4 …… ]， dp[5] = ends数组中4以及4左侧元素个数之和，dp[5] = 4

时间复杂度：
外循环，遍历arr数组 O(N)
内循环，二分法遍历ends数组有效区域 O(logN)
因此总的时间复杂度就是 O(N*logN)
```

```go
package lesson8

func LongestIncrease(order []int) int {
	if len(order) == 0 {
		return 0
	}
	dp := make([]int, len(order))

	for i := 0; i < len(order); i++ {
		// 1.在order[i]之前寻找比起更小的数
		maxDP := 0 // 记录比order[i]更小数字的最大dp累计值
		for j := 0; j < i; j++ {
			if order[j] < order[i] {
				maxDP = getMax(maxDP, dp[j])
			}
		}
		// 2.当前i位置的dp值 == 1 + maxDP(可以是0)
		dp[i] = 1 + maxDP
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = getMax(res, dp[i])
	}

	return res
}
```



## 五、题目五

小Q有一组神奇的数列：1,12,123，……，12345678910,1234567891011，……

小Q对于能否被3整除这个性质很感兴趣。

小Q现在希望你能帮助他计算这个数列的第i个到第j个中有多少个可以被3整除。

比如：输入2，5（12，123,1234,12345）   输出是 3 （12,123，12345都能整除3）

```go
判断一个数能否被3整除，等价于一个数的每位之和能否被3整除。
比如：判断12345678910是否能够整除3，等价于判断 (1+2+3+4+5+6+7+8+9+1+0) 能否整除3
```

```go
package lesson8

import (
	"fmt"
	"strconv"
)

// 生成一组神奇的数列
func magicSeries(start, end int) []int {
	series := make([]int, 0)
	for i := start; i <= end; i++ {
		str := ""
		for j := 1; j <= i; j++ {
			str += fmt.Sprintf("%d", j)
		}
		number, _ := strconv.Atoi(str)
		series = append(series, number)
	}
	return series
}

func MagicSeriesDiv3(start, end int) int {
	series := magicSeries(start, end)

	count := 0

	for _, num := range series {
		if isNotDivBy3(num) {
			fmt.Printf("%d可以被3整除\n", num)
			count++
		}
	}

	return count
}

// 判断一个数是否能被3整除
func isNotDivBy3(num int) bool {

	sum := 0
	cur := num
	for {
		if cur/10 == 0 { // 只剩下一位
			sum += cur
			break
		}
		val := cur % 10 // 获得最低位
		sum += val

		cur = cur / 10 // 去掉最后一位
	}

	// 能否整除3
	if sum%3 == 0 {
		return true
	} else {
		return false
	}
}
```

