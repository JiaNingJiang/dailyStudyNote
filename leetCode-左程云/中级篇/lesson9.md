## 一、题目一

给定一个整型数组A，长度为n，有 1 <= A[i] <= n , 且对于[1,n] 的整数，其中部分整数会重复出现，部分整数不会出现。

实现算法找到[1,n] 中所有未出现在A中的整数。

提示：尝试实现 O(n) 的时间复杂度和 O(1) 的空间复杂度

```go
思路：争取让每一个 i 位置都存储 i+1 , 完成此任务后再遍历数组，就知道缺的数字是什么了。
比如：
    3 2 1 6 2 7 5
    0 1 2 3 4 5 6
1.从arr[0]开始，3应该放在arr[2]上，arr[2] = 1 就拿出来
      2 3 6 2 7 5
    0 1 2 3 4 5 6
2. 1应该放在arr[0]上，arr[0]现在没数字，正好放入：
    1 2 3 6 2 7 5
    0 1 2 3 4 5 6
3.一轮交换以回到起始位置为结束，开始下一轮交换。
4.从arr[1]开始，2就应该放在arr[1]上，因为回到了新起始位置，开始下一轮交换。
5.从arr[2]开始，3就应该放在arr[2]上，因为回到了新起始位置，开始下一轮交换。
6.从arr[3]开始，6应该放到arr[5]上, arr[5] = 7 就拿出来
    1 2 3   2 6 5
    0 1 2 3 4 5 6
7. 7就应该放到arr[6]上，arr[6] = 5 就拿出来
    1 2 3   2 6 7
    0 1 2 3 4 5 6
8. 5就应该放到arr[4]上，arr[4] = 2 就拿出来
    1 2 3   5 6 7
    0 1 2 3 4 5 6
9. 2就应该放到arr[1]上，但是arr[2]已经是1了，因此不需要交换。一旦遇到不需要交换的场景，意味着本轮结束，开始下一轮交换。
10.从arr[4]开始，5就应该放在arr[4]上，因为回到了新起始位置，开始下一轮交换。
11.从arr[5]开始，6就应该放在arr[5]上，因为回到了新起始位置，开始下一轮交换。
12.从arr[6]开始，7就应该放在arr[6]上，因为回到了新起始位置，开始下一轮交换。
13.结束了，所有交换完成。
14.再次遍历arr数组，发现arr[3] != 4 , 说明整个数组缺少4
```

```go
package lesson9

import "math"

func LookUpMissing(arr []int) []int {
	n := len(arr) // 在arr数组中出现数字的最大可能值（最小可能值是1）

	for i := 0; i < n; i++ { // 每轮交易都以从当前第i为开始，以重新回到第i位结束

		if arr[i] == i+1 { // 直接不需要交换，结束本轮
			continue
		}

		curVal := arr[i]
		curIndex := i

		// 需要交换，意味着需要从当前位置拿走一个数，导致该位置的空缺(只有每轮的第一次取是只取不填，因此导致空缺)
		temp := arr[curVal-1]
		arr[curVal-1] = curVal
		curVal = temp // 通过交换而获得的值

		arr[curIndex] = math.MinInt // 用来表示空缺
		curIndex = curVal - 1       // 下一个需要检查的位置

		for {

			if arr[curIndex] == curVal { // 如果arr[curVal-1] == curVal,说明不需要进行交换了，那么本轮结束
				break
			}
			// 此时 arr[curVal-1] != curVal ,那么令 arr[curVal-1] = curVal，并且把 arr[curVal-1] 原本的值拿到
			temp1 := arr[curVal-1]
			arr[curVal-1] = curVal
			curVal = temp1

			if curVal == math.MinInt { // 说明原本的位置是没有值的，刚好把一个值填了进去
				break // 因为本来是空的，因此通过交换无法获取新值，因此结束本轮交换
			}

			curIndex = curVal - 1 // 下一个需要被交换的数组元素的下标是 curVal - 1

		}
	}

	absenceSet := make([]int, 0)
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			absenceSet = append(absenceSet, i+1)
		}
	}
	return absenceSet
}
```

## 二、题目二

