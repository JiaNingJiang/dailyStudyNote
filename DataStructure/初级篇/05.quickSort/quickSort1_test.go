package quickSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	//arr := []int{2, 4, 4, 3, 7, 1, 4, 6}
	//arr := []int{3, 5, 6, 7, 4, 3, 8, 5}
	arr := []int{49, 61, 23, 10, 1, 1, 1}
	QuickSort1(arr)

	fmt.Println("排序后: ", arr)

	comparator.Comparator(QuickSort1)
}
