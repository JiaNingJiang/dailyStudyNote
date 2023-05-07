package nonComparativeSort

import (
	"DataStructure/comparator"
	"fmt"
	"testing"
)

func TestDemp(t *testing.T) {
	arr := []int{17, 13, 25, 100, 72}
	BucketSort(arr)

	fmt.Println("桶排序后: ", arr)

	comparator.Comparator(BucketSort)
}
