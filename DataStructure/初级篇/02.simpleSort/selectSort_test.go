package simpleSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{17, 13, 100, 72, 25}
	SelectSort(arr)

	fmt.Println("选择排序后：", arr)

	comparator.Comparator(SelectSort)
}
