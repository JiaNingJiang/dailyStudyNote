package lesson3

// 矩阵的每一行都是从小到大；每一列也是从小到大
func SearchInMatrix(matrix [][]int, target int) bool {
	maxRow := len(matrix) - 1    // 矩阵的最大行号
	maxCol := len(matrix[0]) - 1 // 矩阵的最大列号

	// 起点从矩阵的右上角开始(第0行，最后一列), 因此每次移动只能往下或者往左走
	curRow := 0
	curCol := maxCol
	for {
		if curRow < 0 || curRow > maxRow || curCol < 0 || curCol > maxCol {
			return false
		}
		if matrix[curRow][curCol] == target {
			return true
		}
		if matrix[curRow][curCol] > target { // 当前矩阵元素大于目标值，则往左走
			curCol--
			continue
		}
		if matrix[curRow][curCol] < target { // 当前矩阵元素小于目标值，则往下走
			curRow++
			continue
		}
	}
}
