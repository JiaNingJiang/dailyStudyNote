package bruteRecursionPromote

// 求解从(0,0)到(x,y)的走法，等价于求从(x,y)到(0,0)的走法
func Chess(x, y int, step int, lineMax, colMax int, lineMin, colMin int) int {
	if x < lineMin || x > lineMax || y < colMin || y > colMax {
		return 0
	}

	if step == 0 {
		if x == lineMin && y == colMin {
			return 1
		} else {
			return 0
		}
	}

	return Chess(x+2, y+1, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x+1, y+2, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x-1, y+2, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x-2, y+1, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x-2, y-1, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x-1, y-2, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x+1, y-2, step-1, lineMax, colMax, lineMin, colMin) +
		Chess(x+2, y-1, step-1, lineMax, colMax, lineMin, colMin)
}

func ChessStrictTable(x, y int, step int, lineMax, colMax int, lineMin, colMin int) int {
	if x < lineMin || x > lineMax || y < colMin || y > colMax || step < 0 {
		return 0
	}

	table := make([][][]int, step+1) // 第一维度是高度( 0~step 共计step+1层)
	for h := 0; h <= step; h++ {
		table[h] = make([][]int, lineMax-lineMin+1) // 第二三维度分别是棋盘的宽度(行数)和长度(列数)
		for i := lineMin; i <= lineMax; i++ {       // 行数 lineMin~lineMax
			table[h][i] = make([]int, colMax-colMin+1)
			for c := colMin; c <= colMax; c++ { // 列数 colMin~colMax
				table[h][i][c] = 0
			}
		}
	}

	table[0][lineMin][colMin] = 1 // 从最下层开始往上推，最下层只有原点值为1，其余各点都为0

	for h := 1; h <= step; h++ {
		for i := lineMin; i <= lineMax; i++ {
			for c := colMin; c <= colMax; c++ {
				table[h][i][c] = getValue(table, i+2, c+1, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i+1, c+2, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i-1, c+2, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i-2, c+1, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i-2, c-1, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i-1, c-2, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i+1, c-2, h-1, lineMin, lineMax, colMin, colMax) +
					getValue(table, i+2, c-1, h-1, lineMin, lineMax, colMin, colMax)
			}
		}
	}
	return table[step][x][y]
}

func getValue(table [][][]int, line, col, height int, lineMin, lineMax, colMin, colMax int) int {
	if line < lineMin || line > lineMax || col < colMin || col > colMax {
		return 0
	}

	return table[height][line][col]
}
