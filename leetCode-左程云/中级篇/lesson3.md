## 一、题目一

现在有一个括号字符串，要求你求出其中最长的完整括号子串。

例如：()()()的返回值是6；())()())返回值是4；)和(的返回值都是0

```go
1.准备一个和总括号字符串长度一样的辅助数组dp
2.从左到右遍历括号字符串，dp[i]记录截止到当前位时的最长括号子串长度
1）如果当前第i位是"("，那么dp[i] = 0，直接跳过
2）如果当前第i位是")"，那么需要按照以下步骤进行dp[i]的计算
① 首先要获取其前一位的dp值即dp[i-1]的值，从当前在总括号字符串的位置向前跳转 dp[i-1]+1个位置。然后判断当前位置是否是"("且有没有越界，如果没有越界且正好是"(",那么dp[i]至少应该是 dp[i-1]+2。至于是不是"至少"，还需要进行第②步
② 完成第一步的跳跃且检查正好是"("，那么接着检查当前"("位置(下标为j)的前一个位置的dp[j-1]是否符合“未越界且大于0”(是一个完整的括号子串)，那么dp[i] 需要在 dp[i-1]+2 的基础上再加上 dp[j-1]

3.完成遍历后，我们返回dp数组中最大的值
```

```go
func LongestPstr(pstr string) int {
	if len(pstr) == 0 {
		return 0
	}
	dp := make([]int, len(pstr))

	for i := 0; i < len(pstr); i++ {
		if pstr[i] == '(' { // pstr[i] == '('
			dp[i] = 0
		} else { // pstr[i] == ')'
			last := i - 1
			if last < 0 { // dp[i-1]是越界的
				continue
			}
			skipInterval := dp[i-1] + 1
			if i-skipInterval < 0 || pstr[i-skipInterval] != '(' { // 跳跃点不存在或者不是'('
				dp[i] = 0
				continue
			}
			dp[i] = dp[i-1] + 2 // dp[i] 至少是 dp[i-1] + 2

			aSkipInterval := skipInterval + 1 // 再往前多跳一个位置
			if i-aSkipInterval >= 0 && dp[i-aSkipInterval] > 0 {
				dp[i] += dp[i-aSkipInterval]
			}
		}
	}
	maxLen := math.MinInt
	for i := 0; i < len(dp); i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
```



## 二、问题二

对一个栈里的整型数据，按照升序进行排序（即排序前，栈里的数据是无序的，排序后最大元素位于栈底），要求最多只能使用一个额外的栈存放临时数据，但不能将元素复制到别的数据结构中。

```go
准备一个辅助的栈，要求辅助栈中的元素必须按照降序(栈顶最大，栈底最小)存储
1.从栈中弹出一个元素到辅助栈
2.再从栈中弹出一个元素，需要查看是否能够放入辅助栈(能否满足辅助栈的降序规则): 如果能放直接放; 如果不能放，则需要将辅助栈中的元素弹到原栈中，弹出到可以在辅助栈放入这个元素。
3.持续进行步骤2，直到栈中的元素全部转移到辅助栈。
4.最后将辅助栈的元素全部弹出到栈中。
```

```go
// increase == true表示变成递增站；increase == false表示变成递减栈
func Orderization(stack *utils.Stack, increase bool) {
	aStack := utils.NewStack() // 辅助栈

	// 将原始栈中的元素全部移动到辅助栈，同时保证辅助栈与原始栈的目标单调性相反
	for {
		if stack.Len == 0 {
			break
		}
		ele := stack.Pop().(int) // 从原始栈中弹出一个元素

		if aStack.Len == 0 { // 仅当辅助栈为空时可以直接加入到辅助栈
			aStack.Push(ele)
			continue
		}
		top := aStack.Top().(int)
		if compare(ele, top, !increase) { // 需要确保辅助栈的栈顶元素与原始栈中弹出的元素符合相反的单调关系
			aStack.Push(ele)
		} else { // 如果不满足单调关系，则需要从辅助栈中持续弹出元素，直到符合单调关系或者辅助栈为空
			for {
				if aStack.Len == 0 {
					aStack.Push(ele)
					break
				}
				top := aStack.Top().(int)
				if compare(ele, top, !increase) {
					aStack.Push(ele)
					break
				}
				data := aStack.Pop()
				stack.Push(data)
			}
		}
	}

	for {
		if aStack.Len == 0 {
			return
		}
		data := aStack.Pop()
		stack.Push(data)
	}

}

func compare(original, assistTop int, increase bool) bool {
	if increase {
		if original > assistTop {
			return true
		} else {
			return false
		}
	} else {
		if original < assistTop {
			return true
		} else {
			return false
		}
	}
}
```



## 三、问题三

现在有一个矩阵(M*N)，它符合以下条件：

1.每一行元素都是从小到大的

2.每一列元素都是从小到大的

要求从从矩阵中查询一个元素是否存在，要求时间复杂度为 O(M+N)

假设矩阵如下，要求查询元素7是否存在？

|  1   |  5   |  9   |  10  |
| :--: | :--: | :--: | :--: |
|  2   |  6   |  11  |  13  |
|  7   |  9   |  15  |  17  |

```go
1.总是以矩阵右上角为起点，这里就是元素10
2.因为目标7小于10，因此10所在列以下的直接跳过。转到左侧，也就是9
3.因为目标7小于9，因此9所在列以下的直接跳过。转到左侧，也就是5
4.因为目标7大于5，因此往下走，走到6
5.因为目标7大于6，因此往下走，走到9
6.因为目标7小于9，因此9所在列以下的直接跳过。转到左侧，也就是7

采用这种方式，要查找一个元素，最多需要经过 M+N 步。
一旦在访问时发生了矩阵越界，那么意味着目标元素不在矩阵中。
```

```go
// 矩阵的每一行都是从小到大；每一列也是从小到大
func SearchInMatrix(matrix [][]int, target int) bool {
	maxRow := len(matrix) - 1    // 矩阵的最大行号
	maxCol := len(matrix[0]) - 1 // 矩阵的最大列号

	// 起点从矩阵的右上角开始(第0行，最后一列), 因此每次移动只能往下或者往左走
	curRow := 0
	curCol := maxCol
	for {
		if curRow < 0 || curRow > maxRow || curCol < 0 || curCol > maxCol {
			return false
		}
		if matrix[curRow][curCol] == target {
			return true
		}
		if matrix[curRow][curCol] > target { // 当前矩阵元素大于目标值，则往左走
			curCol--
			continue
		}
		if matrix[curRow][curCol] < target { // 当前矩阵元素小于目标值，则往下走
			curRow++
			continue
		}
	}
}
```



## 四、问题四

现在有一个矩阵(M*N)，它符合以下条件：

1.元素只能是0或者1

2.对于每一行，可以是全0行，也可以是全1行。但如果是0/1混合行，必须满足以下条件：

1）0必须是连续的且在该行的左侧

2）1必须是连续的且在该行的右侧

要求求出存在1最多的那一行的1的个数。

假设矩阵如下：

| 0    | 0    | 0    | 1    | 1    | 1    |
| ---- | ---- | ---- | ---- | ---- | ---- |
| 0    | 0    | 0    | 0    | 0    | 1    |
| 0    | 0    | 0    | 0    | 1    | 1    |
| 0    | 0    | 1    | 1    | 1    | 1    |
| 0    | 0    | 1    | 1    | 1    | 1    |
| 0    | 1    | 1    | 1    | 1    | 1    |

```go
1.总是以矩阵右上方为起点，准备一个额外的变量max存储最多的1的个数
2.如果当前位置是1，那么往左走，直到走到该行边界的1，更新max，然后往下走；如果当前位置是0，则继续往下走。
3.重复步骤2，直到遍历完所有行。

对于此例子，当遍历完：
第一行，max=3
第二行，max=3
第三行，max=3
第四行，max=4
第五行，max=4
第六行，max=5
```

```go
package lesson3

func MostOne(matrix [][]int) int {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	// 从矩阵的右上角开始
	curRow := 0
	curCol := maxCol

	mostOne := 0
	for {
		if curRow < 0 || curRow > maxRow || curCol < 0 || curCol > maxCol {
			return mostOne
		}
		if matrix[curRow][curCol] == 1 { // 当前位置是1，则累加1的个数，并向左走
			mostOne++
			curCol--
		} else { // 当前位置是0,直接向下走
			curRow++
		}
	}
}
```

