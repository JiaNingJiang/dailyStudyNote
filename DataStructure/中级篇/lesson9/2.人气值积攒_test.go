package lesson9

import (
	"fmt"
	"testing"
)

func TestPopuValue(t *testing.T) {
	start := 0
	end := 50

	add := 3
	twice := 7
	del := 10

	fmt.Println("最少需要的金币数: ", PopuValue(start, end, add, twice, del))
}
