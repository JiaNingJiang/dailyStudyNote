## 一、强关系表达式问题的求解

### 1. 斐波那契数列

我们知道斐波那契数列的求解关系表达式是固定的（也就是强关系表达式）：

F(1) = F(2) = 1 作为初始条件    F(N) = F(N-1) + F(N-2) (N>=2)

```
根据线性代数，对于强关系表达式可以化为以下求解方式：
						  |a  b|
|F(3) F(2)| = |F(1) F(0)|*|c  d|    // 因为 N 需要借助 N-1和N-2两个状态量，因此系数矩阵是 2*2

						  |a  b|
|F(4) F(3)| = |F(2) F(1)|*|c  d|

四个系数可以通过4个表达式求解：
a = b = c = 1 , d = 0

因此任意的：（系数矩阵直接用Factor表示）
|F(N) F(N-1)| = |F(2) F(1)| * Factor^(N-2)   //因此求解斐波那契数列问题，就变成了求系数矩阵的N-2次幂的问题
```

```go
func selfMulMatrix(self interface{}, factor interface{}) interface{} {

	if self == nil { // self总是一个2x2的矩阵（开始是一个单位矩阵(对角线为1，其余位置为0)）
		selfMatrix := make([][]int, 2)
		selfMatrix[0] = make([]int, 2)
		selfMatrix[1] = make([]int, 2)
		selfMatrix[0][0] = 1
		selfMatrix[0][1] = 0
		selfMatrix[1][0] = 0
		selfMatrix[1][1] = 1
		factorMatrix := factor.([][]int)
		return utils.MatrixMultiply(selfMatrix, factorMatrix)
	} else {
		selfMatrix := self.([][]int)
		factorMatrix := factor.([][]int)
		return utils.MatrixMultiply(selfMatrix, factorMatrix)
	}
}

func Fibonacci(n int64) int {
	if n <= 2 {
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		} else if n == 2 {
			return 1
		} else {
			panic("错误的输入")
		}
	}

	factorMatrix := [][]int{{1, 1}, {1, 0}} // 系数矩阵

	factorMatrixPow := EffectivePow(factorMatrix, n-2, selfMulMatrix) // 求系数矩阵的 n-2次幂

	initial := [][]int{{1, 1}}
	res := utils.MatrixMultiply(initial, factorMatrixPow.([][]int))

	return res[0][0]
}

func MatrixMultiply(matrixA [][]int, matrixB [][]int) [][]int {
	rowsA := len(matrixA)
	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	// 检查矩阵尺寸是否兼容
	if colsA != rowsB {
		return nil
	}

	// 创建结果矩阵
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// 矩阵相乘
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}
```



```go
针对求解一个数的N次幂，可以用如下的方式：
假设要求10^73 == ? 可以将73拆分为2的幂的组成形式：73 = 64 + 8 + 1 
因此 10^73 = 10^64 * 10^8 * 10^1 
1）我们让 t 从10^1开始，t = 10^1 。因为存在于10^73组成公式中，因此sum+=10^1。
2）t = t * t ,因此t = 10^2 不在组成公式中(用73的二进制检查是否在组成公式中)，跳过
3) 重复，直到最大的 10^64

通过上述的方式，求解一个数的N次幂，时间复杂度就只是 O(logN) 级别，远远优于O(N)

func EffectivePow(x interface{}, y int64, selfMul func(interface{}, interface{}) interface{}) interface{} {
	// 1.将y拆解为二进制形式
	bitFormat := make([]byte, 0)
	for i := 0; 1<<i <= y; i++ {
		bitRes := ((1 << i) & y) >> i
		if bitRes == 1 {
			bitFormat = append(bitFormat, 1)
		} else {
			bitFormat = append(bitFormat, 0)
		}
	}
	var powRes interface{}

	for i := 0; i < len(bitFormat); i++ {
		if bitFormat[i] == 1 {
			powIndex := 1 << i
			for j := 0; j < powIndex; j++ {
				powRes = selfMul(powRes, x)
			}
		}
	}

	return powRes
}

```

## 2.类推

假设有一问题的强关系表达式为 :  F(N) = 3 * F(N-1) + 2 * F(N-3) + 5 * F(N-5)

那么求解的策略如下：

