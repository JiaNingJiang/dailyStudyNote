package utils

func MatrixMultiply(matrixA [][]int, matrixB [][]int) [][]int {
	rowsA := len(matrixA)
	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	// 检查矩阵尺寸是否兼容
	if colsA != rowsB {
		return nil
	}

	// 创建结果矩阵
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// 矩阵相乘
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}
