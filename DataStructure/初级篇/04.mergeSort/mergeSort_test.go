package mergeSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {

	arr := []int{3, 1, 4, 6, 4, 9, 8}

	MergeSort(arr)

	fmt.Println("排序后: ", arr)

	comparator.Comparator(MergeSort)
}