```
|F(6) F(5) F(4) F(3) F(2)| = |F(5) F(4) F(3) F(2) F(1)| * |5*5矩阵|
根据这样的列举，我们可以求出这个 5*5 矩阵

|F(N) F(N-1) F(N-2) F(N-3) F(N-4)| = |F(5) F(4) F(3) F(2) F(1)| * |5*5矩阵|^(N-5)
只需要获得 |5*5矩阵|^(N-5) 就可以得到任意的 F(N)
```



## 3.奶牛问题

假设有一个牧场，现在只有一头3岁的母牛。

通过人工授精的方式每年让所有年龄 >= 3的母牛生出一只新母牛。假设所有的牛都不会死。

求N年后牧场有多少只母牛？

```go
起始时，F(0) = 1  A
第一年：F(1) = 1+1 = 2  A B(A)
第二年：F(2) = 2+1 = 3  A B(A) C(A)
第三年：F(3) = 3+1 = 4  A B(A) C(A) D(A)
第四年：F(4) = 4+2 = 6  A B(A) C(A) D(A) E(A) F(B)   // B刚好三岁，可以生出新的母牛
……………………

计算公式 F(N) = F(N-1) + F(N-3)  // F(N-1)表示前一年剩余的牛都存活的数量  F(N-3)表示新生牛数量 == 三年前的牛的数量

|F(3) F(2) F(1)| = |F(2) F(1) F(0)| * |3*3矩阵|
求出这个 |3*3矩阵| 矩阵，假设为Factor

|F(N) F(N-1) F(N-2)| = |F(2) F(1) F(0)| * |3*3矩阵|^(N-2)
```



进阶：假设一只牛只能活10年，求N年后牧场有多少只奶牛？

```go
F(N) = F(N-1) + F(N-3) - F(N-10)   // 减去10年前的奶牛数即可
```

## 二、题目二

字符串只有  '0'  和  '1'  两种字符构成，

当字符串长度为1时，所有可能的字符串位"0"、“1”

当字符串长度为2时，所有可能的字符串位"00"、“01”、“10”、“11”

当字符串长度为3时，所有可能的字符串位"000"、“001”、“010”、“011”、“100”、“101”、“110”、“111”

………………

如果某一字符串中，只要出现'0'的位置，左边就靠着'1'，这样的字符串叫做达标字符串。

给定一个正数N，返回所有长度为N的字符串中，达标字符串的数量。

比如，N=3，返回3，因为只有"101"、“110”、“111”达标。

```go
假设字符串总长度为i，求解达标字符串数量的函数为 F(i) 
要想成为达标字符串，那么第1位必然只能是1。

对于第二位，有如下规则：
1.可以是1，如果第二位是1，那么剩余数量字符可以构成的达标字符串的数量完全等于 F(i-1)
2.可以是0，如果第二位是0，这是F(i-1)不曾有的场景(因为F(i-1)中i-1长度字符串首位必须是1，这里反而是让i-1长度字符串的首位是0)。
那么下一位固定地必须是1，此时剩余字符可以构成的达标字符串的数量完全等于 F(i-2)

所以，F(i) = F(i-1) + F(i-2)  // F(i-1)表示继承i-1长度时达标字符串的个数,因为在左边加一个1不会使已有达标字符串失效; F(i-2)则表示因为左侧多加了一个1，导致原本i-1长度字符串首位可以是0，因此前两位固定为10的情况下，剩余长度字符串达标字符串个数则是继承F(i-2) 

很明显，这是一个斐波那契数列，只是初始条件 F(1) = 1 , F(2) = 2

此问题的时间复杂度就是 O(logN) ,与暴力尝试的 O(2^N*N)相比，明显优化非常大
```



## 三、题目三

在草原上，小红捡到了n根木棍，第i根木棍的长度为i，她从中选出其中的三根木棍组成三角形。

现在小明想去掉一些木棍，使得小红任意选三根木棍都不能组成三角形。

请问小明最少去掉多少根木棍？

```go
现在有 1,2,3,4,……,n 这n根不同长度的木棍。
计算位于 1~n 这个范围内的斐波那契数列，出现的斐波那契数列以外的木棍都去掉，即为去掉最少得木棍

木棍：1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 
斐波那契数列：1 1 2 3 5 8 13 21

保留的木棍：1 2 3 5 8 13 
去掉的木棍：4 6 7 9 10 11 12 14 15 16 17
```

