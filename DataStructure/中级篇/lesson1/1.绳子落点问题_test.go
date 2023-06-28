package lesson1

import (
	"fmt"
	"testing"
)

func TestFindRightNearest(t *testing.T) {
	arr := []int{0, 3, 6, 9, 12, 15, 18}
	target := 9

	value, index := FindRightNearest(arr, target)
	fmt.Printf("value = %d  index = %d \n", value, index)
}

func TestScopeIncluedSpot(t *testing.T) {
	arr := []int{0, 4, 5, 6, 9, 12, 15, 18}
	scope := 5

	fmt.Printf("最大覆盖点数为:%d ", ScopeIncluedSpot(arr, scope))
}
