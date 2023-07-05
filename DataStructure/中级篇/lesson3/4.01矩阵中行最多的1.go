package lesson3

func MostOne(matrix [][]int) int {
	maxRow := len(matrix) - 1
	maxCol := len(matrix[0]) - 1

	// 从矩阵的右上角开始
	curRow := 0
	curCol := maxCol

	mostOne := 0
	for {
		if curRow < 0 || curRow > maxRow || curCol < 0 || curCol > maxCol {
			return mostOne
		}
		if matrix[curRow][curCol] == 1 { // 当前位置是1，则累加1的个数，并向左走
			mostOne++
			curCol--
		} else { // 当前位置是0,直接向下走
			curRow++
		}
	}
}
