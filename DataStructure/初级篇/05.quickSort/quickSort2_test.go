package quickSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestQuickSort2(t *testing.T) {
	//arr := []int{2, 4, 4, 3, 7, 1, 4, 6}
	arr := []int{3, 5, 6, 7, 4, 3, 8, 5}
	QuickSort2(arr)

	fmt.Println("排序后: ", arr)

	comparator.Comparator(QuickSort2)
}
