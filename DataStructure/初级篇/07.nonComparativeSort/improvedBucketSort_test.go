package nonComparativeSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestImprovedBucketSort(t *testing.T) {
	arr := []int{13, 21, 11, 52, 62}
	ImprovedBucketSort(arr)

	fmt.Println("改进桶排序后: ", arr)

	comparator.Comparator(ImprovedBucketSort)
}
