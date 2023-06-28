package lesson1

import "math"

// 返回01矩阵中最大边框的边长
func MaxFrame1(matrix [][]int) int {

	width := len(matrix)     // 矩阵的宽度
	length := len(matrix[0]) // 矩阵的长度

	maxFrame := math.MinInt // 记录最大边框长度
	// 选择起点(row,col)
	for row := 0; row < width; row++ { // 横坐标
		for col := 0; col < length; col++ { // 纵坐标
			if matrix[row][col] == 0 { // 左上角起点不能是0
				continue
			}
			// 选择边长
			maxSide := int(math.Min(float64(length-col), float64(width-row))) // 选择的左上角决定最大边长
			for side := 1; side <= maxSide; side++ {
				standardFrame := true // true表示当前边框全为1，false表示边框存在0
				// 1.上方行必须全为1
				for start := col; start < col+side; start++ {
					if matrix[row][start] == 0 {
						standardFrame = false
						break
					}
				}
				if !standardFrame { // 发现边框上有0，即可退出，重新选择边长
					continue
				}
				// 2.下方行必须全为1
				for start := col; start < col+side; start++ {
					if matrix[row+side-1][start] == 0 {
						standardFrame = false
						break
					}
				}
				if !standardFrame {
					continue
				}
				// 3.左侧列必须全为1
				for start := row; start < row+side; start++ {
					if matrix[start][col] == 0 {
						standardFrame = false
						break
					}
				}
				if !standardFrame {
					continue
				}
				// 4.右侧列必须全为1
				for start := row; start < row+side; start++ {
					if matrix[start][col+side-1] == 0 {
						standardFrame = false
						break
					}
				}

				if standardFrame { // 表示边框的四个边全部是1
					maxFrame = int(math.Max(float64(maxFrame), float64(side)))
				}
			}
		}
	}
	return maxFrame
}

func MaxFrame2(matrix [][]int) int {

	width := len(matrix)     // 矩阵的宽度
	length := len(matrix[0]) // 矩阵的长度

	rightwardMatrix := make([][]int, length) // 记录matrix矩阵每一个节点所在行向右的连续1的个数
	downwardMatrix := make([][]int, length)  // 记录matrix矩阵每一个节点所在列向下的连续1的个数

	// 选择一个点(row,col),计算该点在所在行，从所在列开始向右的连续1的个数
	for row := 0; row < width; row++ { // 横坐标
		rightwardMatrix[row] = make([]int, length)
		for col := 0; col < length; col++ { // 纵坐标
			consecutive := 0
			// 计算点在row行,col列向右的连续1的个数
			for index := col; index < length; index++ {
				if matrix[row][index] == 1 {
					consecutive++
				} else {
					break
				}
			}
			rightwardMatrix[row][col] = consecutive
		}
	}

	// 选择一个点(row,col),计算该点在所在列，从所在行开始向下的连续1的个数
	for row := 0; row < width; row++ { // 横坐标
		downwardMatrix[row] = make([]int, length)
		for col := 0; col < length; col++ { // 纵坐标
			consecutive := 0
			// 计算点在row行,col列向下的连续1的个数
			for index := row; index < width; index++ {
				if matrix[index][col] == 1 {
					consecutive++
				} else {
					break
				}
			}
			downwardMatrix[row][col] = consecutive
		}
	}

	maxFrame := math.MinInt // 记录最大边框长度
	// 选择起点(row,col)
	for row := 0; row < width; row++ { // 横坐标
		for col := 0; col < length; col++ { // 纵坐标
			if matrix[row][col] == 0 { // 左上角不能是0
				continue
			}
			// 选择边长
			maxSide := int(math.Min(float64(length-col), float64(width-row))) // 选择的左上角决定最大边长
			for side := 1; side <= maxSide; side++ {
				standardFrame := true // true表示当前边框全为1，false表示边框存在0

				// 1.左上角向右必须全为1
				if rightwardMatrix[row][col] < side {
					standardFrame = false
					continue // 并非全为1，重新选边长
				}
				// 2.左上角向下必须全为1
				if downwardMatrix[row][col] < side {
					standardFrame = false
					continue // 并非全为1，重新选边长
				}
				// 3.左下角向右必须全为1
				if rightwardMatrix[row+side-1][col] < side {
					standardFrame = false
					continue // 并非全为1，重新选边长
				}
				// 4.右上角向下必须全为1
				if downwardMatrix[row][col+side-1] < side {
					standardFrame = false
					continue // 并非全为1，重新选边长
				}

				if standardFrame { // 表示边框的四个边全部是1
					maxFrame = int(math.Max(float64(maxFrame), float64(side)))
				}
			}
		}
	}
	return maxFrame
}