```go
package lesson6

func BrokenTriangle(n int) int {
	fibSet := make(map[int]struct{}, 0)

	var fibCount int64 = 1
	for {
		curFib := Fibonacci(fibCount)
		if curFib > n {
			break
		}

		fibCount++
		fibSet[curFib] = struct{}{}
	}
	remove := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		if _, ok := fibSet[i]; !ok {
			remove = append(remove, i)
		}
	}

	return len(remove)
}

```



## 四、题目四

牛牛准备参加学校组织的春游，出发前牛牛准备往背包里装入一些零食，牛牛的背包容量为w

牛牛家里一共有n袋零食，第i袋零食体积为 v[i]

牛牛想知道在总体积**不超过背包容量**的情况下，他一共有多少种零食放法（总体积为0也算一种放法）。

``` go
可以用递归法和动态规划解决：
递归法: 遍历v[]零食体积数组，每一轮遍历都可以分为 要v[i]和不要v[i] 两种情况

动态规划：（二维表根据最后一行向上求解，最后返回dp[0][0]作为返回值）
1.初始条件 dp[n-1][weight] = 1 (weight <= w )  dp[n-1][weight] = 1 (weight > w )
2. dp[i][j] = dp[i+1][j] + dp[i+1][j-v[i]] 

```

```go
package lesson6

// snacks: 记录各个零食的重量
// bagSize: 背包的大小
// 返回值: 可以的零食放法
func BagSnack(snacks []int, bagSize int) int {
	return bagSnack(snacks, bagSize, 0)
}

func bagSnack(snacks []int, remainBag, index int) int {
	if index >= len(snacks) {
		return -1
	}
	if remainBag == 0 {
		return 1
	}
	if remainBag < 0 {
		return -1
	}

	yao := bagSnack(snacks, remainBag-snacks[index], index+1)
	buyao := bagSnack(snacks, remainBag, index+1)

	if yao == -1 && buyao == -1 {
		return -1
	} else if yao == -1 {
		return buyao
	} else if buyao == -1 {
		return yao
	} else {
		return yao + buyao
	}
}

// snacks: 记录各个零食的重量
// bagSize: 背包的大小
// 返回值: 可以的零食放法
func BagSnackDP(snacks []int, bagSize int) int {
	return bagSnackDP(snacks, bagSize)
}

func bagSnackDP(snacks []int, bagSize int) int {

	// 行表示背包内零食重量，从0~bagSize，共bagSize+1行
	// 列表示当前遍历到零食下标，0~len(snacks)-1, 共len(snacks)列
	// 初始已知：最后一行全1，因为零食重量 == bagSize，且零食数组没有越界
	// 目标求解值: (0,0)
	// 依赖关系： matrix[size][i] = matrix[size+snacks[i]][i+1] + matrix[size][i+1]

	dp := make([][]int, bagSize+1)
	for i := 0; i <= bagSize; i++ {
		dp[i] = make([]int, len(snacks))
	}

	// 初始条件
	for col := 0; col < len(snacks); col++ {
		dp[bagSize][col] = 1
	}

	for row := bagSize - 1; row >= 0; row-- { // 从下往上(因为dp[i][j]依赖于下方)
		for col := len(snacks) - 1; col >= 0; col-- { // 从右向左(因为dp[i][j]依赖于右方)
			lowerRow := row + snacks[col]
			rightCol := col + 1
			if lowerRow <= bagSize && rightCol < len(snacks) { // size+snacks[i]和i+1都没有越界
				dp[row][col] = dp[lowerRow][col+1] + dp[row][col+1]
			} else if rightCol >= len(snacks) { //i+1都越界
				dp[row][col] = 0
			} else if lowerRow > bagSize && rightCol < len(snacks) { // 只有size+snacks[i]越界
				dp[row][col] = dp[row][col+1]
			}
		}
	}
	return dp[0][0]
}
```



## 五、题目五

为了找到满意的工作，牛牛收集到了每种工作的难度和报酬。牛牛选工作的标准是在难度不超过自身能力值的情况下，选择报酬最高的工作。在牛牛选定了自己的工作后，牛牛的小伙伴来找牛牛帮忙选工作。

