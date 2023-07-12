package lesson7

import (
	"fmt"
	"testing"
)

func TestMaxMatrixSum(t *testing.T) {
	matrix := [][]int{
		{-5, 3, 6, 4},
		{-7, 9, -5, 3},
		{-10, 1, -200, 4},
	}
	fmt.Println("最大子矩阵累加和: ", MaxMatrixSum(matrix))
}