最近CC直播平台在举办主播唱歌比赛，假设某一女主播的初始人气值为start，能够晋升到下一轮人气需要刚好达到end，给主播增加人气的方法有：

1. 点赞，花费 x C币，人气+2
2. 送礼，花费 y C币，人气*2
3. 私聊，花费 z C币，人气-2

其中end远远大于start，且start和end都是偶数，请问最少花费多少C币能帮助该主播人气达到end从而完成晋级？

```go
经典的暴力递归问题，关键在于边界条件的设置：

一眼就能看出的边界条件： 假设当前人气值用cur表示，那么cur == end 即为边界条件。
但是仅靠这一个边界条件是不够的。

1.寻找频繁解
因为start和end都是偶数，必定可以只通过点赞就完成晋级，如果只通过点赞的方式，需要花费的C币为： count:= (end - start)/2 * x。如果发现花费的C币总数大于count，那么就必然不是最优解。

2.从业务本身考虑获取边界条件
1）人气数不能为负数，因为start、end以及三种人气变化方式的原因，当前的人气值必然只能是偶数，一旦为负数，那必然只能是先为-2，-2人气变回整数还需要额外的一次点赞。因此人气一旦为负数，那必然不是最优解
2）人气值必然不可能超过 2*end，因为在最优解中使用私聊的方式只会发生在 end < cur < 2*end的情况下。

package lesson9

import "math"

// start,end 分别表示初始人气值和目标人气值
// add 表示人气值+2 需要消耗的金币
// twice 表示人气值*2 需要消耗的金币
// del 表示人气值-2 需要消耗的金币
func PopuValue(start, end int, add, twice, del int) int {
	return popuValue(start, end, add, twice, del, 0, 0)
}

func popuValue(start, end int, add, twice, del int, curPopu int, curCoin int) int {
	if curPopu == end { // 达到目标人气值
		return curCoin
	}
	if curPopu >= 2*end { // 最优情况下，人气值不可能超过end的2倍
		return -1
	}
	if curPopu < 0 { // 最优情况下，人气值不可能小于0
		return -1
	}

	if curCoin > (end-start)/2*add { // 消耗的金币数比只凭点赞消耗的更多，就不可能是最优情况
		return -1
	}

	addRes := popuValue(start, end, add, twice, del, curPopu+2, curCoin+add)
	twiceRes := popuValue(start, end, add, twice, del, curPopu*2, curCoin+twice)
	delRes := popuValue(start, end, add, twice, del, curPopu-2, curCoin+del)

	// 只有当三种分区的结果都为-1的时候，返回值才会是-1
	if addRes == -1 && twiceRes == -1 && delRes == -1 {
		return -1
	}

	// 只要存在有分支！=-1的情况下，返回有效分支中消耗金币数最少的
	addCoin := math.MaxInt
	twiceCoin := math.MaxInt
	delCoin := math.MaxInt

	if addRes != -1 {
		addCoin = addRes
	}
	if twiceRes != -1 {
		twiceCoin = twiceRes
	}
	if delRes != -1 {
		delCoin = delRes
	}

	res := getMin(getMin(addCoin, twiceCoin), delCoin)

	return res
}
```

## 三、题目三

CC直播的运营部门组织了很多运营活动，每个活动需要花费一定的时间参与，主播每参加完一个活动就可以得到一定的奖励，**参与活动可以从任意活动开始，但是一旦参与活动就必须将后续的活动全部参加完毕**。活动之间存在一定的依赖关系(但不存在环的情况)，现在给出所有的活动时间和依赖关系，请计算出在任意天数，能够获得的最大奖励都分别是多少？

假设活动依赖关系图为：

<img src="lesson9.assets/image-20230624145801351.png" alt="image-20230624145801351" style="zoom: 50%;" />

