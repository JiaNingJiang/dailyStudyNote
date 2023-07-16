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
