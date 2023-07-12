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

func f(x *Node) *Info {
    if x == nil {
        return nil
    }
    leftInfo := f(x.left)
    rightInfo := f(x.right)
    
    min := x.value
    max := x.value
    
    if leftInfo != nil {   // 左子树不为空，用左子树的信息更新max和min
        min = math.Min(min,leftInfo.min)
        max = math.Max(max,leftInfo.max)
    }
    
    if rightInfo != nil {   // 右子树不为空，用右子树的信息更新max和min
        min = math.Min(min,rightInfo.min)
        max = math.Max(max,rightInfo.max)
    }
    
    maxBSTSize := 0 
    maxBSTHead := &Node{}
    
    // 1.最大BST子树来自左子树的可能性
    if leftInfo != nil {   
       	maxBSTSize = leftInfo.maxBSTSize
        maxBSTHead = leftInfo.maxBSTHead
    }
    // 2.最大BST子树来自右子树的可能性
    if rightInfo != nil && rightInfo.maxBSTSize > maxBSTSize {   
       	maxBSTSize = rightInfo.maxBSTSize
        maxBSTHead = rightInfo.maxBSTHead
    }
    // 3.最大BST子树就是当前整个二叉树的可能性
    isBST := false
    // 1)保证左右子树都是BST
    if (leftInfo == nil || leftInfo.isBST) && (rightInfo == nil || rightInfo.isBST) {
        // 2) 左子树最大值 < 当前节点值 < 右子树的最小值
        if (leftInfo == nil || leftInfo.max < x.value) && (rightInfo == nil || rightInfo.min < x.value) {
            isBST = true
            // 整棵树都是BST，左右子树必定都是完整的BST,  maxBSTSize == nodeCount
            leftCount := leftInfo == nil ? 0:leftInfo.maxBSTSize   
            rightCount := rightInfo == nil ? 0:rightInfo.maxBSTSize
            
            maxBSTSize = leftCount + rightCount + 1
            maxBSTHead = x
        }
    }
    return &Info{maxBSTHead,isBST,min,max,maxBSTSize}
    
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

func maxSum(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    max := math.MinInt
    cur := 0
    for i:=0;i<len(arr);i++ {
        cur += arr[i]
        max = math.Max(max,cur)   // 历史最大累加和更新
        if cur < 0 {   // 一旦累加和小于0，下一个位置重新从0开始累加
            cur = 0
        } 
    }
    return max
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

