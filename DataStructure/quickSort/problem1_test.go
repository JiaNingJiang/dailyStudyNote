package quickSort

import (
	"fmt"
	"testing"
)

func TestDiscreteSort(t *testing.T) {
	arr := []int{3, 5, 6, 7, 4, 3, 5, 8}
	target := 5
	DiscreteSort(arr, target)
	fmt.Println("排序后：", arr)
}
