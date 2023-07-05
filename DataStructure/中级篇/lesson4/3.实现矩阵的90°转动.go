package lesson4

import "fmt"

func NinetyRotation(matrix *[][]int) {
	maxRow := len(*matrix) - 1
	maxCol := len((*matrix)[0]) - 1

	luRow := 0      // 左上角的行号
	luCol := 0      // 左上角的列号
	rlRow := maxRow // 右下角的行号
	rlCol := maxCol // 右下角的列号

	for {
		if luRow > rlRow || luCol > rlCol { // 对于一个正方形矩阵，不可能出现左上角与右下角同行或同列，因此没有 ==
			return
		}

		group := rlCol - luCol // 矩阵的当前圈被分成的组数，每一组固定会有4个点
		for i := 0; i < group; i++ {
			// 四个点的位置分别是 (luRow,luCol+i)  (luRow+i,rlCol) (rlRow,rlCol-i) (rlRow-i,luCol)
			rotation(&(*matrix)[luRow][luCol+i], &(*matrix)[luRow+i][rlCol], &(*matrix)[rlRow][rlCol-i], &(*matrix)[rlRow-i][luCol])
		}

		// 更新左上角与右下角
		luRow++
		luCol++
		rlRow--
		rlCol--
	}

}

func PrintMatrix(matrix [][]int) {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			fmt.Printf(" %d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func rotation(ele1, ele2, ele3, ele4 *int) {
	val1 := *ele1
	val2 := *ele2
	val3 := *ele3
	val4 := *ele4

	*ele1 = val4
	*ele2 = val1
	*ele3 = val2
	*ele4 = val3
}
