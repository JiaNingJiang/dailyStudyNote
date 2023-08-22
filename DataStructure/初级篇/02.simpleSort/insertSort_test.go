package simpleSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{17, 13, 100, 72, 25}
	InsertSort(arr)

	fmt.Println("插入排序后：", arr)

	comparator.Comparator(InsertSort)
}
