package sort

import (
	"fmt"
	"testing"
)

func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}

	smallSum := SmallSum(arr)

	fmt.Printf("小数之和:%d\n", smallSum)
}
