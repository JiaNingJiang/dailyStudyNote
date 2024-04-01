## 一、题目一

给定一个非负整数n，代表二叉树的节点个数。返回能返回多少种不同的二叉树结构。

```go
假设节点个数为N，根节点需要占用一个，因此可操作的节点实际只有N-1个。
1.如果 N = 0,那么能够返回0种不同的二叉树
2.如果 N = 1,那么能够返回1种二叉树
3.如果 N = 2,那么能够返回2种二叉树，自由节点在左侧或右侧时各一种
4.对于 N >= 3的情况，可以分为以下三种情况进行求解(以 N == 3 为例)：
1) 左子树0个节点，右子树2个节点
2) 左子树1个节点，右子树1个节点
3) 左子树2个节点，右子树0个节点
即左子树有i(i = 0……N-1)个节点,右子树有N-i-1个节点 分别进行递归求解

// 给你n个节点，返回可以组成的二叉树的个数
func BinaryTreeCount(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	count := 0
	for left := 0; left <= n-1; left++ { // 除了根节点外，实际可控制的节点个数为n-1个
		right := n - 1 - left
		leftCount := BinaryTreeCount(left)
		rightCount := BinaryTreeCount(right)

		if left == 0 {
			count += rightCount
		} else if right == 0 {
			count += leftCount
		} else if left != 0 && right != 0 {
			count += leftCount * rightCount
		}
	}
	return count
}

// 用动态规划实现
func BinaryTreeCount2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	if n <= 2 {
		return dp[n]
	}

	for node := 3; node <= n; node++ { // 节点的个数
		for left := 0; left <= node-1; left++ {
			right := node - 1 - left
			if left == 0 {
				dp[node] += dp[right]
			} else if right == 0 {
				dp[node] += dp[left]
			} else if left != 0 && right != 0 {
				dp[node] += dp[left] * dp[right]
			}
		}
	}

	return dp[n]
}
```



## 二、题目二

一个完整的括号字符串的定义如下：

1.空字符串是完整的

2.每一个 左括号"(" 都有与之相匹配的 右括号")"

例如： （（）（））（）是完整的括号字符串     （）)（  和 （）（ 是不完整的括号字符串

现在有一个括号字符串s，请问现在需要在其中任意位置尽量少地添加括号，将其转化为一个完整的括号字符串。请问至少需要添加多少个括号？

```go
1.如何判断一个括号字符串是否是完整的？
定义一个整型变量count。当遇到左括号的时候count++，遇到右括号的时候count——。在遍历过程中，一旦发现count<0，则说明缺一个左括号"(",说明不完整 ; 当遍历完所有字符，若count ！=0 ，说明缺少右括号")",不完整。

2.借助上述的思想完成 需要添加多少个括号 的判断。
准备两个变量：count和sum，count的意义不变。在遍历字符串过程中，一旦发现count < 0,那么让sum++(sum用来统计缺少的左括号个数)，然后让count重新归零，继续向下遍历。遍历完成后，count的值就是缺少的右括号的个数。count+sum就是总共缺少的括号个数。
```

```go
// 返回补全一个括号字符串到完整需要的最少括号数
func CompleteParentheses(str string) int {
	leftLack := 0
	rightLack := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			rightLack++
		} else if str[i] == ')' {
			rightLack--
		}
		if rightLack < 0 {
			leftLack++
			rightLack = 0
		}
	}
	return leftLack + rightLack
}
```



## 三、题目三

给定一个数组arr，求差值为k的数字对，要求不重复，即：如果k = 2，那么(2,4)和(4,2)只能算一个

```go
准备一个哈希表，先将arr数组中的重复元素去掉。
假设arr = [2 4 5 7 3 0 0 ], k == 2  哈希表的key值只会有{2 4 5 7 3 0}，多余的一个0被去掉了
接着遍历哈希表：（只看比当前数大的或者只看比当前数少的，这样可以保证正序和逆序只选一个）
1.遍历到2的时候查看4是否存在
2.遍历到4的时候查看6是否存在
3.遍历到5的时候查看7是否存在
4.遍历到7的时候查看9是否存在
5.遍历到3的时候查看5是否存在
6.遍历到0的时候查看2是否存在
```