```go
需要从最后结尾的活动开始向前依次计算：
图的每一个节点都是一个活动，活动有它自己的消耗天数以及报酬。除此之外，每一个活动节点还包括一个有序表（要求活动天数递增的，报酬也必须是递增的），有序表负责记录从当前活动节点到结尾节点的天数+报酬信息。
1.从活动E开始，有序表的只有一条记录： {5,1000}
2.活动B,计算到结尾的活动信息： {8,1020}, 有序表也只有这一条记录 {8,1020}
3.活动C,计算到结尾的活动信息： {6,1200}, 有序表也只有这一条记录 {6,1200}
4.活动D,计算到结尾的活动信息： {9,1500}, 有序表也只有这一条记录 {9,1500}
5.活动A,计算到结尾的活动信息有三条，分别是：{12,1040} {10,1220} {13,1520}，因为活动路线1{12,1040}明显劣于(消耗天数多，但报酬却少)路线2{10,1220}，所以活动节点A的有序表只保存两条路线 [ {10,1220}，{13,1520} ]

6.最后从头到尾遍历整个图，获取每个节点的有序表，组成一个完整的有序表，合并后要排除掉所有活动天数递增但报酬没有递增的路线
合并结果： [ {10,1220}，{13,1520},{8,1020},{6,1200},{9,1500},{5,1000}  ]
有序化： [ {5,1000}，{6,1200}，{8,1020}，{9,1500}，{10,1220}，{13,1520}  ]
过滤后： [ {5,1000}，{6,1200}，        ，{9,1500}，         ，{13,1520}  ]

最终就得到了各工作天数下能获得的最大报酬都是多少。
```

```go
package lesson9

import "sort"

// 每一条活动路线起点任意，但是终点必须是最后一个活动
type ActRoute struct {
	TotalConsume int // 活动路线总计消耗的天数
	TotalSalary  int // 活动路线总计获得的报酬
}

type Activity struct {
	Consume  int         // 参加活动消耗的天数
	Salary   int         // 参加活动获得的报酬
	NextSet  []*Activity // 保存当前活动的下一个可以的活动
	LastSet  []*Activity // 用于追溯当前活动的上一个可以的活动
	OrderSet []*ActRoute // 活动路线的有序表(按照消耗天数从小到大有序，且天数与报酬需要成正比)
}

// 构建一个新活动
func NewActivity(consume, salary int, next, last []*Activity) *Activity {
	return &Activity{
		Consume:  consume,
		Salary:   salary,
		NextSet:  next,
		LastSet:  last,
		OrderSet: make([]*ActRoute, 0),
	}
}

func (act *Activity) AddNext(next *Activity) {
	act.NextSet = append(act.NextSet, next)
}

func (act *Activity) AddLast(next *Activity) {
	act.LastSet = append(act.LastSet, next)
}

// 整理指定节点为起点的活动路线
func (act *Activity) OrderActRoute() {

	if len(act.NextSet) == 0 {
		newRoute := &ActRoute{TotalConsume: act.Consume, TotalSalary: act.Salary}
		act.OrderSet = append(act.OrderSet, newRoute)
		return
	}

	for _, nextAct := range act.NextSet { // 遍历当前活动所有的后续活动
		for _, actRoute := range nextAct.OrderSet { // 遍历每一个后续活动的有序活动路线
			newRoute := &ActRoute{TotalConsume: act.Consume + actRoute.TotalConsume,
				TotalSalary: act.Salary + actRoute.TotalSalary}
			act.OrderSet = append(act.OrderSet, newRoute)
		}
	}

	// 将所有活动路线按照消耗天数进行排序
	sort.Slice(act.OrderSet, func(i, j int) bool {
		if act.OrderSet[i].TotalConsume < act.OrderSet[j].TotalConsume {
			return true
		} else {
			return false
		}
	})

	if len(act.OrderSet) == 1 {
		return
	}

	// 排除掉所有天数增多，但报酬不递增的活动路线
	targetRoutes := make([]*ActRoute, 0)
	index := 0
	for i := 0; i < len(act.OrderSet); i++ {
		if len(targetRoutes) == 0 {
			targetRoutes = append(targetRoutes, act.OrderSet[0])
			continue
		}
		if act.OrderSet[i].TotalSalary > targetRoutes[index].TotalSalary {
			targetRoutes = append(targetRoutes, act.OrderSet[i])
			index++
		}
	}
}

func SumPerfectRoutes(starts []*Activity) []*ActRoute {
	allActRoutes := make([]*ActRoute, 0)
	for _, start := range starts { // 遍历每一个活动
		for i := 0; i < len(start.OrderSet); i++ { // 遍历每一个活动的有效活动路线
			allActRoutes = append(allActRoutes, start.OrderSet...)
		}
	}

	sort.Slice(allActRoutes, func(i, j int) bool {
		if allActRoutes[i].TotalConsume < allActRoutes[j].TotalConsume {
			return true
		} else {
			return false
		}
	})

	targetRoutes := make([]*ActRoute, 0)
	index := 0
	for i := 0; i < len(allActRoutes); i++ {
		if len(targetRoutes) == 0 {
			targetRoutes = append(targetRoutes, allActRoutes[0])
			continue
		}
		if allActRoutes[i].TotalSalary > targetRoutes[index].TotalSalary {
			targetRoutes = append(targetRoutes, allActRoutes[i])
			index++
		}
	}

	return targetRoutes
}
```

