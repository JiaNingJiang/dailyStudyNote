package bruteRecursionPromote

func CardWin(cards []int) int {
	if cards == nil || len(cards) == 0 {
		return -1
	}
	return getMax(First(cards, 0, len(cards)-1), Second(cards, 0, len(cards)-1))
}

func First(cards []int, left, right int) int {
	if left == right {
		return cards[left]
	}
	leftScore := cards[left] + Second(cards, left+1, right)
	rightScore := cards[right] + Second(cards, left, right-1)

	return getMax(leftScore, rightScore)
}

func Second(cards []int, left, right int) int {
	if left == right {
		return 0
	}
	leftScore := First(cards, left+1, right)
	rightScore := First(cards, left, right-1)

	return getMin(leftScore, rightScore)
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

func CardWinStrictTable(cards []int) int {

	if cards == nil || len(cards) == 0 {
		return -1
	}

	lineMax := len(cards) - 1
	colMax := len(cards) - 1

	firstTable := make([][]int, lineMax+1)
	for i := 0; i <= lineMax; i++ {
		firstTable[i] = make([]int, colMax+1)
	}

	secondTable := make([][]int, lineMax+1)
	for i := 0; i <= lineMax; i++ {
		secondTable[i] = make([]int, colMax+1)
	}

	// 先求先手表的初始
	for line := 0; line <= lineMax; line++ {
		for col := 0; col <= colMax; col++ {
			if line == col {
				firstTable[line][col] = cards[line]
			}
		}
	}

	// 再求后手表的初始
	for line := 0; line <= lineMax; line++ {
		for col := 0; col <= colMax; col++ {
			if line == col {
				secondTable[line][col] = 0
			}
		}
	}

	curLine := 0
	diff := 1
	curCol := curLine + diff

	for {
		if curLine == 0 && curCol == colMax+1 {
			break
		}

		secondTable[curLine][curCol] = getMin(firstTable[curLine][curCol-1], firstTable[curLine+1][curCol])
		firstTable[curLine][curCol] = getMax(cards[curLine]+secondTable[curLine+1][curCol],
			cards[curCol]+secondTable[curLine][curCol-1])

		curLine++
		curCol++
		if curCol > colMax {
			curLine = 0
			diff++
			curCol = curLine + diff
		}
	}

	return getMax(firstTable[0][colMax], secondTable[0][colMax])
}