```go
func NumberCouple(arr []int, dvalue int) {
	vSet := make(map[int]struct{})

	for _, v := range arr {
		if _, ok := vSet[v]; !ok {
			vSet[v] = struct{}{}
		}
	}

	for firV, _ := range vSet {
		if _, ok := vSet[firV+dvalue]; ok {
			fmt.Printf("数字对(%d - %d)\n", firV, firV+dvalue)
		}
	}
}
```



## 四、题目四

给一个包含n个整数元素的集合a，一个包含m个整数元素的集合b。

定义magic操作，从一个集合中取出一个元素，放到另一个集合里，且操作过后每个集合的平均值都大于操作前。

注意以下两点：

1.不可以把一个集合的元素清空，这样就没有平均值了

2.值为x的元素从集合b取出放入集合a，但集合a中已经有值为x的元素，则a的平均值不变（因为集合元素不会重复），b的平均值可能会改变（因为x被取出了）

问：最多可以进行多少次magic操作？

```go
1.如果arr1和arr2的平均值相等，假设都是100，此时是无法进行移动的，原因是：
	1）假设我们将arr1中小于100的数移动到arr2，那么arr1平均值上升，但是arr2平均值下降
	2) 假设我们将arr1中等于100的数移动到arr2，那么arr1和arr2的平均值都不会变
	3）假设我们将arr1中大于100的数移动到arr2，那么arr1平均值下降，但是arr2平均值上升

2.如果arr1平均值 < arr2 平均值，无法从arr1中移动一个数到arr2中实现magic操作，原因是：
	如果取出的数 <= arr1平均值，会导致arr2平均值下降 ； 如果取出的数 > arr1平均值，不管arr2平均值是否变大，arr1平均值都是减小。因此不可实现magic操作

3.如果arr1平均值 > arr2 平均值， 符合下面条件的值可以从arr1中移动一个数到arr2中实现magic操作
假设arr1平均值为avg1，arr2平均值为avg2.
	1）从arr1取出的数字必须大于avg2，小于avg1
	2）从arr1取出的数组不能已经存在于arr2  （用一个哈希表hash记录）

对于符合上述两点的arr1中的数，我们每次取一个符合条件且最小的数移动到arr2。原因是这样可以最大程度拉大avg1和avg2的差距，增大arr1中可以实现magic操作的数据范围。	

在实现时，还需要注意：平均值可能是小数，因此需要用两个浮点型存储avg1和avg2；arr1、arr2、avg1和avg2、hash每一轮都要重新生成。且arr1需要进行排序
```