## 四、题目四

给定一个只由 0（假）、1（真）、&（逻辑与）、|（逻辑或）、和^（异或）五种字符组成的字符串express，再给定一个布尔值desired。返回express能有多少种组合的方式，可以达到desired的结果？

```
提示：这个问题其实就是在问怎么加()字符串中的各个关系运算式最终得到desired
```

1. 方法一（暴力递归）

```go
1.首先要判断字符串的有效性，有效性需要满足下面三个点：
1）偶数下标位置(0、2、4、……)必须是0或者1
2）奇数下标位置(1、3、5、……)必须是位运算符
3) 字符串总长度必须是奇数(如 1&0、 1|0&1)

2.外循环每次都访问奇数下标位置(也就是位运算符)，i从1开始，i每轮后+2
3.外循环访问到位运算符后，根据desired是true还是false分为不同的分支，需要通过递归的方式分别获取位运算符左右返回指定bool值的数量

func LogicalOpt(logicStr string, desire bool) int {
	if logicStr == "" {
		return 0
	}
	if !isValid(logicStr) {
		return 0
	}

	return logicalOpt(logicStr, desire, 0, len(logicStr)-1)

}

func logicalOpt(logicStr string, desire bool, left, right int) int {
	// 边界条件
	if left > right {
		return 0
	}

	if left == right { // 此时只剩一个字符
		if desire { // 目标渴望是true
			if logicStr[left] == '1' {
				return 1
			} else if logicStr[left] == '0' {
				return 0
			}
		} else { // 目标渴望是false
			if logicStr[left] == '1' {
				return 0
			} else if logicStr[left] == '0' {
				return 1
			}
		}
	}

	// 正常递归
	res := 0 // 可能的结果数

	if desire { // 目标渴望是true
		for i := left + 1; i < right; i += 2 { // 遍历每一个逻辑运算符(注意：这里每次只需要遍历当前区域内(left~right)的奇数位置)
			op := logicStr[i]
			switch op {
			case '&': // 左右必须都是true
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '|': // 左右有一个是true即可
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '^': // 左右必须相异
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			}
		}
	} else { // 目标渴望是false
		for i := left + 1; i < right; i += 2 { // 遍历每一个逻辑运算符
			switch logicStr[i] {
			case '&': // 左右有一个是false即可
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, false, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, true, i+1, right)
			case '|': // 左右必须都是false
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
			case '^': // 左右必须相同
				res += logicalOpt(logicStr, true, left, i-1) * logicalOpt(logicStr, true, i+1, right)
				res += logicalOpt(logicStr, false, left, i-1) * logicalOpt(logicStr, false, i+1, right)
			}
		}
	}
	return res
}
```

2. 方法二（动态规划）

