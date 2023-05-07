package nonComparativeSort

import (
	"fmt"
	"strconv"
)

const (
	baseNum = 10 // 默认要统计的数组元素都是十进制
)

func BucketSort(arr []int) {

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

	buckets := make([][]string, baseNum) // 事先创建与进制数量相等的 "桶(bucket)"

	for i := eleMaxLen - 1; i >= 0; i-- { // 循环次数 == 位数最多元素的位数  (从最低位开始--- 重要：元素的最高位是位数的最低位)
		for _, eleStr := range strArr { // 第一次遍历，将所有元素 "入桶"
			bucketNum, _ := strconv.ParseInt(string(eleStr[i]), baseNum, 32) // 找到当前位需要存放的bucket的编号
			buckets[bucketNum] = append(buckets[bucketNum], eleStr)
		}
		// 将所有元素从桶中排出（每个桶都要按照先进先出的原则），完成i位的排序
		strArr = make([]string, 0, len(arr))
		for index, bucket := range buckets {
			strArr = append(strArr, bucket...)
			buckets[index] = []string{} // 重要：对应的桶清空上一次的缓存
		}
	}

	// 完成排序，将元素从字符串转换回数字形式
	for i, eleStr := range strArr {
		ele, _ := strconv.ParseInt(eleStr, baseNum, 32)
		arr[i] = int(ele)
	}

}