```go
func MagicTime(arr1, arr2 []int) int {
	time := 0

	for {
		avg1 := average(arr1)
		avg2 := average(arr2)

		if avg1 == avg2 {
			return time
		}
		smallMap := make(map[int]struct{})

		if avg1 > avg2 { // 将arr1中 (avg2,avg1)范围内的数字移动到arr2中
			if len(arr1) <= 1 { // arr1不能为空
				return time
			}
			sort.Slice(arr1, func(i, j int) bool { // 对arr1进行排序
				if arr1[i] >= arr1[j] {
					return false
				} else {
					return true
				}
			})
			for _, v := range arr2 { // smallMap存储具有较小平均值的数组所有元素
				smallMap[v] = struct{}{}
			}
			res := 0
			arr1, arr2, res = magicOperation(arr1, arr2, avg1, avg2, smallMap)
			if res == 0 { // 找不到符合条件的数字可以移动
				return time
			} else {
				time += res
			}
		} else { // 将arr2中 (avg1,avg2)范围内的数字移动到arr1中
			if len(arr2) <= 1 { // arr2不能为空
				return time
			}
			sort.Slice(arr2, func(i, j int) bool { // 对arr2进行排序
				if arr2[i] >= arr2[j] {
					return false
				} else {
					return true
				}
			})

			for _, v := range arr1 { // smallMap存储具有较小平均值的数组所有元素
				smallMap[v] = struct{}{}
			}

			res := 0
			arr2, arr1, res = magicOperation(arr2, arr1, avg2, avg1, smallMap)
			if res == 0 { // 找不到符合条件的数字可以移动
				return time
			} else {
				time += res
			}
		}

	}

}

func MagicTimeRecursion(arr1, arr2 []int) int {
	avg1 := average(arr1)
	avg2 := average(arr2)
	smallMap := make(map[int]struct{})
	if avg1 == avg2 {
		return 0
	} else if avg1 > avg2 {
		if len(arr1) <= 1 {
			return 0
		}
		sort.Slice(arr1, func(i, j int) bool { // 对arr1进行排序
			if arr1[i] >= arr1[j] {
				return false
			} else {
				return true
			}
		})

		for _, v := range arr2 { // smallMap存储具有较小平均值的数组所有元素
			smallMap[v] = struct{}{}
		}
		if count := magicTimeRecursion(arr1, arr2, avg1, avg2, smallMap); count == -1 {
			return 0
		} else {
			return count
		}
	} else {
		if len(arr2) <= 1 {
			return 0
		}
		sort.Slice(arr2, func(i, j int) bool { // 对arr2进行排序
			if arr2[i] >= arr2[j] {
				return false
			} else {
				return true
			}
		})

		for _, v := range arr1 { // smallMap存储具有较小平均值的数组所有元素
			smallMap[v] = struct{}{}
		}
		if count := magicTimeRecursion(arr2, arr1, avg2, avg1, smallMap); count == -1 {
			return 0
		} else {
			return count
		}
	}
}

func magicTimeRecursion(bigArr, smallArr []int, bigAvg, smallAvg float64, smallMap map[int]struct{}) int {
	if bigAvg == smallAvg { // 需要结束
		return -1
	}
	if len(bigArr) == 1 { // 需要结束
		return -1
	}

	arr1, arr2, res := magicOperation(bigArr, smallArr, bigAvg, smallAvg, smallMap)
	if res == 0 { // 无法转移，需要结束
		return -1
	}
	avg1 := average(arr1)
	avg2 := average(arr2)
	newSmallMap := make(map[int]struct{})
	if avg1 == avg2 { // 完成了一次，但是不能继续了
		return 1
	} else if avg1 > avg2 { // 可能还能继续，下一次从arr1中移动数字到arr2
		sort.Slice(arr1, func(i, j int) bool { // 对arr1进行排序
			if arr1[i] >= arr1[j] {
				return false
			} else {
				return true
			}
		})
		for _, v := range arr2 { // smallMap存储具有较小平均值的数组所有元素
			newSmallMap[v] = struct{}{}
		}
		if res := magicTimeRecursion(arr1, arr2, avg1, avg2, newSmallMap); res == -1 {
			return 1
		} else {
			return 1 + res
		}

	} else { // 可能还能继续，下一次从arr2中移动数字到arr1
		sort.Slice(arr2, func(i, j int) bool { // 对arr2进行排序
			if arr2[i] >= arr2[j] {
				return false
			} else {
				return true
			}
		})
		for _, v := range arr1 { // smallMap存储具有较小平均值的数组所有元素
			newSmallMap[v] = struct{}{}
		}
		if res := magicTimeRecursion(arr2, arr1, avg2, avg1, newSmallMap); res == -1 {
			return 1
		} else {
			return 1 + res
		}
	}
}

func magicOperation(bigArr, smallArr []int, bigAvg, smallAvg float64, smallMap map[int]struct{}) ([]int, []int, int) {

	for i, v := range bigArr { // 从bigArr选出第一个符合条件的数(也是符合数中最小的一个)移动到smallArr
		if _, ok := smallMap[v]; ok { // 如果bigArr中的这个数在smallArr中已经存在，则跳过
			continue
		}
		if float64(v) > smallAvg && float64(v) <= bigAvg {
			smallArr = append(smallArr, v)
			bigArr = append(bigArr[:i], bigArr[i+1:]...)

			smallMap[v] = struct{}{}
			return bigArr, smallArr, 1
		}
	}
	return bigArr, smallArr, 0
}

func average(arr []int) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += float64(v)
	}
	avg := sum / float64(len(arr))
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", avg), 64)
	return value
}
```

