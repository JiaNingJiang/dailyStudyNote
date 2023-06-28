package lesson1

import "math"

// grass 是草的份数
func SheepEatGrass(grass int) string {
	// 设置边界条件(边界条件越多，递归花费的时间越少。 只有0和1两个边界条件是必须的)
	// 0  1  2  3  4
	// 后 先 后 先  先
	if grass == 0 || grass == 2 {
		return "后手"
	} else if grass == 1 || grass == 3 || grass == 4 {
		return "先手"
	}

	first := 1 // 先手羊吃的份数

	for {
		if first > grass { // 先手羊遍历了所有可能，也不能赢
			return "后手"
		}
		// 后手羊的先手函数返回后手，就等同于先手羊获得胜利
		if SheepEatGrass(grass-first) == "后手" {
			return "先手"
		}

		if first > math.MaxInt/4 { // 下一次先手羊吃的份数将溢出
			return "后手"
		}
		first *= 4

	}

}

func SheepEatGrass2(n int) string {
	if n%5 == 0 || n%5 == 2 {
		return "后手"
	} else {
		return "先手"
	}
}
