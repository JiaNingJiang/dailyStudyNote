package lesson5

import "math"

func HoldWater(vessel []int) int {

	leftMax := make([]int, len(vessel))
	rightMax := make([]int, len(vessel))

	max := math.MinInt
	for i := 0; i < len(vessel); i++ {
		if vessel[i] > max {
			max = vessel[i]
		}
		leftMax[i] = max
	}
	max = math.MinInt
	for i := len(vessel) - 1; i >= 0; i-- {
		if vessel[i] > max {
			max = vessel[i]
		}
		rightMax[i] = max
	}

	waterSet := make([]int, len(vessel))
	totalWater := 0
	for i := 1; i < len(vessel)-1; i++ { // 容器的左右两侧是无法存水的，因此跳过即可
		water_i := getMin(leftMax[i], rightMax[i]) - vessel[i]
		waterSet[i] = water_i
		totalWater += water_i
	}
	return totalWater
}

func HoldWaterByPoint(vessel []int) int {
	leftMax := vessel[0]
	rightMax := vessel[len(vessel)-1]

	leftIndex := 1                // 最左侧位置不能蓄水，因此跳过
	rightIndex := len(vessel) - 2 // 最右侧位置不能蓄水，因此跳过

	waterSet := make([]int, len(vessel))
	totalWater := 0

	for {
		if leftIndex > rightIndex {
			break
		}

		if vessel[leftIndex] > leftMax { // leftIndex位置比之前左侧所有位置都高，那么leftIndex位置无法存水
			waterSet[leftIndex] = 0
			leftMax = vessel[leftIndex]
			leftIndex++ // 当前位置蓄水值已求出，因此可以继续leftIndex++求下一个为止
		} else {
			if leftMax <= rightMax { // leftIndex位置水位取决于leftMax
				waterSet[leftIndex] = leftMax - vessel[leftIndex]
				leftIndex++ // 当前位置蓄水值已求出，因此可以继续leftIndex++求下一个为止
			}
		}

		if vessel[rightIndex] > rightMax {
			waterSet[rightIndex] = 0
			rightMax = vessel[rightIndex]
			rightIndex-- // 当前位置蓄水值已求出，因此可以继续rightIndex--求下一个为止
		} else {
			if rightMax <= leftMax {
				waterSet[rightIndex] = rightMax - vessel[rightIndex]
				rightIndex-- // 当前位置蓄水值已求出，因此可以继续rightIndex--求下一个为止
			}
		}
	}

	for i := 0; i < len(waterSet); i++ {
		totalWater += waterSet[i]
	}
	return totalWater
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