```go
1.可变参量只有L和R，因此是一个二维表
2.期待的结果desired可能是true，也可能是false，因此需要有两张表
3.L不可能小于R，因此表格的下半区域都是无效的
4.对角线区域，意味着 L == R , 因此对角线区域就是初始条件可以获得的
5.因为最后一行只有一个位置需要求解，也就是对角线，而对角线已知，因此求解的顺序是从下往上，从左往右。

func LogicalOptDP(logicStr string, desire bool) int {
	if logicStr == "" {
		return 0
	}
	if !isValid(logicStr) {
		return 0
	}

	return logicalOptDP(logicStr, desire)

}

func logicalOptDP(logicStr string, desire bool) int {
	N := len(logicStr)

	trueDP := make([][]int, N)
	for i := 0; i < N; i++ {
		trueDP[i] = make([]int, N)
	}

	falseDP := make([][]int, N)
	for i := 0; i < N; i++ {
		falseDP[i] = make([]int, N)
	}

	// 1.设置初始条件(矩阵对角线) -- 两个矩阵(一个是最终返回结果是true，一个是最终返回结果为false)
	for i := 0; i < N; i += 2 { // 每次都要访问数字(left和right都在偶数位置上)
		if logicStr[i] == '0' {
			trueDP[i][i] = 0
			falseDP[i][i] = 1
		} else if logicStr[i] == '1' {
			trueDP[i][i] = 1
			falseDP[i][i] = 0
		}
	}

	// 2.根据初始条件获取其他位置(从下往上，从左向右求解)
	for left := N - 3; left >= 0; left -= 2 { // 范围的左边界(需要是数字)  N-1作为对角线是初始条件不用求
		for right := left + 2; right < N; right += 2 { // 范围的右边界(需要是数字)
			for oper := left + 1; oper < right; oper += 2 { // 遍历从左边界到右边界范围内的所有逻辑运算符
				// 2.1 计算trueDP
				switch logicStr[oper] {
				case '&': // logicStr[oper] 左右两侧都必须是true
					trueDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
				case '|':
					trueDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
					trueDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					trueDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				case '^':
					trueDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					trueDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				}
				// 2.2 计算falseDP
				switch logicStr[oper] {
				case '&':
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
					falseDP[left][right] += trueDP[left][oper-1] * falseDP[oper+1][right]
					falseDP[left][right] += falseDP[left][oper-1] * trueDP[oper+1][right]
				case '|':
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
				case '^':
					falseDP[left][right] += trueDP[left][oper-1] * trueDP[oper+1][right]
					falseDP[left][right] += falseDP[left][oper-1] * falseDP[oper+1][right]
				}
			}
		}
	}

	// 3.返回结果(左边界为0，右边界为N-1的矩阵元素即是要求解的目标)
	if desire {
		return trueDP[0][N-1]
	} else {
		return falseDP[0][N-1]
	}

}

// 判断一个逻辑字符串是否有效
func isValid(logicStr string) bool {
	length := len(logicStr)

	if length%2 == 0 { // 字符串必须是奇数长度
		return false
	}

	for i := 0; i < length; i += 2 { // 偶数位置必须是0或1
		if logicStr[i] != '0' && logicStr[i] != '1' {
			return false
		}
	}

	for i := 1; i < length; i += 2 { // 奇数位置必须是逻辑运算符 & | ^
		if logicStr[i] != '&' && logicStr[i] != '|' && logicStr[i] != '^' {
			return false
		}
	}

	return true
}
```

## 五、题目五

在一个字符串中找到没有重复字符子串中最长的长度。

例如：

`abcabcbb` 没有重复字符的最长子串是 `abc` ，长度是3

`bbbbb`，答案是b，长度为1

`pwwkew`，答案是`wke`，长度为3

```go
遍历整个字符串arr。
假设当前遍历的是arr[i]，查询以arr[i]作为结尾的最长不重复子串是什么。
以arr[i]为结尾的不重复子串的长度取决于两个量：
1.假设arr[i]上的字符是'a',其中一个变量就是arr[i]左侧最近一次'a'出现的位置。
2.以arr[i-1]字符结尾的不重复子串的长度。

以arr[i]结尾的不重复子串的左边界取自上述两种情况中，距离arr[i]最近的那一个

package lesson9

import "math"

func MaxUnique(str string) int {
	if str == "" {
		return 0
	}

	recentCharMap := make(map[uint8]int) // 用于记录不同字符出现在字符串中的距离结尾最近的位置

	for i := 0; i <= 255; i++ {
		recentCharMap[uint8(i)] = -1
	}

	maxLen := math.MinInt
	lastLeft := -1 // i-1 位置的无重复子串的左边界位置
	index := 0     // 访问下标i

	for {
		if index >= len(str) {
			return maxLen
		}
		lastLeft = getMax(lastLeft, recentCharMap[str[index]]) // 当前以index结尾的无重复字符串的左边界
		recentCharMap[str[index]] = index
		curLen := index - lastLeft
		maxLen = getMax(maxLen, curLen)
		index++
	}

}
```

