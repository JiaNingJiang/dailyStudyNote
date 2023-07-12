package lesson7

func MaxMatrixSum(matrix [][]int) int {
	rowCount := len(matrix)    // 行数
	colCount := len(matrix[0]) // 列数

	res := 0
	compressRow := make([]int, colCount)                 // 压缩行
	for startRow := 0; startRow < rowCount; startRow++ { // 起始行号
		for endRow := startRow; endRow < rowCount; endRow++ { // 终止行号
			// 1.进行列压缩
			for col := 0; col < colCount; col++ {
				for row := startRow; row <= endRow; row++ {
					compressRow[col] += matrix[row][col]
				}
			}
			// 2.获取压缩行的最大累加和
			compressMax := PostMaxScore(compressRow)
			// 3.更新子矩阵最大累加和
			res = getMax(res, compressMax)
			// 4.重新清空压缩行
			compressRow = make([]int, colCount)
		}
	}
	return res
}
