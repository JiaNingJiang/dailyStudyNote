## 一、题目一

在数据加密和数据压缩中常需要对特殊的字符串进行编码。给定的字母表A由26个小写英文字符{a,b,c,……，z}组成。由该字母表产生的长序字符串必须是最小字母序的，每个符号最多出现一次。

整体的编码顺序为：a(1)，b(2)，c(3)，……，z(26)，ab(27)，ac(28)，……，az，bc，……，bz，yz，abc……，

假设给你一个任意的编码字符串，请你给出这是第几个编码字符串？

```go
1.长度小的一定排在长度达的后面：先是长度为1的，接着是为2的，为3的，……，为26的
2.每一个字符串必定是按照字母序的

先给出下面两个函数：
1.函数一：可以求出以i号字符开头，总长度为len的子序列有多少个？
2.函数二：可以求出总长度为len的子序列有多少个？

// 求出以begin开头(begin: 0~25 表示26个不同字母)，长度为length的字符串个数
func spBeginFixedLen(begin int, length int) int {
	if length == 1 { // 长度为1，开头又固定，那必然只有一种可能
		return 1
	}

	sum := 0
	for i := begin + 1; i <= 25; i++ { // 子字符串的开头必须从 begin+1 开始，且长度-1
		sum += spBeginFixedLen(i, length-1)
	}

	return sum
}

// 求出指定长度的所有可能字符串的数量
func fixedLen(length int) int {
	sum := 0

	for i := 0; i <= 25; i++ {
		sum += spBeginFixedLen(i, length)
	}
	return sum
}

func DesignatedStrIndex(str string) int {
	index := 0
	strLen := len(str)

	// 长度为1的字符串求解是特殊的
	if strLen == 1 {
		end := int(str[0] - 'a')
		for i := 0; i < end; i++ {
			index++
		}
		return index + 1
	}

	for i := 1; i <= strLen-1; i++ {
		index += fixedLen(i) // 长度比当前字符串小的必然在前面
	}

	curStr := str
	for {
		if len(curStr) <= 1 {
			return index + 1 // index只是当前字符串之前的字符串个数，因此还需要+1
		}
		left := int(curStr[0]-'a') + 1 // 长度至少为2
		right := int(curStr[1]-'a') - 1

		for j := left; j <= right; j++ {
			index += spBeginFixedLen(j, len(curStr)-1)
		}

		curStr = curStr[1:] // 每次循环去掉一个最左侧的字符

	}

}
```

使用上面两个函数可以求任意一个子序列的位置：

```go
假设要求解的子序列为 "d,j,v"

1.我们先求总长度为1和2的子序列有多少个，他们一定在"d,j,v"这个总长度为3的子序列的前面
2.再求分别以 "a" 、"b"、……、"e"开头的总长度为3的子序列一共有多少个，他们也一定在"d,j,v"这个总长度为3的子序列的前面
3.再求分别以 “e”（d的下一个）、……、“i”（j的前一个）开头的总长度为2的子序列一共有多少个，他们也一定在"d,j,v"这个总长度为3的子序列的前面
4.再求分别以 “k”（j的下一个）、……、“u”（v的前一个）开头的总长度为1的子序列一共有多少个，他们也一定在"d,j,v"这个总长度为3的子序列的前面

将上述获得的子序列数量 + 1，即可得到目标子序列的位置
```

