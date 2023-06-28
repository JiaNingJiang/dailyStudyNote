package lesson1

import "fmt"

func MinBags(apples int) int {

	if apples <= 0 || apples%2 != 0 {
		return -1
	}

	bag6 := -1         // 使用的6容积袋子数目，初始为-1
	bag8 := apples / 8 //  使用的8容积袋子数目，初始为apples/8

	remain := apples - bag8*8 // 剩余需要6容积袋子装的苹果数

	for {
		if bag8 < 0 || remain >= 24 { // 凑不齐
			return -1
		}
		curBag6 := getBag6Count(remain)
		if curBag6 != -1 {
			bag6 = curBag6
			break
		}
		bag8--
		remain = apples - bag8*8
	}
	fmt.Printf("苹果数:%d  6袋子:%d   8袋子:%d\n", apples, bag6, bag8)
	return bag6 + bag8
}

// 剩余苹果能否正好全部用6容积袋子装起来，如果可以返回需要的袋子数
func getBag6Count(remain int) int {
	if remain%6 == 0 {
		return remain / 6
	} else {
		return -1
	}
}

func minBags2(apples int) int {
	if apples == 0 || apples%2 != 0 {
		return -1
	}
	if apples < 18 {
		switch apples {
		case 0:
			return 0
		case 6, 8:
			return 1
		case 12, 14, 16:
			return 2
		default:
			return -1
		}
	} else {
		return (apples-18)/8 + 3
	}
}
