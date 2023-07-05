package lesson4

import "math"

func WashMachine(washer []int) int {
	totalCloth := 0
	for _, laundry := range washer {
		totalCloth += laundry
	}
	washerCount := len(washer)
	avgCloth := totalCloth / washerCount // 平均每台洗衣机要洗的衣物

	roundSet := make([]int, washerCount) // 以每台洗衣机为中心，达成待洗衣物的均分需要的轮次各自是多少

	for i := 0; i < washerCount; i++ {
		leftLaundry := 0 // 左侧已有的待洗衣物
		for left := 0; left < i; left++ {
			leftLaundry += washer[left]
		}
		rightLaundry := 0 // 右侧已有的待洗衣物
		for right := i + 1; right < washerCount; right++ {
			rightLaundry += washer[right]
		}
		leftAdjust := leftLaundry - avgCloth*i                   // 左侧需要调整的待洗衣物数量
		rightAdjust := rightLaundry - avgCloth*(washerCount-1-i) // 右侧需要调整的待洗衣物数量

		if leftAdjust < 0 && rightAdjust < 0 { // 左右都需要移入待洗衣物
			roundSet[i] = abs(leftAdjust) + abs(rightAdjust)
		} else if leftAdjust > 0 && rightAdjust > 0 { // 左右都需要移入待洗衣物
			roundSet[i] = getMax(leftAdjust, rightAdjust)
		} else if leftAdjust > 0 && rightAdjust < 0 {
			roundSet[i] = getMax(leftAdjust, abs(rightAdjust))
		} else if leftAdjust < 0 && rightAdjust > 0 {
			roundSet[i] = getMax(abs(leftAdjust), rightAdjust)
		} else if leftAdjust == 0 && rightAdjust == 0 {
			roundSet[i] = 0
		} else if leftAdjust == 0 && rightAdjust != 0 {
			roundSet[i] = abs(rightAdjust)
		} else if leftAdjust != 0 && rightAdjust == 0 {
			roundSet[i] = abs(leftAdjust)
		}
	}

	minRound := math.MaxInt
	for _, round := range roundSet {
		if round < minRound {
			minRound = round
		}
	}
	return minRound
}

func abs(before int) int {
	if before < 0 {
		return -before
	} else {
		return before
	}
}

func getMax(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
