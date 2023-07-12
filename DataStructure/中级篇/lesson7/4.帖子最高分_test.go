package lesson7

import (
	"fmt"
	"testing"
)

func TestPostMaxScore(t *testing.T) {
	scores := []int{1, 1, -1, -10, 11, 4, -6, 9, 20, -10, -2}

	fmt.Println("最高分数: ", PostMaxScore(scores))
}
