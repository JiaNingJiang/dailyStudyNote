package lesson1

import (
	"fmt"
	"testing"
)

func TestMaxFrame1(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 1, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 1, 0, 1, 1},
	}

	fmt.Println("方法一 最大边框长度: ", MaxFrame1(matrix))
	fmt.Println("方法二 最大边框长度: ", MaxFrame2(matrix))
}