## 六、题目六

给定两个字符串str1和str2，再给定三个整数ic 、dc和rc，分别代表插入、删除和替换一个字符串的代价，返回将str1编辑成str2的最小代价。

举例：

1. str1 == "abc" , str2 == "adc"  , ic = 5 , dc =3 , rc = 2

​       从"abc"编辑成"adc" , 把'b' 替换成'd'是代价最小的，所以返回2

2. str1 == "abc" , str2 == "adc"  , ic = 5 , dc =3 , rc = 100

   从"abc"编辑成"adc" , 先删除'b'，然后再插入'd'是代价最小的，所以返回8

3. str1 == "abc" , str2 == "abc"  , ic = 5 , dc =3 , rc = 2

   因为本来就是一样的字符串，所以不用编辑，返回0

```go
这是一个编辑距离问题，需要用动态规划解决。

假设str1 == "abcdef"   str2 == "skbcdf"

1.下面的dp表，行表示str1的各个字符；列表示str2的各个字符

2.动态规划表dp的每一个元素 dp[i][j] 表示：
str1中以str1[i]结尾的子串完全转化为str2中以str2[j]结尾的子串需要的代价。

3.第一行表示str1的空串要与str2对应字符结尾的子串相等，那么str1的子串只能执行插入操作

4.第一列表示str2的空串要与str1对应字符结尾的子串相等，那么str1的子串只能执行删除操作
```

|       |  0   | 1-"s" | 2-"k" | 3-"b" | 4-"c" | 5-"d" | 6-"f" |
| :---: | :--: | :---: | :---: | :---: | :---: | :---: | :---: |
|   0   |  0   |  ic   | 2*ic  | 3*ic  | 4*ic  | 5*ic  | 6*ic  |
| 1-"a" |  dc  |       |       |       |       |       |       |
| 2-"b" | 2*dc |       |       |       |       |       |       |
| 3-"c" | 3*dc |       |       |       |       |       |       |
| 4-"d" | 4*dc |       |       |       |       |       |       |
| 5-"e" | 5*dc |       |       |       |       |       |       |
| 6-"f" | 6*dc |       |       |       |       |       |       |

```go
对于余下的普通位置，也就是常规的dp[i][j]的求法分为四种：
1.将以 str1[i-1]结尾的子串编辑成以str2[j]结尾的子串，然后删除str1[i] , 总代价为 dp[i-1][j] + dc (依赖于上方格子)
2.将以 str1[i]结尾的子串编辑成以str2[j-1]结尾的子串，然后在str1[i]上添加一个str2[j], 总代价为 dp[i][j-1] + ic (依赖于左侧格子)
3.将以 str1[i-1]结尾的子串编辑成以str2[j-1]结尾的子串,然后用str2[j]替换str1[i],总代价为 dp[i-1][j-1] + rc (依赖于左上角格子)
4.当str1[i] == str2[j]这种特殊情况下，将以 str1[i-1]结尾的子串编辑成以str2[j-1]结尾的子串 ，总代价为 dp[i-1][j-1](依赖于左上角格子)

每一个dp[i][j] 所需要的代价是上述4种代价中最小的那一个。
```

