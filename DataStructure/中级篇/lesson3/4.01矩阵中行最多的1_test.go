package lesson3

import (
	"fmt"
	"testing"
)

func TestMostOne(t *testing.T) {
	matrix := [][]int{
		{1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1},
		{0, 1, 1, 1, 1, 1},
	}

	fmt.Println("行最多1的个数为: ", MostOne(matrix))
}
