package nonComparativeSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{5, 5, 4, 3, 3, 2, 1}

	CountingSort(arr)

	fmt.Println("排序后：", arr)

	comparator.Comparator(CountingSort)
}