```go
package lesson9

import "math"

// 将str1编辑成str2，消耗的最小代价
func StrEditDistance(str1, str2 string, icost, dcost, rcost int) int {
	rowCount := len(str1) + 1
	colCount := len(str2) + 1

	dp := make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		dp[i] = make([]int, colCount)
	}

	// 1.初始条件（第一行和第一列是已知的）
	for col := 0; col < colCount; col++ {
		dp[0][col] = col * icost // 第一行，str1为空，str1变成str2只能插入
	}
	for row := 1; row < rowCount; row++ {
		dp[row][0] = row * dcost // 第一行，str2为空，str1变成str2只能删除
	}

	// 2.普通dp求解
	for row := 1; row < rowCount; row++ {
		for col := 1; col < colCount; col++ {
			cost := math.MaxInt
			// 2.1 结尾字符相等的情况，依赖于左上角的dp[row-1][col-1]
			if str1[row-1] == str2[col-1] { // 注意：字符串索引下标要比矩阵中的行与列小1(因为矩阵引入了0长度字符串概念)
				cost = getMin(cost, dp[row-1][col-1])
			}
			// 2.2 将str1[0……i-1]变成str2[0……j]，然后将str1[i]删除
			cost = getMin(cost, dp[row-1][col]+dcost)

			// 2.3 将str1[0……i]变成str2[0……j-1]，然后在str1[i]之后新加一个str2[j]
			cost = getMin(cost, dp[row][col-1]+icost)

			// 2.4 将str1[0……i-1]变成str2[0……j-1]，然后将str1[i]变成str2[j]
			cost = getMin(cost, dp[row-1][col-1]+rcost)

			dp[row][col] = cost
		}
	}

	// 3.返回结果
	return dp[rowCount-1][colCount-1]
}
```

## 七、题目七

给定一个全是小写字母的字符串str，删除多余字符，使得每种字符只保留一个，并让最终结果字符串的字典序最小。

举例：

str = "acbc"，删除第一个'c'，得到"abc"，是所有结果字符串中字典序最小的。

str = "dbcacbca"，删除第一个'b'、第一个'c'、第二个'c'、第二个'a'，得到"dabc"，是所有结果字符串中字典序最小的。

```go
1.先遍历一遍字符串，组建一个词频表
2.准备一个变量minACSIndex,用来记录每一轮中ACSII码最小的字符的下标。这个变量开始为0
3.用一个新字符串res保留结果字符串
3.每一轮都从头开始遍历字符串，每访问一个字符，将其在词频表中的记录--。一旦遇到词频 == 0 的情况，退出本轮循环，让res记录本轮最小ascii码的字符，res += str[minACSInex],然后将minACSInex以及其之前的字符全部变为空字符''，同时将后续字符串中的 str[minACSInex] 字符也变成空字符''
4.接着进行下一轮循环。下一轮循环从minACSInex+1开始，重新组建词频表，然后访问。
5.结束条件：直到字符串长度为1（将这个字符加入到res中）
```

```go
package lesson9

func Operation(str string) string {
	if str == "" {
		return ""
	}
	freqMap := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		freqMap[str[i]]++
	}
	minACSIndex := 0
	resSubStr := ""
	curStr := str
	for {
		if len(curStr) == 1 {
			resSubStr += curStr
			return resSubStr
		}
		for i := 0; i < len(curStr); i++ {
			freqMap[curStr[i]]--
			if freqMap[curStr[i]] == 0 {
				// 1.计算在出现词频为0时，之前字符串中具有最小ascii码的下标
				minAscii := uint8(255)
				for index := 0; index <= i; index++ {
					if curStr[index] < minAscii {
						minAscii = curStr[index]
						minACSIndex = index
					}
				}
				// 2.字符追加
				resSubStr += string(curStr[minACSIndex])
				// 3.冗余字符删除
				newCurStr := curStr[minACSIndex+1:]   // 删除掉包含minACSIndex在内的之前的所有字符
				for j := 0; j < len(newCurStr); j++ { // 删除从 minACSIndex+1开始到末尾的所有 curStr[minACSIndex]字符
					if newCurStr[j] == curStr[minACSIndex] {
						newCurStr = newCurStr[:j] + newCurStr[j+1:]
					}
				}
				curStr = newCurStr
				break
			}
		}
		// 重新组件词频表
		freqMap = make(map[uint8]int)
		for i := 0; i < len(curStr); i++ {
			freqMap[curStr[i]]++
		}
	}

}
```
