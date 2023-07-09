package utils

import (
	"fmt"
	"testing"
)

func TestMatrixMultiply(t *testing.T) {
	matrixA := [][]int{
		{1, 0},
		{0, 1},
	}

	matrixB := [][]int{
		{1, 1},
		{1, 0},
	}

	result := MatrixMultiply(matrixA, matrixB)

	// 输出结果矩阵
	for _, row := range result {
		fmt.Println(row)
	}
}
