package lesson4

import "fmt"

func SpiralPrint(matrix [][]int) {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	luRow := 0      // 左上角的行号
	luCol := 0      // 左上角的列号
	rlRow := maxRow // 右下角的行号
	rlCol := maxCol // 右下角的列号

	defer fmt.Println()
	for {
		if luRow == rlRow { // 左上角与右下角在同一行
			for col := luCol; col <= rlCol; col++ { // 打印完这一行后退出
				fmt.Printf(" %d ", matrix[luRow][col])
			}
			return
		}
		if luCol == rlCol { // 左上角与右下角在同一列
			for row := luRow; row <= rlRow; row++ { // 打印完这一列后退出
				fmt.Printf(" %d ", matrix[row][luCol])
			}
			return
		}
		if luRow > rlRow || luCol > rlCol { // 左上角在右下角的下方或者右侧时，退出
			return
		}

		// 1.打印矩阵顶层行
		for col := luCol; col < rlCol; col++ {
			fmt.Printf(" %d ", matrix[luRow][col])
		}
		// 2.打印矩阵的最右列
		for row := luRow; row < rlRow; row++ {
			fmt.Printf(" %d ", matrix[row][rlCol])
		}
		// 3.打印矩阵的最底层
		for col := rlCol; col > luCol; col-- {
			fmt.Printf(" %d ", matrix[rlRow][col])
		}
		// 4.打印矩阵的最左列
		for row := rlRow; row >= luRow+1; row-- {
			fmt.Printf(" %d ", matrix[row][luCol])
		}

		luRow++
		luCol++
		rlRow--
		rlCol--
	}

}