给定一个Job结构体（包括工作难度hard和工作报酬money）数组jobarr，再给定一个int类型的数组arr，表示所有小伙伴的能力。

返回一个int型数组，表示每一个小伙伴按照牛牛的标准选工作后能获得的报酬。

```go
1.先对jobarr按照工作难度hard从小到大排序。
2.工作难度hard相等的归为一组，组内按照报酬money从大到小排序，只保留每一种money最多的job
3.剩余的jobarr数组元素，再从小到大遍历，因为工作难度hard必然是递增的，所以一旦遇到报酬money小于前一项的，直接将该job移出。
4.最后jobarr中的元素：按照hard从小到大，按照money从小到大。每个人只需要选择最接近自己能力的hard的job即可。


```

```go
package lesson6

import (
	"sort"
)

type Job struct {
	Ability int // 工作需要的能力
	Salary  int // 工作提供的薪资
}

// 工作数组jobArr，个人能力数组capArr
func AcAbilityGetWork(jobArr []Job, capArr []int) []int {
	// 1.将工作数组按照工作难度从小到大排序
	sort.Slice(jobArr, func(i, j int) bool {
		if jobArr[i].Ability < jobArr[j].Ability {
			return true
		} else {
			return false
		}
	})

	filterArr := make([]Job, 0) // 只记录各工作难度下，薪资最高的工作
	maxSalary := jobArr[0].Salary
	curCap := jobArr[0].Ability
	// 2.对于工作难度相等的工作，jobArr中只保留薪资最高的那一个
	for i := 1; i < len(jobArr); i++ {
		if jobArr[i].Ability == curCap { // 工作难度相等
			if jobArr[i].Salary > maxSalary {
				maxSalary = jobArr[i].Salary // 在相同的工作难度下，记录薪资最高的一个
			} else {
				if i == len(jobArr)-1 {
					filterArr = append(filterArr, Job{Salary: maxSalary, Ability: curCap})
				}
			}
		} else { // 只能是工作难度增大了，而不可能是变小了
			filterArr = append(filterArr, Job{Salary: maxSalary, Ability: curCap}) // 记录上一个工作难度下的，薪资最高的工作
			maxSalary = jobArr[i].Salary
			curCap = jobArr[i].Ability
		}
	}

	increaseJob := make([]Job, 0) // 工作如果难度上升，那么薪资也需要提高

	increaseJob = append(increaseJob, filterArr[0])
	lastJob := 0
	for i := 1; i < len(filterArr); i++ {
		if filterArr[i].Salary >= increaseJob[lastJob].Salary {
			increaseJob = append(increaseJob, filterArr[i])
			lastJob++
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(capArr); i++ {
		personCap := capArr[i]
		salary := 0
		for _, job := range increaseJob {
			if personCap >= job.Ability {
				salary = job.Salary
			}
		}
		res = append(res, salary)
	}

	return res
}
```



## 六、题目六

现在有一个数字字符串，需要将其转化为int型数值，请问如何操作？

```go
1.无论字符串表示的是正整数还是负整数，统一转换为负整数（因为负整数个数比正整数多一个）
2.假设字符串为“-1023”，转换过程为:
0 + (-1) = -1  (-1)*10 +(-0) = -10  -10*10+(-2) = -102   -102*10+(-3) = -1023

package lesson6

import "math"

func Convert(numStr string) int {
	var neg bool
	if numStr[0] == '-' {
		neg = true
		numStr = numStr[1:]
	} else {
		neg = false
	}

	var res int // 最后转换得到的整数值（总是负数值）
	// 为了防止转化后数值溢出，准备了下面两个变量
	spillCheck := math.MinInt / 10
	spillCheckRe := math.MinInt % 10

	for i := 0; i < len(numStr); i++ {
		eleInt := -int(numStr[i] - '0')
		if res < spillCheck || res == spillCheck && eleInt < spillCheckRe {
			panic("数值转换溢出")
		}

		res = res*10 + eleInt
	}

	if !neg && res == math.MinInt {
		panic("数值转换溢出")
	}
	if neg {
		return res
	} else {
		return -res
	}
}
```

