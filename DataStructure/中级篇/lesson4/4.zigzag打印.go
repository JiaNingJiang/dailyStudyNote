package lesson4

import "fmt"

type Pos struct {
	row int
	col int
}

func ZigZag(matrix [][]int) {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	leftLower := Pos{0, 0}  // 对角线的左下角
	rightUpper := Pos{0, 0} // 对角线的右上角

	direction := true // 打印对角线时的方向

	for {
		if leftLower.row == maxRow && leftLower.col == maxCol { // 到达右下角
			fmt.Printf(" %d \n", matrix[leftLower.row][leftLower.col])
			return
		}

		if direction { // 从右上角到左下角顺序打印
			start := rightUpper
			end := leftLower

			for {
				if start.row == end.row && start.col == end.col {
					fmt.Printf(" %d ", matrix[start.row][start.col])
					break
				}
				fmt.Printf(" %d ", matrix[start.row][start.col])
				start.row++
				start.col--
			}

		} else { // 从左下角到右上角顺序打印
			start := leftLower
			end := rightUpper

			for {
				if start.row == end.row && start.col == end.col {
					fmt.Printf(" %d ", matrix[start.row][start.col])
					break
				}
				fmt.Printf(" %d ", matrix[start.row][start.col])
				start.row--
				start.col++
			}
		}

		direction = !direction
		if leftLower.row < maxRow {
			leftLower.row++
		} else {
			leftLower.col++
		}
		if rightUpper.col < maxCol {
			rightUpper.col++
		} else {
			rightUpper.row++
		}
	}
}
