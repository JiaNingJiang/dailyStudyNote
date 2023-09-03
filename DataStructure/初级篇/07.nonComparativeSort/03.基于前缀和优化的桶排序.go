package nonComparativeSort

import (
	"fmt"
	"strconv"
)

func ImprovedBucketSort(arr []int) {
	strArr := make([]string, 0, len(arr))
	eleMaxLen := 0            // 位数最多元素的位数
	for _, ele := range arr { // 将arr数组转化为各元素数值相等的字符串数组
		eleStr := fmt.Sprintf("%d", ele)
		if len(eleStr) > eleMaxLen {
			eleMaxLen = len(eleStr)
		}
		strArr = append(strArr, eleStr)
	}

	// 将其余各元素的位数补齐至eleMaxLen位(前方补0)
	for index, eleStr := range strArr {
		eleLen := len(eleStr)
		diff := eleMaxLen - eleLen
		if diff != 0 {
			for i := 0; i < diff; i++ {
				strArr[index] = "0" + strArr[index] // 重要：修改的是strArr[index]，而非eleStr
			}
		}
	}

	// 从这里开始改进算法
	bucket := make([]string, len(strArr)) // 代替桶的临时数组
	count := make([]int, baseNum)         // 前缀和数组

	for i := eleMaxLen - 1; i >= 0; i-- { // 从个位开始，进行排序

		for index, _ := range count { // 重要：每一轮运行前，清空前缀和数组count
			count[index] = 0
		}

		for _, eleStr := range strArr { // 第一次遍历待排序数组(顺序无所谓)，填充前缀和数组count
			bitCount, _ := strconv.ParseInt(string(eleStr[i]), baseNum, 32)
			for j := int(bitCount); j < len(count); j++ { //
				count[j]++
			}
		}

		for k := len(strArr) - 1; k >= 0; k-- { // 第二次遍历(必须从后往前)，相当于同时完成了入桶+出桶
			bitCount, _ := strconv.ParseInt(string(strArr[k][i]), baseNum, 32)
			suffixCount := count[bitCount]
			bucket[suffixCount-1] = strArr[k] // 等价于有桶的后进后出(后进桶的排在后面)
			count[bitCount]--
		}
		// 相当于完成了一轮 针对一位的排序。用bucket更新strArr
		for index, eleStr := range bucket {
			strArr[index] = eleStr
		}

	}

	// 完成排序，将元素从字符串转换回数字形式
	for i, eleStr := range strArr {
		ele, _ := strconv.ParseInt(eleStr, baseNum, 32)
		arr[i] = int(ele)
	}

}
