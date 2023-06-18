package bruteRecursionPromote

import "math"

// 以(x,y)作为起点进行移动step步，一旦越界就会死亡
func BeyondEdge(x, y int, step int, lineMax, colMax int, lineMin, colMin int) float64 {
	totalCase := math.Pow(float64(4), float64(step)) // 每次移动都有 上下左右 四种选择，总共有 4^step种移动方式

	surviveCase := beyondEdge(x, y, step, lineMax, colMax, lineMin, colMin) // 存活的可能性

	return float64(surviveCase) / totalCase
}

func beyondEdge(x, y int, step int, lineMax, colMax int, lineMin, colMin int) int {
	if x < lineMin || x > lineMax || y < colMin || y > colMax {
		return 0
	}
	if step == 0 { // 只要走完step步没越界就是存活
		return 1
	}

	return beyondEdge(x-1, y, step-1, lineMax, colMax, lineMin, colMin) + // 向左
		beyondEdge(x+1, y, step-1, lineMax, colMax, lineMin, colMin) + // 向右
		beyondEdge(x, y-1, step-1, lineMax, colMax, lineMin, colMin) + // 向下
		beyondEdge(x, y+1, step-1, lineMax, colMax, lineMin, colMin) // 向上
}

func BeyondEdgeStrictTable(x, y int, step int, lineMax, colMax int, lineMin, colMin int) float64 {
	totalCase := math.Pow(float64(4), float64(step)) // 每次移动都有 上下左右 四种选择，总共有 4^step种移动方式

	table := make([][][]int, step+1) // 第一维度是高度( 0~step 共计step+1层)
	for h := 0; h <= step; h++ {
		table[h] = make([][]int, lineMax-lineMin+1) // 第二三维度分别是区域的宽度(行数)和长度(列数)
		for i := lineMin; i <= lineMax; i++ {       // 行数 lineMin~lineMax
			table[h][i] = make([]int, colMax-colMin+1)
			for c := colMin; c <= colMax; c++ { // 列数 colMin~colMax
				table[h][i][c] = 0
			}
		}
	}
	for i := lineMin; i <= lineMax; i++ { // 最后一层全部为1(因为只要在区域内就是存活)
		for c := colMin; c <= colMax; c++ {
			table[0][i][c] = 1
		}
	}

	for h := 1; h <= step; h++ {
		for i := lineMin; i <= lineMax; i++ { // 最后一层全部为1(因为只要在区域内就是存活)
			for c := colMin; c <= colMax; c++ {
				table[h][i][c] = getValue(table, i-1, c, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i+1, c, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i, c-1, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i, c+1, h-1, lineMin, lineMax, colMin, colMax)
			}
		}
	}

	return float64(table[step][x][y]) / totalCase
}
