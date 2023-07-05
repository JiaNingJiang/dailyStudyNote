package lesson3

import (
	"fmt"
	"testing"
)

func TestSearchInMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 5, 9, 10},
		{2, 6, 11, 13},
		{7, 9, 15, 17},
	}
	target := 8

	if isExist := SearchInMatrix(matrix, target); isExist {
		fmt.Printf("目标: %d 存在于矩阵中\n", target)
	} else {
		fmt.Printf("目标: %d 不存在于矩阵中\n", target)
	}

}
