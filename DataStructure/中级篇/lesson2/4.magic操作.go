package lesson2

import (
	"fmt"
	"sort"
	"strconv"
)

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
